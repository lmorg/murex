package autocomplete

import (
	"crypto/md5"
	"strings"
	"time"
)

type dynamicCacheT struct {
	hash map[string]dynamicCacheItemT
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
	go dc.garbageCollection()
	return dc
}

func (dc dynamicCacheT) CreateHash(exe string, args []string, block []rune) []byte {
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
	cache := dc.hash[s]
	if cache.time.After(time.Now()) {
		return cache.stdout, cache.dataType
	}

	return nil, ""
}

func (dc *dynamicCacheT) Set(hash []byte, stdout []byte, dataType string, TTL int) {
	dc.hash[string(hash)] = dynamicCacheItemT{
		time:     time.Now().Add(time.Duration(TTL) * time.Second),
		stdout:   stdout,
		dataType: dataType,
	}
}

func (dc *dynamicCacheT) garbageCollection() {
	for {
		time.Sleep(5 * time.Minute)
		for hash := range dc.hash {
			if dc.hash[hash].time.Before(time.Now()) {
				delete(dc.hash, hash)
			}
		}
	}
}
