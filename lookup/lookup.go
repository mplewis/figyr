package lookup

import "os"

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
		args = &LookupArgs{
			OSArgs:     os.Args,
			FilePaths:  []string{}, // todo: which?
			EnvFetcher: os.Getenv,
		}
	}

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
