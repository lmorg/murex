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

func read(namespace string, key string, ptr any) bool {
	if ptr == nil {
		return false
	}

	var b []byte
	ic, ok := cache[namespace]
	if !ok {
		//panic(fmt.Sprintf("invalid namespace: '%s'", namespace))
		initNamespace(namespace)
		ic = cache[namespace]
	}

	ok = ic.Read(key, unsafe.Pointer(&b))
	if !ok {
		return false
	}

	if err := json.Unmarshal(b, ptr); err != nil {
		if debug.Enabled {
			os.Stderr.WriteString(fmt.Sprintf("!!! error unmarshalling cache in '%s': %s !!!\n!!! cache value: '%s' !!!", namespace, err.Error(), string(b)))
		}
		return false
	}

	return true
}

type dumpT struct {
	Internal any
	CacheDb  any
}

func Dump(ctx context.Context) (any, error) {
	dump := make(map[string]dumpT)

	for _, namespace := range ListNamespaces() {
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

type trimmedT struct {
	Internal []string
	CacheDb  []string
}

// Trim removes stale cache values from the cache databases
func Trim(ctx context.Context) (any, error) {
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

// Clear the cache completely
func Clear(ctx context.Context) (any, error) {
	flushed := make(map[string]trimmedT)

	for namespace := range cache {
		internal := cache[namespace].Clear(ctx)
		cacheDb, err := clearDb(ctx, namespace)
		flushed[namespace] = trimmedT{internal, cacheDb}

		if err != nil {
			return flushed, err
		}
	}

	return flushed, nil
}
