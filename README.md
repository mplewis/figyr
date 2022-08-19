# Figyr

[![Go Reference](https://pkg.go.dev/badge/github.com/mplewis/figyr.svg)](https://pkg.go.dev/github.com/mplewis/figyr)

Configure your Go app simply with zero configuration. Figyr parses config values
from various sources into a struct that you provide.

Figyr is extremely implicit and may not be for you! See the [FAQ](#faq) for more
information about the design decisions behind this library.

# Usage

The only way to configure Figyr is by building a config struct into which the
data will be loaded and tagging its fields:

```go
import "github.com/mplewis/figyr"

type Config struct {
	SiteName      string        `figyr:"required,description=The public name of the website"`
	Development   bool          `figyr:"optional,description=Run the server with development logging on"`
	CheckInterval time.Duration `figyr:"default=15s,description=How often to check for updates"`
}

func LoadConfig() (Config, error) {
  var cfg Config
  err := figyr.Parse(&cfg)
  return cfg, err
}
```

Figyr supports a help message generated from your flags:

```
$ ./my_app --help
Options:
    --site-name        required       The public name of the website
    --development      optional       Run the server with development logging on
    --check-interval   default: 15s   How often to check for updates
```

# Supported Types

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

# FAQ

**Why can't I change how Figyr works?**

Figyr is for people writing a quick Go program who need to be able to set some
keys and values, and parse them as typed Go values. That's all it supports.

**Why doesn't Figyr follow idiomatic Go principles?**

I wanted a library to pull config values in from the execution environment by
name and type without having to configure _anything else._ Figyr does exactly
this and nothing more.

**Why does Figyr try to guess the casing (e.g. `MyKey` vs `my-key` vs `MY_KEY`)
of the variables I use in my config?**

This makes Figyr compatible with config files that contain keys of any naming
convention. You shouldn't have to rework your team's existing conventions to
write a config file that's compatible with Figyr.

**This implicit configuration approach carries risk, right?**

Yes. If you're configuring your app with config files, environment variables,
and argument flags, it may be hard to track down how your variables are being
set. Cascading value overrides are explicitly a supported feature, but you
should be conservative in using them for your own sanity.

Figyr may not be for you or your use case, and I encourage you to evaluate
critically before using Figyr in your app.

**Why does Figyr have so many different ways to configure my app?**

Different runtime environments make it easier or harder to use various
configuration methods to set up your app. For example: config files support big
lists of values very well, but require a lot of boilerplate to add to a
Kubernetes deployment. Environment variables and config flags work in most
runtime environments, but one might be easier to use with your Infrastructure as
Code system. Use whichever config methods work best for you.

# TODO

- `--help` messages
- Cut a proper release
