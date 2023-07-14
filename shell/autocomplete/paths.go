package autocomplete

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils/cd/cache"
	"github.com/lmorg/murex/utils/consts"
)

func MatchDirectories(prefix string, act *AutoCompleteT) {
	act.append(matchDirs(prefix, act)...)
}

func matchDirs(s string, act *AutoCompleteT) []string {
	return matchFilesystem(s, false, "", act)
}

func matchFilesAndDirs(s string, act *AutoCompleteT) []string {
	return matchFilesystem(s, true, "", act)
}

func matchFilesAndDirsWithRegexp(s string, fileRegexp string, act *AutoCompleteT) []string {
	return matchFilesystem(s, true, fileRegexp, act)
}

func matchFilesystem(s string, filesToo bool, fileRegexp string, act *AutoCompleteT) []string {
	// compile regex
	var (
		rx  *regexp.Regexp
		err error
	)

	act.DoNotSort = true

	if len(fileRegexp) > 0 {
		rx, err = regexp.Compile(fileRegexp)
		if err != nil {
			act.ErrCallback(err)
		}
	}

	// Is recursive search enabled?
	enabled, _ := lang.ShellProcess.Config.Get("shell", "recursive-enabled", types.Boolean)
	//if err != nil {
	//	enabled = false
	//}

	// If not, fallback to the faster surface level scan
	if !enabled.(bool) {
		if filesToo {
			return matchFilesAndDirsOnce(s, rx)
		}
		return matchDirsOnce(s)
	}

	// If so, get timeout and depth, then start the scans in parallel
	var (
		once      []string
		recursive []string
	)

	softTimeout, _ := lang.ShellProcess.Config.Get("shell", "autocomplete-soft-timeout", types.Integer)
	hardTimeout, _ := lang.ShellProcess.Config.Get("shell", "autocomplete-hard-timeout", types.Integer)

	softCtx, _ := context.WithTimeout(context.Background(), time.Duration(int64(softTimeout.(int)))*time.Millisecond)
	hardCtx, _ := context.WithTimeout(context.Background(), time.Duration(int64(hardTimeout.(int)))*time.Millisecond)
	done := make(chan bool)

	act.largeMin() // assume recursive overruns

	go func() {
		recursive = matchRecursive(hardCtx, s, filesToo, rx, act)

		formatSuggestionsArray(act.ParsedTokens, recursive)
		act.DelayedTabContext.AppendSuggestions(recursive)
	}()

	go func() {
		if filesToo {
			once = matchFilesAndDirsOnce(s, rx)
		} else {
			once = matchDirsOnce(s)
		}
		done <- true
		select {
		case <-softCtx.Done():
			formatSuggestionsArray(act.ParsedTokens, once)
			act.DelayedTabContext.AppendSuggestions(once)
		default:
		}
	}()

	select {
	case <-done:
		return append(once, recursive...)

	case <-softCtx.Done():
		return []string{}
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

func matchFilesAndDirsOnce(s string, rx *regexp.Regexp) (items []string) {
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
		if rx != nil && !rx.MatchString(f.Name()) {
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

func matchRecursive(ctx context.Context, s string, filesToo bool, rx *regexp.Regexp, act *AutoCompleteT) (hierarchy []string) {
	s = variables.ExpandString(s)

	maxDepth, _ := lang.ShellProcess.Config.Get("shell", "recursive-max-depth", types.Integer)

	split := strings.Split(s, consts.PathSlash)
	path := strings.Join(split[:len(split)-1], consts.PathSlash)
	partial := split[len(split)-1]

	if len(s) > 0 && s[0] == consts.PathSlash[0] {
		path = consts.PathSlash + path
	}

	//var mutex sync.Mutex

	walker := func(walkedPath string, info os.FileInfo, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-act.DelayedTabContext.Context.Done():
			return act.DelayedTabContext.Context.Err()
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
			switch {
			case info.IsDir():
				//mutex.Lock()
				hierarchy = append(hierarchy, walkedPath[len(s):]+consts.PathSlash)
				//mutex.Unlock()
			case rx != nil && !rx.MatchString(info.Name()):
				return nil
			default:
				//mutex.Lock()
				hierarchy = append(hierarchy, walkedPath[len(s):])
				//mutex.Unlock()
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

	success := cache.WalkCompletions(pwd, walker)
	if !success {
		go cache.GatherFileCompletions(pwd)
		filepath.Walk(pwd, walker)
		return
	}

	go func() {
		filepath.Walk(pwd, walker)

		formatSuggestionsArray(act.ParsedTokens, hierarchy)
		act.DelayedTabContext.AppendSuggestions(hierarchy)
	}()

	/*err = filepath.Walk(pwd, walker)
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	}*/

	return
}
