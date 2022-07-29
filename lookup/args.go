package lookup

import "regexp"

var argMatcher = regexp.MustCompile(`^--([^=]+)=(.*)$`)

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
