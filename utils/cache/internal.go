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

func (ic *internalCacheT) Read(key string, ptr unsafe.Pointer) bool {
	if disabled {
		return false
	}

	ic.mutex.Lock()
	v, ok := ic.cache[key]
	ic.mutex.Unlock()

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

func (ic *internalCacheT) Dump(ctx context.Context) any {
	if disabled {
		return nil
	}

	if ic == nil {
		return nil
	}

	var s []internalDumpT

	ic.mutex.Lock()

	for key, item := range ic.cache {
		select {
		case <-ctx.Done():
			ic.mutex.Unlock()
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

	ic.mutex.Unlock()

	return s
}

func (ic *internalCacheT) Write(key string, value *[]byte, ttl time.Time) {
	if disabled {
		return
	}

	if ttl.Before(time.Now().Add(59 * time.Minute)) {
		// lets not bother writing anything to cache.db if the TTL < 1 hr
		return
	}

	ic.mutex.Lock()
	ic.cache[key] = &cacheItemT{value, ttl}
	ic.mutex.Unlock()
}

func (ic *internalCacheT) Trim(ctx context.Context) []string {
	if disabled {
		return nil
	}

	var s []string

	ic.mutex.Lock()

	for key, item := range ic.cache {
		select {
		case <-ctx.Done():
			ic.mutex.Unlock()
			return s

		default:
			if item.ttl.Before(time.Now()) {
				delete(ic.cache, key)
				s = append(s, key)
			}
		}
	}

	ic.mutex.Unlock()

	return s
}

func (ic *internalCacheT) Clear(ctx context.Context) []string {
	if disabled {
		return nil
	}

	var s []string

	ic.mutex.Lock()

	for key := range ic.cache {
		select {
		case <-ctx.Done():
			ic.mutex.Unlock()
			return s

		default:
			delete(ic.cache, key)
			s = append(s, key)
		}
	}

	ic.mutex.Unlock()

	return s
}
