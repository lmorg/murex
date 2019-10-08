package autocomplete

import (
	"sort"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func autoBranch(items *[]string) {
	// Is recursive search enabled?
	recursiveSearch, err := lang.ShellProcess.Config.Get("shell", "recursive-enabled", types.Boolean)
	if err != nil {
		recursiveSearch = false
	}

	if recursiveSearch.(bool) {
		sort.Sort(treeSorter(*items))
	} else {
		// Only show top level
		cropBranches(items)
		dedup(items)
	}
}

type treeSorter []string

func (ts treeSorter) Len() int      { return len(ts) }
func (ts treeSorter) Swap(i, j int) { ts[i], ts[j] = ts[j], ts[i] }

func (ts treeSorter) Less(i, j int) bool {
	iLen := abstractSize(ts[i])
	jLen := abstractSize(ts[j])

	if iLen == jLen {
		return ts[i] < ts[j]
	}

	return iLen < jLen
}

func abstractSize(s string) int {
	count := strings.Count(s, "/")
	switch {
	case count == 0:
		return 0
	case count == 1:
		if s[len(s)-1] != '/' {
			return 1
		}
		return 2
	default:
		return 3
	}
}

func cropBranches(tree *[]string) {
	for branch := range *tree {

		for i := 0; i < len((*tree)[branch])-1; i++ {
			if (*tree)[branch][i] == '/' {
				(*tree)[branch] = (*tree)[branch][:i+1]
				break
			}
		}

	}
}

// This is pretty inefficient. It really should be rewritten
func dedup(items *[]string) {
	m := make(map[string]bool)
	for i := range *items {
		m[(*items)[i]] = true
	}

	*items = make([]string, len(m))
	var i int
	for s := range m {
		(*items)[i] = s
		i++
	}

	sort.Strings(*items)
}
