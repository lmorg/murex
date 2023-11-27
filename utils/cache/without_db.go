//go:build no_cachedb
// +build no_cachedb

package cache

import "time"

func SetPath(path string) {
	// no nothing
}

func createDb(namespace string) {
	// do nothing
}

func Read(namespace string, key string, ptr any) bool {
	return read(namespace, key, ptr)
}

func Write(namespace string, key string, value any, ttl time.Time) {
	write(namespace, key, value, ttl)
}

func CloseDb() {
	// do nothing
}
