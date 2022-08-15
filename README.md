# Figyr

[![Go Reference](https://pkg.go.dev/badge/github.com/mplewis/figyr.svg)](https://pkg.go.dev/github.com/mplewis/figyr)

Configure your Go app simply with zero configuration. Figyr parses config values
from various sources into a struct that you provide.

# Usage

The only way to configure Figyr is by building a config struct into which the
data will be loaded and tagging its fields:

```go
import "github.com/mplewis/figyr"

type Config struct {
	SiteName      string        `figyr:"required"`
	Development   bool          `figyr:"optional"`
	CheckInterval time.Duration `figyr:"default=15s"`
}

func LoadConfig() (Config, error) {
  var cfg Config
  err := figyr.Parse(&cfg)
  return cfg, err
}
```

Figyr supports the following types, defined in [coerce.go](refparse/coerce.go):

- string
- [bool](https://pkg.go.dev/strconv#ParseBool)
- integer
- float
- [time.Duration](https://pkg.go.dev/time#ParseDuration)
- [time.Time](https://pkg.go.dev/time#RFC3339) (RFC3339 format)

For an extended example, run [`bin/demo`](bin/demo) and read the source in
[`examples/demo/main.go`](examples/demo/main.go).

# Value Precedence

Values are parsed from sources in this order (values higher in the list take
priority):

1. `--key=val` command-line flags
2. Values parsed from `--config=cfg-file.yaml` files. If `--config` is specified
   more than once, files specified later take priority over files specified
   earlier.
3. `KEY=val` environment variables

Figyr uses [strcase](https://github.com/iancoleman/strcase) to convert key
names, so you can use any mix of `CamelCase`, `kebab-case`, `snake_case`, and
`SCREAMING_SNAKE_CASE` in all places where you specify config values.
