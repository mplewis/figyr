package main

import (
	"reflect"
	"time"

	"github.com/k0kubun/pp/v3"
	"github.com/mplewis/figyr"
	"github.com/mplewis/figyr/lookup"
)

type Config struct {
	// SiteName    string `figyr:"default:example.com"`
	// ConnCount   int    `figyr:"default:10"`
	// Development bool   `figyr:"default:false"`
	SiteName           string        `figyr:"required"`
	ConnCount          int           `figyr:"default:10"`
	Development        bool          `figyr:"optional"`
	Hosts              []string      `figyr:"default:localhost,1.2.3.4"`
	CheckInterval      time.Duration `figyr:"required"`
	StartDate          time.Time     `figyr:"optional"`
	SomeAdditionalData string        // not set by config
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var cfg Config

	vals, err := lookup.NewFromDefaults(nil)
	check(err)
	defs, err := figyr.ParseFieldDefs(&cfg)
	pp.Println(defs)
	for _, def := range defs {
		t := def.Type
		pp.Println(t, t.PkgPath(), t.Name())
	}

	check(err)
	for name, def := range defs {
		rawVal, found := vals.Get(name)
		if found {
			val, err := def.Coerce(rawVal)
			check(err)
			// fmt.Printf("setting %s to %s -> %#v\n", name, rawVal, val)
			reflect.ValueOf(&cfg).Elem().FieldByName(name).Set(reflect.ValueOf(val))
		}
	}

	pp.Println(cfg)
}
