package figyr

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type FieldDef struct {
	Name     string
	Type     reflect.Type
	Required bool
	Default  string
}

func (f *FieldDef) hasDefault() bool {
	return f.Default != ""
}

func (f *FieldDef) ignored() bool {
	return f.Type == nil
}

func (f *FieldDef) Coerce(in string) (any, error) {
	if f.Required && in == "" {
		return nil, fmt.Errorf("missing required field %s", f.Name)
	}

	switch f.Type.Kind() {
	case reflect.String:
		if !f.hasDefault() {
			return "", nil
		}
		return in, nil

	case reflect.Int:
		if !f.hasDefault() {
			return 0, nil
		}
		return strconv.Atoi(in)

	case reflect.Bool:
		if !f.hasDefault() {
			return false, nil
		}
		return strconv.ParseBool(in)

	case reflect.Float64:
		if !f.hasDefault() {
			return 0, nil
		}
		return strconv.ParseFloat(in, 64)

	default:
		return nil, fmt.Errorf("unsupported type %s", f.Type.Kind())
	}
}

func buildFieldDef(name string, typ reflect.Type, val string) (FieldDef, error) {
	bits := strings.Split(val, ":")
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

func ParseFieldDefs(dst interface{}) ([]FieldDef, error) {
	errors := []error{}

	typ := reflect.TypeOf(dst)
	if typ.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("dst must be a pointer")
	}

	el := typ.Elem()
	count := el.NumField()
	fields := make([]FieldDef, count)
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

		fields[i] = def
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("%d errors occurred: %v", len(errors), errors)
	}
	return fields, nil
}
