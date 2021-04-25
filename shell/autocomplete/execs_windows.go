// +build windows

package autocomplete

import (
	"io/ioutil"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

// SplitPath takes a %PATH% (or similar) string and splits it into an array of paths
func SplitPath(envPath string) []string {
	split := strings.Split(envPath, ";")
	return split
}

func listExes(path string, exes map[string]bool) {
	var showExts bool

	v, err := lang.ShellProcess.Config.Get("shell", "extensions-enabled", types.Boolean)
	if err != nil {
		showExts = false
	} else {
		showExts = v.(bool)
	}

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		name := strings.ToLower(f.Name())

		if len(name) < 5 {
			continue
		}

		ext := name[len(name)-4:]

		if ext == ".exe" || ext == ".com" || ext == ".bat" || ext == ".cmd" || ext == ".scr" {
			if showExts {
				exes[name] = true
			} else {
				exes[name[:len(name)-4]] = true
			}
		}
	}
	return
}

func matchExes(s string, exes map[string]bool, includeColon bool) (items []string) {
	colon := ""
	if includeColon {
		colon = ":"
	}

	for name := range exes {
		lc := strings.ToLower(s)
		if strings.HasPrefix(strings.ToLower(name), lc) {
			switch {
			case isSpecialBuiltin(name):
				items = append(items, name[len(s):]+colon)
			case consts.NamedPipeProcName == name:
				// do nothing
			default:
				items = append(items, name[len(s):])
			}
		}
	}

	//sortColon(items, 0, len(items)-1)

	return
}
