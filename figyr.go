package figyr

import (
	"fmt"
	"reflect"

	"github.com/k0kubun/pp/v3"
)

func ParseType(dst interface{}) error {
	typ := reflect.TypeOf(dst)
	if typ.Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}
	el := typ.Elem()
	fmt.Println(el)
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
		pp.Println(def)

		result, err := def.Coerce(def.Default)
		if err != nil {
			return err
		}
		reflect.ValueOf(dst).Elem().Field(i).Set(reflect.ValueOf(result))
	}
	return nil
}
