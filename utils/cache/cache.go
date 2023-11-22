//go:build !no_cache
// +build !no_cache

package cache

import (
	"time"

	"github.com/lmorg/murex/utils/cache/cachelib"
)

func Read(namespace string, key string, ptr any) bool {
	return cachelib.Read(namespace, key, ptr)
}

func Write(namespace string, key string, value any, ttl time.Time) {
	cachelib.Write(namespace, key, value, ttl)
}
