package cd

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
	path     string
	fileInfo os.FileInfo
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
}

func cacheFileCompletions(pwd string) {
	mutex.Lock()
	cancel()

	if lastScan[pwd].Add(1 * time.Hour).After(time.Now()) {
		mutex.Unlock()
		return
	}

	var ctx context.Context

	maxDepth, _ := lang.ShellProcess.Config.Get("shell", "recursive-max-depth", types.Integer)

	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(60*time.Second))
	mutex.Unlock()

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

		if len(dirs) > maxDepth.(int) {
			return filepath.SkipDir
		}

		if len(dirs) != 0 && len(dirs[len(dirs)-1]) == 0 {
			return nil
		}

		cw = append(cw, cachedWalkT{walkedPath, info})
		return nil
	}

	filepath.Walk(pwd, walker)

	mutex.Lock()
	cachedWalk[pwd] = cw
	lastScan[pwd] = time.Now()
	mutex.Unlock()
}

func WalkCachedCompletions(pwd string, walker filepath.WalkFunc) {
	if len(pwd) == 0 {
		return
	}

	if pwd[0] != '/' {
		wd, err := os.Getwd()
		if err == nil {
			pwd = wd + "/" + pwd
		}
	}
	pwd = path.Clean(pwd)

	mutex.Lock()

	if lastScan[pwd].Add(1 * time.Hour).After(time.Now()) {
		mutex.Unlock()
		return
	}

	cw := cachedWalk[pwd]
	mutex.Unlock()

	var err error
	for _, file := range cw {
		err = walker(file.path, file.fileInfo, nil)
		if err != nil {
			return
		}
	}
}
