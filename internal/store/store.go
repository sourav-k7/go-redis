package store

var cache =  make(map[string]string);

func Set(key string, value string) (bool, error) {
	cache[key] = value
	return true, nil
}

func Get(key string) (string, bool) {
	value, exist := cache[key]
	if !exist {
		return "", exist
	}
	return value, true
}

func Delete(key string) {
	delete(cache, key)
}
