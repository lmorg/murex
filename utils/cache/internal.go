package cache

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var (
	cache    = make(map[string]*internalCacheT)
	disabled = true // avoid the cache for unit tests
)

type internalCacheT struct {
	mutex sync.Mutex
	cache map[string]*cacheItemT
}

type cacheItemT struct {
	value *[]byte
	ttl   time.Time
}

func (lc *internalCacheT) Read(key string, ptr unsafe.Pointer) bool {
	if disabled {
		return false
	}

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

type internalDumpT struct {
	Key   string
	Value string
	TTL   string
}

func (lc *internalCacheT) Dump(ctx context.Context) interface{} {
	if disabled {
		return nil
	}

	var s []internalDumpT

	lc.mutex.Lock()

	for key, item := range lc.cache {
		select {
		case <-ctx.Done():
			lc.mutex.Unlock()
			return nil

		default:
			if item.ttl.After(time.Now()) {
				s = append(s, internalDumpT{
					Key:   key,
					Value: string(*item.value),
					TTL:   item.ttl.Format(time.UnixDate),
				})
			}
		}
	}

	lc.mutex.Unlock()

	return s
}

func (lc *internalCacheT) Write(key string, value *[]byte, ttl time.Time) {
	if disabled {
		return
	}

	lc.mutex.Lock()
	lc.cache[key] = &cacheItemT{value, ttl}
	lc.mutex.Unlock()
}

func (lc *internalCacheT) Trim(ctx context.Context) []string {
	if disabled {
		return nil
	}

	var s []string

	lc.mutex.Lock()

	for key, item := range lc.cache {
		select {
		case <-ctx.Done():
			lc.mutex.Unlock()
			return s

		default:
			if item.ttl.Before(time.Now()) {
				delete(lc.cache, key)
				s = append(s, key)
			}
		}
	}

	lc.mutex.Unlock()

	return s
}

func (lc *internalCacheT) Flush(ctx context.Context) []string {
	if disabled {
		return nil
	}

	var s []string

	lc.mutex.Lock()

	for key := range lc.cache {
		select {
		case <-ctx.Done():
			lc.mutex.Unlock()
			return s

		default:
			delete(lc.cache, key)
			s = append(s, key)
		}
	}

	lc.mutex.Unlock()

	return s
}
