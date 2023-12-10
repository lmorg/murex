package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"unsafe"

	"github.com/lmorg/murex/debug"
)

func initCache(namespace string) {
	if configCacheDisabled {
		return
	}

	cache[namespace] = new(internalCacheT)
	cache[namespace].cache = make(map[string]*cacheItemT)
	disabled = false
	createDb(namespace)
}

func read(namespace string, key string, ptr any) bool {
	if ptr == nil {
		return false
	}

	var b []byte
	ok := cache[namespace].Read(key, unsafe.Pointer(&b))
	if !ok {
		return false
	}

	if err := json.Unmarshal(b, ptr); err != nil {
		if debug.Enabled {
			os.Stderr.WriteString(fmt.Sprintf("Error unmarshalling cache in "+namespace, err))
		}
		return false
	}

	return true
}

type dumpT struct {
	Internal interface{}
	CacheDb  interface{}
}

func Dump(ctx context.Context) (interface{}, error) {
	dump := make(map[string]dumpT)

	for namespace := range cache {
		internal := cache[namespace].Dump(ctx)
		cacheDb, err := listDb(ctx, namespace)
		dump[namespace] = dumpT{internal, cacheDb}

		if err != nil {
			return dump, err
		}
	}

	return dump, nil
}

func write(namespace string, key string, value any, ttl time.Time) {
	if value == nil {
		return
	}

	b, err := json.Marshal(value)
	if err != nil {
		if debug.Enabled {
			os.Stderr.WriteString(fmt.Sprintf("Error marshalling cache in "+namespace, err))
		}
		return
	}

	cache[namespace].Write(key, &b, ttl)
}

func ListCaches() []string {
	var (
		ret = make([]string, len(cache))
		i   int
	)

	for namespace := range cache {
		ret[i] = namespace
		i++
	}

	return ret
}

type trimmedT struct {
	Internal []string
	CacheDb  []string
}

func Trim(ctx context.Context) (interface{}, error) {
	trimmed := make(map[string]trimmedT)

	for namespace := range cache {
		internal := cache[namespace].Trim(ctx)
		cacheDb, err := trimDb(ctx, namespace)
		trimmed[namespace] = trimmedT{internal, cacheDb}

		if err != nil {
			return trimmed, err
		}
	}

	return trimmed, nil
}
