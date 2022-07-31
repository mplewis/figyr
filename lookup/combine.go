package lookup

// combined looks for config values from multiple getters sequentially.
type combined struct {
	getters []getter
}

// Get returns the value for the given key.
func (c combined) Get(key string) (string, bool) {
	for _, getter := range c.getters {
		val, ok := getter.Get(key)
		if ok {
			return val, true
		}
	}
	return "", false
}

// Combine combines the given getters into a single getter.
func Combine(getters ...getter) combined {
	return combined{getters: getters}
}
