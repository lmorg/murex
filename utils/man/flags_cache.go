//go:build !windows
// +build !windows

package man

import "sync"

type flagsCacheT struct {
	mutex sync.Mutex
	flags map[string]*flagsT
}

type flagsT struct {
	Flags        []string
	Descriptions map[string]string
}

func NewFlagsCache() *flagsCacheT {
	fc := new(flagsCacheT)
	fc.flags = make(map[string]*flagsT)
	return fc
}

func (fc *flagsCacheT) Get(cmd string) *flagsT {
	fc.mutex.Lock()
	flags, ok := fc.flags[cmd]
	fc.mutex.Unlock()
	if ok {
		return flags
	}
	return nil
}

func (fc *flagsCacheT) Set(cmd string, flags []string, descriptions map[string]string) {
	f := &flagsT{
		Flags:        flags,
		Descriptions: descriptions,
	}
	fc.mutex.Lock()
	fc.flags[cmd] = f
	fc.mutex.Unlock()
}

var Flags = NewFlagsCache()
