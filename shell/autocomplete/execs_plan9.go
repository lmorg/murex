// +build plan9

package autocomplete

import (
	"io/ioutil"
	"strings"
)

// SplitPath takes a $PATH (or similar) string and splits it into an array of paths
func SplitPath(envPath string) []string {
	split := strings.Split(envPath, ":")
	return split
}

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
