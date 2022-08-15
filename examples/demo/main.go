package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mplewis/figyr"
)

type Config struct {
	RequiredString   string        `figyr:"required"`
	RequiredBool     bool          `figyr:"required"`
	RequiredInt      int           `figyr:"required"`
	RequiredFloat    float64       `figyr:"required"`
	RequiredDuration time.Duration `figyr:"required"`
	RequiredTime     time.Time     `figyr:"required"`

	OptionalString   string        `figyr:"optional"`
	OptionalBool     bool          `figyr:"optional"`
	OptionalInt      int           `figyr:"optional"`
	OptionalFloat    float64       `figyr:"optional"`
	OptionalDuration time.Duration `figyr:"optional"`
	OptionalTime     time.Time     `figyr:"optional"`

	DefaultString   string        `figyr:"default=hello"`
	DefaultBool     bool          `figyr:"default=true"`
	DefaultInt      int           `figyr:"default=42"`
	DefaultFloat    float64       `figyr:"default=3.14"`
	DefaultDuration time.Duration `figyr:"default=1s"`
	DefaultTime     time.Time     `figyr:"default=2022-08-15T02:03:39.570Z"`
}

func prettyPrint(x any) {
	typ := reflect.TypeOf(x)
	val := reflect.ValueOf(x)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		v := val.Field(i)
		vs := fmt.Sprintf("%v", v)
		if vs == "" {
			vs = "<empty string>"
		}
		fmt.Printf("%s: %s\n", f.Name, vs)
	}
}

func main() {
	var cfg Config
	figyr.MustParse(&cfg)
	prettyPrint(cfg)
}
