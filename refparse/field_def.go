package refparse

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

var mandateSepMatcher = regexp.MustCompile(`[^\\],`)
var mustHaveOneMandateOf = []string{"required", "optional", "default"}

// FieldDef holds the type definition of a field in a custom struct, as well as its Figyr metadata.
type FieldDef struct {
	Name        string
	Type        reflect.Type
	Required    bool
	Default     string
	Description string
}

func (f FieldDef) Constraint() string {
	if f.Required {
		return "required"
	}
	if f.Default != "" {
		return fmt.Sprintf("default: %s", f.Default)
	}
	return "optional"
}

// Coerce converts the given value to a type compatible with the field.
func (f *FieldDef) Coerce(raw string) (any, error) {
	if f.Required && raw == "" {
		return nil, fmt.Errorf("missing value for required field")
	}
	if raw == "" {
		raw = f.Default
	}
	return coerce(f.Name, f.Type, raw)
}

func splitMandate(s string) (string, string) {
	k, v, _ := strings.Cut(s, "=")
	v = strings.ReplaceAll(v, `\,`, `,`)
	return k, v
}

func BreakValueIntoMandates(value string) map[string]string {
	result := map[string]string{}
	out := mandateSepMatcher.FindAllStringIndex(value, -1)
	last := 0
	for _, pair := range out {
		_, y := pair[0], pair[1]
		k, v := splitMandate(value[last : y-1])
		result[k] = v
		last = y
	}
	k, v := splitMandate(value[last:])
	result[k] = v
	return result
}

// BuildFieldDef builds a FieldDef from the given field name, type, and Figyr tag.
func BuildFieldDef(name string, typ reflect.Type, tagVal string) (FieldDef, error) {
	mandates := BreakValueIntoMandates(tagVal)
	def := FieldDef{Name: name, Type: typ}

	found := []string{}
	for k := range mandates {
		if slices.Contains(mustHaveOneMandateOf, k) {
			found = append(found, k)
		}
	}
	if len(found) != 1 {
		return FieldDef{}, fmt.Errorf("expected exactly one of %v, but found %v", mustHaveOneMandateOf, found)
	}

	for mandate, value := range mandates {
		switch mandate {
		case "required":
			def.Required = true

		case "optional":
			def.Required = false

		case "default":
			def.Required = false
			if value == "" {
				return def, fmt.Errorf("missing default value for field %s", name)
			}
			def.Default = value

		case "description":
			def.Description = value

		default:
			return def, fmt.Errorf("unknown mandate %s for field %s", mandate, name)
		}
	}

	return def, nil
}
