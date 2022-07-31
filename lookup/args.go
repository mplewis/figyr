package lookup

import "regexp"

// argMatcher matches a single `--key=value` argument.
var argMatcher = regexp.MustCompile(`^--([^=]+)=(.*)$`)

// NewFromArgs builds a config lookup from the given `os.Args` values.
func NewFromArgs(args []string) valMap {
	v := newValMap()
	for _, arg := range args {
		matches := argMatcher.FindStringSubmatch(arg)
		if matches == nil {
			continue
		}
		v.Set(matches[1], matches[2])
	}
	return v
}
