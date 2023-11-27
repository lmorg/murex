package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"unsafe"

	"github.com/lmorg/murex/debug"
)

func initCache(namespace string) {
	cache[namespace] = new(localCacheT)
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
