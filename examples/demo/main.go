package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mplewis/figyr"
)

const desc = "This is a demonstration of figyr's basic functions."

type Config struct {
	RequiredString   string        `figyr:"required,description=A string that must be provided"`
	RequiredBool     bool          `figyr:"required,description=A bool that must be provided"`
	RequiredInt      int           `figyr:"required,description=An int that must be provided"`
	RequiredFloat    float64       `figyr:"required,description=A float that must be provided"`
	RequiredDuration time.Duration `figyr:"required,description=A duration that must be provided"`
	RequiredTime     time.Time     `figyr:"required,description=A time that must be provided"`

	OptionalString   string        `figyr:"optional,description=A string that defaults to \"\""`
	OptionalBool     bool          `figyr:"optional,description=A bool that defaults to false"`
	OptionalInt      int           `figyr:"optional,description=An int that defaults to 0"`
	OptionalFloat    float64       `figyr:"optional,description=A float that defaults to 0.0"`
	OptionalDuration time.Duration `figyr:"optional,description=A duration that defaults to 0s"`
	OptionalTime     time.Time     `figyr:"optional,description=A time that defaults to the zero time"`

	DefaultString   string        `figyr:"default=hello,description=A string that defaults to \"hello\""`
	DefaultBool     bool          `figyr:"default=true,description=A bool that defaults to true"`
	DefaultInt      int           `figyr:"default=42,description=An int that defaults to 42"`
	DefaultFloat    float64       `figyr:"default=3.14,description=A float that defaults to 3.14"`
	DefaultDuration time.Duration `figyr:"default=1s,description=A duration that defaults to 1s"`
	DefaultTime     time.Time     `figyr:"default=2022-08-15T02:03:39.570Z,description=A time that defaults to 2022-08-15T02:03:39.570Z"`
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
	figyr.New(desc).MustParse(&cfg)
	prettyPrint(cfg)
}
