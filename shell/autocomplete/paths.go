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
)

func matchDirs(s string) []string {
	return matchFilesystem(s, false)
}

func matchFilesAndDirs(s string) []string {
	return matchFilesystem(s, true)
}

func matchFilesystem(s string, filesToo bool) []string {
	var (
		once      []string
		recursive []string
		wg        sync.WaitGroup
	)

	wg.Add(1)

	timeout, err := lang.ShellProcess.Config.Get("shell", "recursive-timeout", types.Integer)
	if err != nil {
		timeout = 200
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(int64(timeout.(int)))*time.Millisecond)
	defer cancel()

	done := make(chan bool)
	go func() {
		recursive = matchRecursive(ctx, s, filesToo)
		done <- true
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
		return recursive

	case <-ctx.Done():
		wg.Wait() // make sure the once search has done. We might be working on a slow storage media
		return once
	}
}

func partialPath(s string) (path, partial string) {
	expanded := variables.Expand([]rune(s))
	split := strings.Split(string(expanded), consts.PathSlash)
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

func matchRecursive(ctx context.Context, s string, filesToo bool) (hierarchy []string) {
	s = variables.ExpandString(s)

	maxDepth, err := lang.ShellProcess.Config.Get("shell", "recursive-max-depth", types.Integer)
	if err != nil {
		maxDepth = 5
	}

	//expanded := variables.Expand([]rune(s))
	split := strings.Split(s, consts.PathSlash)
	path := strings.Join(split[:len(split)-1], consts.PathSlash)
	//partial = split[len(split)-1]

	if len(s) > 0 && s[0] == consts.PathSlash[0] {
		path = consts.PathSlash + path
	}

	walker := func(walkedPath string, info os.FileInfo, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err != nil {
			return nil
		}

		if !info.IsDir() && !filesToo {
			return nil
		}

		if info.Name()[0] == '.' {
			return nil
		}

		dirs := strings.Split(walkedPath, consts.PathSlash)

		if len(dirs)-len(split) > maxDepth.(int) {
			return filepath.SkipDir
		}

		if len(dirs) != 0 && len(dirs[len(dirs)-1]) == 0 {
			return nil
		}

		if len(split) == 1 {
			if (len(dirs)) > 1 && strings.HasPrefix(dirs[len(dirs)-2], ".") &&
				!strings.HasPrefix(s, ".") && !strings.HasPrefix(s, "..") {
				//panic(fmt.Sprint(dirs, len(split), split))
				return filepath.SkipDir
			}

		} else {
			if (len(dirs)) > 1 && strings.HasPrefix(dirs[len(dirs)-2], ".") && !strings.HasPrefix(dirs[len(dirs)-2], "..") &&
				!strings.HasPrefix(split[len(split)-1], ".") && !strings.HasPrefix(split[len(split)-1], "..") {
				//panic(fmt.Sprint(dirs, len(split), split))
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

	/*err :=*/
	filepath.Walk(pwd, walker)
	//if err != nil {
	//	lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	//}

	/*if path != consts.PathSlash {
		if len(s) < 3 { // TODO: there is a better way of doing this
			hierarchy = append(hierarchy, ".."[len(s):]+consts.PathSlash)
		}
	}*/

	return
}
