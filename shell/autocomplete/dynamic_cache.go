package autocomplete

import (
	"strings"
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

func (dc *dynamicCacheT) CreateHash(exe string, args []string, block []rune) []byte {
	s := exe + " " + strings.Join(args, " ")
	return []byte(s)
}

func (dc *dynamicCacheT) Get(hash []byte) ([]byte, string) {
	sHash := string(hash)

	dc.mutex.Lock()
	item := dc.hash[sHash]
	dc.mutex.Unlock()

	if item.time.After(time.Now()) {
		return item.Stdout, item.DataType
	}

	if cache.Read(cache.AUTOCOMPLETE_DYNAMIC, sHash, &item) {
		return item.Stdout, item.DataType
	}

	return nil, ""
}

func (dc *dynamicCacheT) Set(hash []byte, stdout []byte, dataType string, TTL int) {
	sHash := string(hash)
	item := dynamicCacheItemT{
		time:     time.Now().Add(time.Duration(TTL) * time.Second),
		Stdout:   stdout,
		DataType: dataType,
	}

	dc.mutex.Lock()
	dc.hash[sHash] = item
	dc.mutex.Unlock()

	if TTL >= 60*60*24 { // 1 day
		cache.Write(cache.AUTOCOMPLETE_DYNAMIC, sHash, item, item.time)
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
