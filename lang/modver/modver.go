package modver

import (
	"sync"

	"github.com/lmorg/murex/utils/semver"
)

var (
	modver   = make(map[string]*semver.Version)
	mutex    sync.Mutex
	baseline = &semver.Version{6, 0, 0}
)

func Set(module string, version *semver.Version) {
	mutex.Lock()
	modver[module] = version
	mutex.Unlock()
}

func Get(module string) *semver.Version {
	mutex.Lock()
	ver, ok := modver[module]
	mutex.Unlock()

	if ok {
		return ver
	}

	return baseline
}
