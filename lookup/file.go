package lookup

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// decoder is an interface to decode data from a reader into a destination struct.
type decoder func(in io.Reader, dest any) error

// knownExtensions is the list of supported config file extensions.
var knownExtensions = map[string]func(string) (valMap, error){
	".json": newFromJSON,
	".yaml": newFromYAML,
	".yml":  newFromYAML,
}

// newFromDecoder builds a valMap from a file path using the given decoder.
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

// newFromJson builds a valMap from a JSON file at the given path.
func newFromJSON(path string) (valMap, error) {
	return newFromDecoder(path, func(in io.Reader, dest any) error {
		return json.NewDecoder(in).Decode(dest)
	})
}

// newFromYAML builds a valMap from a YAML file at the given path.
func newFromYAML(path string) (valMap, error) {
	return newFromDecoder(path, func(in io.Reader, dest any) error {
		return yaml.NewDecoder(in).Decode(dest)
	})
}

// newFromFile builds a valMap from a file at the given path.
func NewFromFile(path string) (valMap, error) {
	ext := filepath.Ext(path)
	builder, supported := knownExtensions[ext]
	if !supported {
		return newValMap(), fmt.Errorf("unknown extension %s", ext)
	}
	return builder(path)
}
