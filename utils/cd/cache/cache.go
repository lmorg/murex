package cache

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

type cachedWalkT struct {
	Path     string
	FileInfo os.FileInfo
}

var (
	cancel     = func() {}
	cachedWalk map[string][]cachedWalkT
	lastScan   map[string]time.Time
	mutex      sync.Mutex
)

func init() {
	cachedWalk = make(map[string][]cachedWalkT)
	lastScan = make(map[string]time.Time)

	go garbageCollection()
}

func garbageCollection() {
	for {
		time.Sleep(time.Duration(gcSleep) * time.Second)

		mutex.Lock()
		for s := range lastScan {
			if lastScan[s].Add(time.Duration(cacheTimeout) * time.Second).Before(time.Now()) {
				delete(lastScan, s)
				delete(cachedWalk, s)
			}
		}
		mutex.Unlock()
	}
}

func cleanPath(pwd string) string {
	if len(pwd) == 0 {
		return ""
	}

	if pwd[0] != '/' {
		wd, err := os.Getwd()
		if err == nil {
			pwd = wd + "/" + pwd
		}
	}
	return path.Clean(pwd)
}

func GatherFileCompletions(pwd string) {
	pwd = cleanPath(pwd)

	if len(pwd) == 0 {
		return
	}

	mutex.Lock()
	cancel()

	if lastScan[pwd].Add(time.Duration(cacheTimeout) * time.Second).After(time.Now()) {
		mutex.Unlock()
		return
	}

	var ctx context.Context

	maxDepth, err := lang.ShellProcess.Config.Get("shell", "recursive-max-depth", types.Integer)
	if err != nil {
		maxDepth = 0 // This should only crop up in testing
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(walkTimeout)*time.Second)
	mutex.Unlock()

	currentDepth := len(strings.Split(pwd, consts.PathSlash))
	var cw []cachedWalkT

	walker := func(walkedPath string, info os.FileInfo, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err != nil {
			return nil
		}

		dirs := strings.Split(walkedPath, consts.PathSlash)

		if len(dirs)-currentDepth > maxDepth.(int) {
			return filepath.SkipDir
		}

		/*if len(dirs) != 0 && len(dirs[len(dirs)-1]) == 0 {
			return nil
		}*/

		cw = append(cw, cachedWalkT{walkedPath, info})
		return nil
	}

	filepath.Walk(pwd, walker)

	mutex.Lock()
	cachedWalk[pwd] = cw
	lastScan[pwd] = time.Now()
	mutex.Unlock()
}

func WalkCompletions(pwd string, walker filepath.WalkFunc) bool {
	pwd = cleanPath(pwd)

	if len(pwd) == 0 {
		return false
	}

	mutex.Lock()

	if lastScan[pwd].Add(time.Duration(cacheTimeout) * time.Second).Before(time.Now()) {
		mutex.Unlock()
		return false
	}

	cw := cachedWalk[pwd]
	mutex.Unlock()

	var err error
	for _, file := range cw {
		err = walker(file.Path, file.FileInfo, nil)
		if err != nil {
			return false
		}
	}

	return true
}

func DumpCompletions() interface{} {
	type dumpT struct {
		Walk     []cachedWalkT
		LastScan string
	}

	dump := make(map[string]dumpT)

	mutex.Lock()
	for s := range cachedWalk {
		walk := make([]cachedWalkT, len(cachedWalk[s]))
		copy(walk, cachedWalk[s])
		dump[s] = dumpT{
			Walk:     walk,
			LastScan: lastScan[s].String(),
		}
	}
	mutex.Unlock()

	return dump
}
