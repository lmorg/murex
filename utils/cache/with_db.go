//go:build !no_cachedb
// +build !no_cachedb

package cache

import (
	"time"

	"github.com/lmorg/murex/utils/cache/cachedb"
)

func SetPath(path string) {
	cachedb.Path = path
}

func createDb(namespace string) {
	cachedb.CreateTable(namespace)
}

func Read(namespace string, key string, ptr any) bool {
	if !read(namespace, key, ptr) {
		return cachedb.Read(namespace, key, ptr)
	}
	return true
}

func Write(namespace string, key string, value any, ttl time.Time) {
	write(namespace, key, value, ttl)
	cachedb.Write(namespace, key, value, ttl)
}

func CloseDb() {
	cachedb.CloseDb()
}
