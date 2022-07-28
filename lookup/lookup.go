package lookup

type fetcher interface {
	Fetch(string) (string, bool)
}

type ValMap map[string]string

func (v ValMap) Fetch(key string) (string, bool) {
	val, ok := v[key]
	return val, ok
}
