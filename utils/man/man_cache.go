//go:build !windows
// +build !windows

package man

import (
	"sync"

	"github.com/lmorg/murex/utils/cache"
)

type manCacheT struct {
	mutex sync.Mutex
	paths map[string][]string
}

func NewManCache() *manCacheT {
	mc := new(manCacheT)
	mc.paths = make(map[string][]string)
	return mc
}

func (mc *manCacheT) Get(cmd string) []string {
	mc.mutex.Lock()
	s, ok := mc.paths[cmd]
	mc.mutex.Unlock()

	if !ok {
		cache.Read(cache.MAN_PATHS, cmd, &s)
	}

	return s
}

func (sc *manCacheT) Set(cmd string, paths []string) {
	sc.mutex.Lock()
	sc.paths[cmd] = paths
	sc.mutex.Unlock()

	cache.Write(cache.MAN_PATHS, cmd, paths, cacheTtl())
}

var Paths = NewManCache()
