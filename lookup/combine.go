package lookup

type combined struct {
	getters []getter
}

func (c combined) Get(key string) (string, bool) {
	for _, getter := range c.getters {
		val, ok := getter.Get(key)
		if ok {
			return val, true
		}
	}
	return "", false
}

func Combine(getters ...getter) combined {
	return combined{getters: getters}
}
