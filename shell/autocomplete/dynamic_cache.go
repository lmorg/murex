package autocomplete

import (
	"sync"
	"time"

	"github.com/lmorg/murex/utils/cache"
)

type dynamicCacheT struct {
	hash    map[string]dynamicCacheItemT
	gcSleep int
	mutex   sync.Mutex
}

type dynamicCacheItemT struct {
	time     time.Time
	Stdout   []byte
	DataType string
}

var dynamicCache = NewDynamicCache()

func NewDynamicCache() *dynamicCacheT {
	dc := new(dynamicCacheT)
	dc.hash = make(map[string]dynamicCacheItemT)
	dc.gcSleep = 5 // minutes
	go dc.garbageCollection()
	return dc
}

func (dc *dynamicCacheT) Get(hash string) ([]byte, string) {
	dc.mutex.Lock()
	item := dc.hash[hash]
	dc.mutex.Unlock()

	if item.time.After(time.Now()) {
		return item.Stdout, item.DataType
	}

	if cache.Read(cache.AUTOCOMPLETE_DYNAMIC, hash, &item) {
		return item.Stdout, item.DataType
	}

	return nil, ""
}

func (dc *dynamicCacheT) Set(hash string, stdout []byte, dataType string, TTL int) {
	item := dynamicCacheItemT{
		time:     time.Now().Add(time.Duration(TTL) * time.Second),
		Stdout:   stdout,
		DataType: dataType,
	}

	dc.mutex.Lock()
	dc.hash[hash] = item
	dc.mutex.Unlock()

	if TTL >= 60*60*24 { // 1 day
		cache.Write(cache.AUTOCOMPLETE_DYNAMIC, hash, item, item.time)
	}
}

func (dc *dynamicCacheT) garbageCollection() {
	for {
		dc.mutex.Lock()
		sleep := dc.gcSleep
		dc.mutex.Unlock()
		time.Sleep(time.Duration(sleep) * time.Minute)

		dc.mutex.Lock()
		for hash := range dc.hash {
			if dc.hash[hash].time.Before(time.Now()) {
				delete(dc.hash, hash)
			}
		}
		dc.mutex.Unlock()
	}
}
