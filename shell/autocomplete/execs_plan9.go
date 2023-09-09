//go:build plan9
// +build plan9

package autocomplete

import (
	"io/ioutil"
	"strings"

	"github.com/lmorg/murex/utils/consts"
)

func listExes(path string, exes map[string]bool) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		// TODO: There is a log of logic missing in here that appears in the unix source
		exes[f.Name()] = true
	}
}

func matchExes(s string, exes map[string]bool) (items []string) {
	for name := range exes {
		if strings.HasPrefix(name, s) {
			if name != consts.NamedPipeProcName {
				items = append(items, name[len(s):])
			}
		}
	}

	return
}
