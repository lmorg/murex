//go:build no_cachedb
// +build no_cachedb

package cache

import (
	"context"
	"time"
)

func SetPath(_ string) {
	// no nothing
}

func createDb(_ string) {
	// do nothing
}

func Read(namespace string, key string, ptr any) bool {
	return read(namespace, key, ptr)
}

func listDb(_ context.Context, _ string) (interface{}, error) {
	return nil, nil
}

func Write(namespace string, key string, value any, ttl time.Time) {
	write(namespace, key, value, ttl)
}

func trimDb(_ context.Context, _ string) ([]string, error) {
	return nil, nil
}

func flushDb(_ context.Context, _ string) ([]string, error) {
	return nil, nil
}

func CloseDb() {
	// do nothing
}
