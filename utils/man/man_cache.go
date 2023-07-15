//go:build !windows
// +build !windows

package man

import "sync"

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
	s := mc.paths[cmd]
	mc.mutex.Unlock()
	return s
}

func (sc *manCacheT) Set(cmd string, paths []string) {
	sc.mutex.Lock()
	sc.paths[cmd] = paths
	sc.mutex.Unlock()
}

var Paths = NewManCache()
