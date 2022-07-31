package lookup

import "github.com/iancoleman/strcase"

// valMap is a convenience map which provides the getter interface.
type valMap struct {
	data map[string]string
}

// newValMap initializes a new empty valMap.
func newValMap() valMap {
	return valMap{data: map[string]string{}}
}

// Get returns the value for the given key.
func (v valMap) Get(key string) (string, bool) {
	key = strcase.ToSnake(key)
	val, ok := v.data[key]
	return val, ok
}

// Set sets the value for the given key.
func (v valMap) Set(key string, val string) {
	key = strcase.ToSnake(key)
	v.data[key] = val
}
