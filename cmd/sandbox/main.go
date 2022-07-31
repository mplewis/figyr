package main

import (
	"time"

	"github.com/k0kubun/pp/v3"
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

func main() {
	// var cfg Config
	// pp.Println(figyr.ParseFieldDefs(&cfg))
	// fmt.Println(cfg)
	pp.Println(lookup.NewFromDefaults(nil))
}
