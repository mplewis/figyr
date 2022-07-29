package lookup

import (
	"os"

	"github.com/iancoleman/strcase"
)

type envFetcher struct {
	fetchFromEnv func(string) string
}

func (e envFetcher) Get(key string) (string, bool) {
	key = strcase.ToScreamingSnake(key)
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
	return envFetcher{fetchFromEnv: fetcher}
}
