package cache

import (
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var cache = make(map[string]*localCacheT)

type localCacheT struct {
	mutex sync.Mutex
	cache map[string]*cacheItemT
}

type cacheItemT struct {
	value *[]byte
	ttl   time.Time
}

func (lc *localCacheT) Read(key string, ptr unsafe.Pointer) bool {
	lc.mutex.Lock()
	v, ok := lc.cache[key]
	lc.mutex.Unlock()

	if !ok {
		return false
	}

	if v.ttl.After(time.Now()) {
		return false
	}

	atomic.StorePointer(&ptr, unsafe.Pointer(v.value))
	return true
}

func (lc *localCacheT) Write(key string, value *[]byte, ttl time.Time) {
	lc.mutex.Lock()
	lc.cache[key] = &cacheItemT{value, ttl}
	lc.mutex.Unlock()
}
