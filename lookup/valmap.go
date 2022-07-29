package lookup

import "github.com/iancoleman/strcase"

type valMap struct {
	data map[string]string
}

func newValMap() valMap {
	return valMap{data: map[string]string{}}
}

func (v valMap) Get(key string) (string, bool) {
	key = strcase.ToSnake(key)
	val, ok := v.data[key]
	return val, ok
}

func (v valMap) Set(key string, val string) {
	key = strcase.ToSnake(key)
	v.data[key] = val
}
