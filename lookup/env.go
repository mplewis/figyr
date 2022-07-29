package lookup

import "os"

type envFetcher struct {
	fetchFromEnv func(string) string
}

func (e envFetcher) Fetch(key string) (string, bool) {
	val := e.fetchFromEnv(key)
	if val == "" {
		return "", false
	}
	return val, true
}

func NewFromEnv(fetcher func(string) string) envFetcher {
	if fetcher == nil {
		fetcher = os.Getenv
	}
	return envFetcher{fetcher}
}
