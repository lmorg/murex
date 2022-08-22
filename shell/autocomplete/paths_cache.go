package autocomplete

import (
	"crypto/md5"
	"sync"
	"time"
)

type pathsCacheT struct {
	hash    map[string]pathsCacheItemT
	gcSleep int
	mutex   sync.Mutex
}

type pathsCacheItemT struct {
	time  time.Time
	paths []string
}

var pathsCache = NewPathsCache()

func NewPathsCache() *pathsCacheT {
	pc := new(pathsCacheT)
	pc.hash = make(map[string]pathsCacheItemT)
	pc.gcSleep = 5 // minutes
	go pc.garbageCollection()
	return pc
}

var (
	t = []byte{':', 't'}
	f = []byte{':', 'f'}
)

func (pc *pathsCacheT) CreateHash(exe string, filesToo bool, regexp bool) []byte {
	h := md5.New()

	_, err := h.Write([]byte(exe))
	if err != nil {
		return nil
	}

	if filesToo {
		_, err = h.Write(t)
	} else {
		_, err = h.Write(f)
	}
	if err != nil {
		return nil
	}

	if regexp {
		_, err = h.Write(t)
	} else {
		_, err = h.Write(f)
	}
	if err != nil {
		return nil
	}

	hash := h.Sum(nil)
	return hash
}

func (pc *pathsCacheT) Get(hash []byte) []string {
	s := string(hash)

	pc.mutex.Lock()

	cache := pc.hash[s]
	if cache.time.After(time.Now()) {
		s := cache.paths
		pc.mutex.Unlock()
		return s
	}

	pc.mutex.Unlock()
	return nil
}

func (pc *pathsCacheT) Set(hash []byte, paths []string) {
	pc.mutex.Lock()

	pc.hash[string(hash)] = pathsCacheItemT{
		time:  time.Now().Add(time.Duration(30) * time.Second),
		paths: paths,
	}

	pc.mutex.Unlock()
}

func (pc *pathsCacheT) garbageCollection() {
	for {
		pc.mutex.Lock()
		sleep := pc.gcSleep
		pc.mutex.Unlock()
		time.Sleep(time.Duration(sleep) * time.Minute)

		pc.mutex.Lock()
		for hash := range pc.hash {
			if pc.hash[hash].time.Before(time.Now()) {
				delete(pc.hash, hash)
			}
		}
		pc.mutex.Unlock()
	}
}
