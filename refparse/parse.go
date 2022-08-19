package refparse

import (
	"fmt"
	"reflect"

	"github.com/mplewis/figyr/lookup"
)

// Parse parses a config struct and fills it with coerced data, or returns an error if something went wrong.
func Parse(dst interface{}, source lookup.Getter) error {
	typ := reflect.TypeOf(dst)
	if typ.Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}

	el := typ.Elem()
	defs := []FieldDef{}
	errors := []error{}
	for i := 0; i < el.NumField(); i++ {
		f := el.Field(i)
		val, annotated := f.Tag.Lookup("figyr")
		if !annotated {
			continue
		}

		def, err := BuildFieldDef(f.Name, f.Type, val)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		defs = append(defs, def)

		val, found := source.Get(f.Name)
		if !found && def.Default != "" {
			val = def.Default
		}
		result, err := def.Coerce(val)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		reflect.ValueOf(dst).Elem().Field(i).Set(reflect.ValueOf(result))
	}

	printHelpAndExitIfRequested(defs)

	if len(errors) > 0 {
		return fmt.Errorf("%v", errors)
	}
	return nil
}
