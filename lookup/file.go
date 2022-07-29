package lookup

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type decoder func(in io.Reader, dest any) error

var knownExtensions = map[string]func(string) (valMap, error){
	".json": newFromJSON,
	".yaml": newFromYAML,
	".yml":  newFromYAML,
}

func newFromDecoder(path string, dec decoder) (valMap, error) {
	vm := newValMap()
	f, err := os.Open(path)
	if err != nil {
		return vm, err
	}
	defer f.Close()
	kvs := map[string]any{}
	err = dec(f, &kvs)
	if err != nil {
		return vm, err
	}
	for k, v := range kvs {
		vm.Set(k, fmt.Sprint(v))
	}
	return vm, nil
}

func newFromJSON(path string) (valMap, error) {
	return newFromDecoder(path, func(in io.Reader, dest any) error {
		return json.NewDecoder(in).Decode(dest)
	})
}

func newFromYAML(path string) (valMap, error) {
	return newFromDecoder(path, func(in io.Reader, dest any) error {
		return yaml.NewDecoder(in).Decode(dest)
	})
}

func NewFromFile(path string) (valMap, error) {
	ext := filepath.Ext(path)
	builder, supported := knownExtensions[ext]
	if !supported {
		return newValMap(), fmt.Errorf("unknown extension %s", ext)
	}
	return builder(path)
}
