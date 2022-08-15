package refparse

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// coerce converts the given value to a type compatible with the field.
func coerce(fieldName string, typ reflect.Type, raw string) (any, error) {
	empty := raw == ""
	ident := typ.Name()
	if pkg := typ.PkgPath(); pkg != "" {
		ident = fmt.Sprintf("%s.%s", pkg, typ.Name())
	}

	switch ident {
	case "string":
		return raw, nil

	case "bool":
		if empty {
			return false, nil
		}
		return strconv.ParseBool(raw)

	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr":
		if empty {
			return 0, nil
		}
		return strconv.Atoi(raw)

	case "float32", "float64":
		if empty {
			return 0.0, nil
		}
		return strconv.ParseFloat(raw, 64)

	case "time.Duration":
		if empty {
			return 0 * time.Second, nil
		}
		return time.ParseDuration(raw)

	case "time.Time":
		if empty {
			return time.Time{}, nil
		}
		return time.Parse(time.RFC3339, raw)

	default:
		return nil, fmt.Errorf("unsupported type %s for field %s", ident, fieldName)
	}
}
