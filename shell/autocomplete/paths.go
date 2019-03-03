package autocomplete

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils/consts"
)

func matchDirs(s string) []string {
	return matchRecursiveDirs(s)
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

func matchFilesAndDirs(s string) (items []string) {
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

func matchRecursiveDirs(s string) (hierarchy []string) {
	s = variables.ExpandString(s)

	//expanded := variables.Expand([]rune(s))
	split := strings.Split(s, consts.PathSlash)
	path := strings.Join(split[:len(split)-1], consts.PathSlash)
	//partial = split[len(split)-1]

	if len(s) > 0 && s[0] == consts.PathSlash[0] {
		path = consts.PathSlash + path
	}

	walker := func(walkedPath string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			return nil
		}

		if info.Name()[0] == '.' {
			return nil
		}

		dirs := strings.Split(walkedPath, consts.PathSlash)

		if len(dirs)-len(split) > 5 {
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
			hierarchy = append(hierarchy, walkedPath[len(s):]+consts.PathSlash)
		}

		return nil
	}

	var pwd string
	if path == "" {
		pwd = "./"
	} else {
		pwd = path
	}

	err := filepath.Walk(pwd, walker)
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	}

	/*if path != consts.PathSlash {
		if len(s) < 3 { // TODO: there is a better way of doing this
			hierarchy = append(hierarchy, ".."[len(s):]+consts.PathSlash)
		}
	}*/

	return
}
