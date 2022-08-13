package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/mplewis/figyr"
)

type Config struct {
	RequiredString   string        `figyr:"required"`
	RequiredBool     bool          `figyr:"required"`
	RequiredInt      int           `figyr:"required"`
	RequiredFloat    float64       `figyr:"required"`
	RequiredDuration time.Duration `figyr:"required"`

	OptionalString   string        `figyr:"optional"`
	OptionalBool     bool          `figyr:"optional"`
	OptionalInt      int           `figyr:"optional"`
	OptionalFloat    float64       `figyr:"optional"`
	OptionalDuration time.Duration `figyr:"optional"`

	DefaultString   string        `figyr:"default=hello"`
	DefaultBool     bool          `figyr:"default=true"`
	DefaultInt      int           `figyr:"default=42"`
	DefaultFloat    float64       `figyr:"default=3.14"`
	DefaultDuration time.Duration `figyr:"default=1s"`
}

func prettyPrint(x any) {
	all := fmt.Sprintf("%#v\n", x)
	for _, s := range regexp.MustCompile(`[{},]`).Split(all, -1) {
		fmt.Println(strings.TrimSpace(s))
	}
}

func main() {
	var cfg Config
	err := figyr.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	prettyPrint(cfg)
}
