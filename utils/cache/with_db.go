//go:build !no_cachedb
// +build !no_cachedb

package cache

import (
	"context"
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

func listDb(ctx context.Context, namespace string) (interface{}, error) {
	return cachedb.List(ctx, namespace)
}

func Write(namespace string, key string, value any, ttl time.Time) {
	write(namespace, key, value, ttl)
	cachedb.Write(namespace, key, value, ttl)
}

func trimDb(ctx context.Context, namespace string) ([]string, error) {
	return cachedb.Trim(ctx, namespace)
}

func clearDb(ctx context.Context, namespace string) ([]string, error) {
	return cachedb.Clear(ctx, namespace)
}

func CloseDb() {
	cachedb.CloseDb()
}
