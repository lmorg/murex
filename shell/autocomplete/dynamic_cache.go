package autocomplete

import (
	"crypto/md5"
	"strings"
	"sync"
	"time"
)

type dynamicCacheT struct {
	hash    map[string]dynamicCacheItemT
	gcSleep int
	mutex   sync.Mutex
}

type dynamicCacheItemT struct {
	time     time.Time
	stdout   []byte
	dataType string
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
	h := md5.New()

	_, err := h.Write([]byte(exe))
	if err != nil {
		return nil
	}

	s := strings.Join(args, " ")
	_, err = h.Write([]byte(s))
	if err != nil {
		return nil
	}

	b := []byte(string(block))
	_, err = h.Write(b)
	if err != nil {
		return nil
	}

	hash := h.Sum(nil)
	return hash
}

func (dc *dynamicCacheT) Get(hash []byte) ([]byte, string) {
	s := string(hash)

	dc.mutex.Lock()

	cache := dc.hash[s]
	if cache.time.After(time.Now()) {
		b, s := cache.stdout, cache.dataType
		dc.mutex.Unlock()
		return b, s
	}

	dc.mutex.Unlock()
	return nil, ""
}

func (dc *dynamicCacheT) Set(hash []byte, stdout []byte, dataType string, TTL int) {
	dc.mutex.Lock()

	dc.hash[string(hash)] = dynamicCacheItemT{
		time:     time.Now().Add(time.Duration(TTL) * time.Second),
		stdout:   stdout,
		dataType: dataType,
	}

	dc.mutex.Unlock()
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
