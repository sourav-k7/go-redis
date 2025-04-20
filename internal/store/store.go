package store

type StringCache struct {
	cache map[string]string 
}

func (store StringCache) Set(key string, value string) (bool, error) {
	store.cache[key] = value
	return true, nil
}

func (store StringCache) Get(key string) (string, bool) {
	value, exist := store.cache[key]
	if !exist {
		return "", exist
	}
	return value, true
}

func (store StringCache) Delete(key string) {
	delete(store.cache, key)
}
