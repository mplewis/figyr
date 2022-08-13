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
	for i := 0; i < el.NumField(); i++ {
		f := el.Field(i)
		val, annotated := f.Tag.Lookup("figyr")
		if !annotated {
			continue
		}

		def, err := buildFieldDef(f.Name, f.Type, val)
		if err != nil {
			return err
		}

		val, found := source.Get(f.Name)
		if !found && def.Default != "" {
			val = def.Default
		}
		result, err := def.Coerce(val)
		if err != nil {
			return err
		}
		reflect.ValueOf(dst).Elem().Field(i).Set(reflect.ValueOf(result))
	}
	return nil
}
