package figyr

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// FieldDef holds the type definition of a field in a custom struct, as well as its Figyr metadata.
type FieldDef struct {
	Name     string
	Type     reflect.Type
	Required bool
	Default  string
}

// Coerce converts the given value to a type compatible with the field.
func (f *FieldDef) Coerce(in string) (any, error) {
	if f.Required && in == "" {
		return nil, fmt.Errorf("missing required field %s", f.Name)
	}
	if in == "" {
		in = f.Default
	}

	ident := f.Type.Name()
	if pkg := f.Type.PkgPath(); pkg != "" {
		ident = fmt.Sprintf("%s.%s", pkg, f.Type.Name())
	}
	switch ident {

	case "string":
		return in, nil

	case "int", "int64":
		return strconv.Atoi(in)

	case "bool":
		return strconv.ParseBool(in)

	case "float64":
		return strconv.ParseFloat(in, 64)

	case "time.Duration":
		return time.ParseDuration(in)

	default:
		return nil, fmt.Errorf("unsupported type %s for field %s", ident, f.Name)
	}
}

// buildFieldDef builds a FieldDef from the given field name, type, and Figyr tag.
func buildFieldDef(name string, typ reflect.Type, tagVal string) (FieldDef, error) {
	bits := strings.Split(tagVal, ":")
	mandate := bits[0]
	def := FieldDef{Name: name, Type: typ}

	switch mandate {
	case "required":
		def.Required = true

	case "optional":
		def.Required = false

	case "default":
		def.Required = false
		fallback := bits[1]
		if fallback == "" {
			return def, fmt.Errorf("missing default value for field %s", name)
		}
		def.Default = fallback

	default:
		return def, fmt.Errorf("unknown mandate %s for field %s", mandate, name)
	}

	return def, nil
}

// ParseFieldDefs parses the given struct and returns a map of struct field names to FieldDefs.
func ParseFieldDefs(dst interface{}) (map[string]FieldDef, error) {
	errors := []error{}

	typ := reflect.TypeOf(dst)
	if typ.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("dst must be a pointer")
	}

	el := typ.Elem()
	count := el.NumField()
	defs := map[string]FieldDef{}
	for i := 0; i < count; i++ {
		f := el.Field(i)
		val, annotated := f.Tag.Lookup("figyr")
		if !annotated {
			continue
		}

		def, err := buildFieldDef(f.Name, f.Type, val)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		defs[f.Name] = def
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("%d errors occurred: %v", len(errors), errors)
	}
	return defs, nil
}
