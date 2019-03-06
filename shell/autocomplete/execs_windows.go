// +build windows

package autocomplete

import (
	"io/ioutil"
	"sort"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

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
	colon := " "
	if includeColon {
		colon = ": "
	}

	for name := range exes {
		lc := strings.ToLower(s)
		if strings.HasPrefix(strings.ToLower(name), lc) {
			switch name {
			case ">", ">>", "[", "=":
				items = append(items, name[len(s):])
			case consts.NamedPipeProcName:
			default:
				items = append(items, name[len(s):])
			}
		}
	}
	sort.Strings(items)
	for i := range items {
		switch items[i] {
		case ">", ">>", "[", "=":
			items[i] += " "
		default:
			items[i] += colon
		}
	}
	return
}
