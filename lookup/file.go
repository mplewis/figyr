package lookup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

var knownExtensions = map[string]func(string) (ValMap, error){
	".yaml": newFromYAML,
	".yml":  newFromYAML,
	".json": newFromJSON,
	".toml": newFromTOML,
}

func newFromYAML(path string) (ValMap, error) {
	v := ValMap{}
	f, err := os.Open(path)
	if err != nil {
		return v, err
	}
	defer f.Close()
	err = yaml.NewDecoder(f).Decode(&v)
	return v, err
}

func newFromJSON(path string) (ValMap, error) {
	v := ValMap{}
	f, err := os.Open(path)
	if err != nil {
		return v, err
	}
	defer f.Close()
	err = yaml.NewDecoder(f).Decode(&v)
	return v, err
}

func newFromTOML(path string) (ValMap, error) {
	v := ValMap{}
	f, err := os.Open(path)
	if err != nil {
		return v, err
	}
	defer f.Close()
	_, err = toml.NewDecoder(f).Decode(&v)
	return v, err
}

func NewFromFile(path string) (ValMap, error) {
	ext := filepath.Ext(path)
	builder, supported := knownExtensions[ext]
	if !supported {
		return nil, fmt.Errorf("unknown extension %s", ext)
	}
	return builder(path)
}
