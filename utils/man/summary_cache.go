//go:build !windows
// +build !windows

package man

import (
	"sync"
	"time"

	"github.com/lmorg/murex/utils/cache"
)

var SummaryCache = NewSummaryCache()

type summaryCacheT struct {
	mutex sync.Mutex
	//mutex   debug.BadMutex
	summary map[string]string
}

func NewSummaryCache() *summaryCacheT {
	sc := new(summaryCacheT)
	sc.summary = make(map[string]string)
	return sc
}

func (sc *summaryCacheT) Get(path string) string {
	sc.mutex.Lock()

	s, ok := sc.summary[path]
	if !ok {
		cache.Read(cache.MAN_SUMMARY, path, &s)
	}

	sc.mutex.Unlock()
	return s
}

func (sc *summaryCacheT) Set(path, summary string) {
	sc.mutex.Lock()
	sc.summary[path] = summary
	cache.Write(cache.MAN_SUMMARY, path, summary, time.Now().Add(time.Hour*24*31))
	sc.mutex.Unlock()
}
