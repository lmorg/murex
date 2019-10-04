package autocomplete

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/readline"
)

func matchDirs(s string, errCallback func(error), dtc *readline.DelayedTabContext) []string {
	return matchFilesystem(s, false, errCallback, dtc)
}

func matchFilesAndDirs(s string, errCallback func(error), dtc *readline.DelayedTabContext) []string {
	return matchFilesystem(s, true, errCallback, dtc)
}

func matchFilesystem(s string, filesToo bool, errCallback func(error), dtc *readline.DelayedTabContext) []string {
	// Is recursive search enabled?
	enabled, err := lang.ShellProcess.Config.Get("shell", "recursive-enabled", types.Boolean)
	if err != nil {
		enabled = false
	}

	// If not, fallback to the faster surface level scan
	if !enabled.(bool) {
		if filesToo {
			return matchFilesAndDirsOnce(s)
		}
		return matchDirsOnce(s)
	}

	// If so, get timeout and depth, then start the scans in parallel
	var (
		once      []string
		recursive []string
		wg        sync.WaitGroup
	)

	wg.Add(1)

	softTimeout, err := lang.ShellProcess.Config.Get("shell", "recursive-soft-timeout", types.Integer)
	if err != nil {
		softTimeout = 150
	}

	hardTimeout, err := lang.ShellProcess.Config.Get("shell", "recursive-hard-timeout", types.Integer)
	if err != nil {
		hardTimeout = 5000
	}

	softCtx, _ := context.WithTimeout(context.Background(), time.Duration(int64(softTimeout.(int)))*time.Millisecond)
	hardCtx, _ := context.WithTimeout(context.Background(), time.Duration(int64(hardTimeout.(int)))*time.Millisecond)
	//defer cancel()

	done := make(chan bool)
	go func() {
		recursive = matchRecursive(hardCtx, s, filesToo, dtc)
		select {
		case <-softCtx.Done():
			dtc.AppendSuggestions(recursive)
		case <-hardCtx.Done():
			dtc.AppendSuggestions(recursive)
		default:
			done <- true
		}
	}()

	go func() {
		if filesToo {
			once = matchFilesAndDirsOnce(s)
		} else {
			once = matchDirsOnce(s)
		}
		wg.Done()
	}()

	select {
	case <-done:
		// The surface search should have already been completed but lets wait
		// for it regardless because the last thing we need is a completely
		// avoidable race condition
		wg.Wait()
		return append(once, recursive...)

	case <-softCtx.Done():
		// Make sure the surface search has done. It should have, but we might
		// be working on impossibly slow storage media
		wg.Wait()
		return once

	case <-hardCtx.Done():
		// Make sure the surface search has done. It should have, but we might
		// be working on impossibly slow storage media
		wg.Wait()
		return once
	}
}

func partialPath(s string) (path, partial string) {
	expanded := variables.ExpandString(s)
	split := strings.Split(expanded, consts.PathSlash)
	path = strings.Join(split[:len(split)-1], consts.PathSlash)
	partial = split[len(split)-1]

	if len(s) > 0 && s[0] == consts.PathSlash[0] {
		path = consts.PathSlash + path
	}

	if path == "" {
		path = "."
	}
	return
}

func matchLocal(s string, includeColon bool) (items []string) {
	path, file := partialPath(s)
	exes := make(map[string]bool)
	listExes(path, exes)
	items = matchExes(file, exes, includeColon)
	return
}

func matchFilesAndDirsOnce(s string) (items []string) {
	s = variables.ExpandString(s)
	path, partial := partialPath(s)

	var item []string

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.Name()[0] == '.' && (len(partial) == 0 || partial[0] != '.') {
			// hide hidden files and directories unless you press dot / period.
			// (this behavior will also hide files and directories in Windows if
			// those file system objects are prefixed with a dot / period).
			continue
		}
		if f.IsDir() {
			item = append(item, f.Name()+consts.PathSlash)
		} else {
			item = append(item, f.Name())
		}
	}

	item = append(item, ".."+consts.PathSlash)

	for i := range item {
		if strings.HasPrefix(item[i], partial) {
			items = append(items, item[i][len(partial):])
		}
	}
	return
}

func matchRecursive(ctx context.Context, s string, filesToo bool, dtc *readline.DelayedTabContext) (hierarchy []string) {
	s = variables.ExpandString(s)

	maxDepth, err := lang.ShellProcess.Config.Get("shell", "recursive-max-depth", types.Integer)
	if err != nil {
		maxDepth = 5
	}

	split := strings.Split(s, consts.PathSlash)
	path := strings.Join(split[:len(split)-1], consts.PathSlash)
	partial := split[len(split)-1]

	if len(s) > 0 && s[0] == consts.PathSlash[0] {
		path = consts.PathSlash + path
	}

	walker := func(walkedPath string, info os.FileInfo, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-dtc.Context.Done():
			return dtc.Context.Err()
		default:
		}

		if err != nil {
			return nil
		}

		if !info.IsDir() && !filesToo {
			return nil
		}

		if info.Name()[0] == '.' && (len(partial) == 0 || partial[0] != '.') {
			return nil
		}

		dirs := strings.Split(walkedPath, consts.PathSlash)

		if len(dirs) == len(split) {
			return nil
		}

		if len(dirs)-len(split) > maxDepth.(int) {
			return filepath.SkipDir
		}

		if len(dirs) != 0 && len(dirs[len(dirs)-1]) == 0 {
			return nil
		}

		switch {
		case strings.HasSuffix(s, consts.PathSlash):
			if (len(dirs)) > 1 && strings.HasPrefix(dirs[len(dirs)-2], ".") {

				return filepath.SkipDir
			}

		case len(split) == 1:
			if (len(dirs)) > 1 && strings.HasPrefix(dirs[len(dirs)-2], ".") &&
				(!strings.HasPrefix(s, ".") || strings.HasPrefix(s, "..")) {

				return filepath.SkipDir
			}

		default:
			if (len(dirs)) > 1 && strings.HasPrefix(dirs[len(dirs)-2], ".") && !strings.HasPrefix(dirs[len(dirs)-2], "..") &&
				(!strings.HasPrefix(partial, ".") || strings.HasPrefix(partial, "..")) {

				return filepath.SkipDir
			}
		}

		if strings.HasPrefix(walkedPath, s) {
			if info.IsDir() {
				hierarchy = append(hierarchy, walkedPath[len(s):]+consts.PathSlash)
			} else {
				hierarchy = append(hierarchy, walkedPath[len(s):])
			}
		}

		return nil
	}

	var pwd string
	if path == "" {
		pwd = "./"
	} else {
		pwd = path
	}

	filepath.Walk(pwd, walker)
	/*err = filepath.Walk(pwd, walker)
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	}*/

	return
}
