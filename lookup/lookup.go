package lookup

import "github.com/iancoleman/strcase"

type fetcher interface {
	Fetch(string) (string, bool)
}

type ValMap struct {
	data map[string]string
}

func NewValMap() ValMap {
	return ValMap{data: map[string]string{}}
}

func (v ValMap) Get(key string) (string, bool) {
	key = strcase.ToSnake(key)
	val, ok := v.data[key]
	return val, ok
}

func (v ValMap) Set(key string, val string) {
	key = strcase.ToSnake(key)
	v.data[key] = val
}
