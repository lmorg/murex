//go:build windows
// +build windows

package autocomplete

import (
	"strings"

	"github.com/lmorg/murex/utils/consts"
)

// listExes called listExesWindows which exists in execs.go because it needs to
// be called when murex runs inside WSL
func listExes(path string, exes map[string]bool) {
	listExesWindows(path, exes)
}

func matchExes(s string, exes map[string]bool) (items []string) {
	for name := range exes {
		lc := strings.ToLower(s)
		if strings.HasPrefix(strings.ToLower(name), lc) {
			switch {
			case isSpecialBuiltin(name):
				items = append(items, name[len(s):])
			case consts.NamedPipeProcName == name:
				// do nothing
			default:
				items = append(items, name[len(s):])
			}
		}
	}

	return
}
