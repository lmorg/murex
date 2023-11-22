//go:build no_cache
// +build no_cache

package cache

import "time"

func Read(namespace string, key string, ptr any) bool {
	return false
}

func Write(namespace string, key string, value any, ttl time.Time) {
	// do nothing
}
