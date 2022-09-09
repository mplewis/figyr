package refparse

import (
	"fmt"
	"reflect"

	"github.com/mplewis/figyr/lookup"
	"github.com/mplewis/figyr/types"
)

// Parse parses a config struct and fills it with coerced data, or returns an error if something went wrong.
func Parse(po types.ParserOptions, dst interface{}, source lookup.Getter) error {
	typ := reflect.TypeOf(dst)
	if typ.Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}

	el := typ.Elem()
	defs := []FieldDef{}
	parseErrors := []error{}
	valueErrors := []error{}
	for i := 0; i < el.NumField(); i++ {
		f := el.Field(i)
		val, annotated := f.Tag.Lookup("figyr")
		if !annotated {
			continue
		}

		def, err := BuildFieldDef(f.Name, f.Type, val)
		if err != nil {
			parseErrors = append(parseErrors, fmt.Errorf("%s: %v", f.Name, err))
			continue
		}
		defs = append(defs, def)

		val, found := source.Get(f.Name)
		if !found && def.Default != "" {
			val = def.Default
		}
		result, err := def.Coerce(val)
		if err != nil {
			valueErrors = append(valueErrors, fmt.Errorf("%s: %v", f.Name, err))
			continue
		}
		reflect.ValueOf(dst).Elem().Field(i).Set(reflect.ValueOf(result))
	}

	if len(parseErrors) > 0 {
		return fmt.Errorf("%v", parseErrors)
	}
	printHelpAndExitIfRequested(po, defs)
	if len(valueErrors) > 0 {
		return fmt.Errorf("%v", valueErrors)
	}
	return nil
}
