package lookup

import (
	"os"
	"regexp"
)

var configArgMatcher = regexp.MustCompile(`^--config=(.*)$`)

type getter interface {
	Get(string) (string, bool)
}

type LookupArgs struct {
	OSArgs     []string
	FilePaths  []string
	EnvFetcher func(string) string
}

func NewFromDefaults(args *LookupArgs) (getter, error) {
	if args == nil {
		args = &LookupArgs{OSArgs: os.Args, EnvFetcher: os.Getenv}
	}
	args.FilePaths = getConfigPaths(args.OSArgs)

	fetchers := []getter{}
	fetchers = append(fetchers, NewFromArgs(args.OSArgs))
	for _, path := range args.FilePaths {
		fetcher, err := NewFromFile(path)
		if err != nil {
			return nil, err
		}
		fetchers = append(fetchers, fetcher)
	}
	fetchers = append(fetchers, NewFromEnv(args.EnvFetcher))

	return Combine(fetchers...), nil
}

func getConfigPaths(args []string) []string {
	var configPaths []string
	for _, arg := range args {
		matches := configArgMatcher.FindStringSubmatch(arg)
		if matches == nil {
			continue
		}
		configPaths = append(configPaths, matches[1])
	}
	// reverse it - config files listed last should take priority
	for i, j := 0, len(configPaths)-1; i < j; i, j = i+1, j-1 {
		configPaths[i], configPaths[j] = configPaths[j], configPaths[i]
	}
	return configPaths
}
