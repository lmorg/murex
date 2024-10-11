package modver

import (
	"sync"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/utils/semver"
)

const ModuleDefault = "6.0"

var (
	modver   = make(map[string]*semver.Version)
	mutex    sync.Mutex
	baseline = app.Semver()
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

func Dump() any {
	mutex.Lock()
	defer mutex.Unlock()

	dump := make(map[string]string)
	for mod, v := range modver {
		dump[mod] = v.String()
	}

	return dump
}
