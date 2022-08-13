package lookup

import (
	"os"
	"regexp"
)

// configArgMatcher matches a config file path argument.
var configArgMatcher = regexp.MustCompile(`^--config=(.*)$`)

// Getter is an interface for getting a named value if it exists.
type Getter interface {
	Get(string) (string, bool)
}

// LookupArgs is a config that specifies what config sources should be used.
type LookupArgs struct {
	OSArgs     []string
	FilePaths  []string
	EnvFetcher func(string) string
}

// NewFromDefaults builds a config lookup from the default config sources.
// In order of priority:
// 1. `os.Args`
// 2. `--config=myconfig.json` arguments
// 3. Environment variables
func NewFromDefaults(args *LookupArgs) (Getter, error) {
	if args == nil {
		args = &LookupArgs{OSArgs: os.Args, EnvFetcher: os.Getenv}
	}
	args.FilePaths = getConfigPaths(args.OSArgs)

	fetchers := []Getter{}
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

// getConfigPaths extracts the config file paths from the given args,
// returning them in reverse order to maintain the expected cascade priority.
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
