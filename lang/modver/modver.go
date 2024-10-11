package modver

import (
	"sync"

	"github.com/lmorg/murex/utils/semver"
)

const Default = "6.0"

var (
	modver   = make(map[string]*semver.Version)
	mutex    sync.Mutex
	baseline *semver.Version
)

func init() {
	var err error
	baseline, err = semver.Parse(Default)
	if err != nil {
		panic(err.Error())
	}
}

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
