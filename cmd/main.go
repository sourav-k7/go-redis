package main

import "github.com/sourav-k7/go-redis/internal/server"

type Cache struct {
	cache map[string]string
}

func (store Cache) SetCache(key string, value string) (bool, error) {
	store.cache[key] = value
	return true, nil
}

func (store Cache) GetCache(key string) (string, bool) {
	value, exist := store.cache[key]
	if !exist {
		return "", exist
	}
	return value, true
}

func (store Cache) DeleteCache(key string) {
	delete(store.cache, key)
}

func main() {
	// cache := Cache{
	// 	cache: make(map[string]string),
	// }
	// cache.SetCache("test", "val")
	// value, exist := cache.GetCache("test")
	// if exist {
	// 	fmt.Println(value)
	// }
	// cache.DeleteCache("test")
	server.ServerSetup()
}
