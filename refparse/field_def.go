package refparse

import (
	"fmt"
	"reflect"
	"strings"
)

// FieldDef holds the type definition of a field in a custom struct, as well as its Figyr metadata.
type FieldDef struct {
	Name     string
	Type     reflect.Type
	Required bool
	Default  string
}

// Coerce converts the given value to a type compatible with the field.
func (f *FieldDef) Coerce(raw string) (any, error) {
	if f.Required && raw == "" {
		return nil, fmt.Errorf("missing required field %s", f.Name)
	}
	if raw == "" {
		raw = f.Default
	}
	return coerce(f.Name, f.Type, raw)
}

// buildFieldDef builds a FieldDef from the given field name, type, and Figyr tag.
func buildFieldDef(name string, typ reflect.Type, tagVal string) (FieldDef, error) {
	mandate, fallback, _ := strings.Cut(tagVal, "=")
	def := FieldDef{Name: name, Type: typ}

	switch mandate {
	case "required":
		def.Required = true

	case "optional":
		def.Required = false

	case "default":
		def.Required = false
		if fallback == "" {
			return def, fmt.Errorf("missing default value for field %s", name)
		}
		def.Default = fallback

	default:
		return def, fmt.Errorf("unknown mandate %s for field %s", mandate, name)
	}

	return def, nil
}
