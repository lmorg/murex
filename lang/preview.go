package lang

import (
	"sync"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/expressions/functions"
	"github.com/lmorg/murex/lang/stdio"
)

type previewCacheT struct {
	mutex   sync.Mutex
	raw     []string
	cache   []*cacheT
	streams []cacheStreamT
}

type cacheStreamT struct {
	stdout stdio.Io
	stderr stdio.Io
}

type cacheT struct {
	stdout []byte
	stderr []byte
}

var previewCache = new(previewCacheT)

// compare returns index+1 for last match, 0 for no match, or -1 for no change
func (pc *previewCacheT) compare(s []string) int {
	if len(pc.raw) > len(s) {
		return len(s)
	}

	if len(pc.raw) < len(s) {
		return len(pc.raw)
	}

	for i := range pc.raw {
		if pc.raw[i] != s[i] {
			return i
		}
	}

	return -1
}

func (pc *previewCacheT) compile(tree *[]functions.FunctionT, procs *[]Process) *[]Process {
	s := make([]string, len(*tree))
	for i := range s {
		s[i] = string((*tree)[i].Raw)
	}

	comp := pc.compare(s)
	switch comp {
	case -1:
		(*procs)[len(pc.raw)-1].cache = previewCache.cache[len(pc.raw)-1]
		p := (*procs)[len(pc.raw):]
		return &p

	case 0:
		previewCache = new(previewCacheT)

	default:
		(*procs)[comp-1].cache = previewCache.cache[comp-1]
	}

	for i := comp; i < len(*tree); i++ {
		(*procs)[i].Stdout, pc.streams[i].stdout = streams.NewTee((*procs)[i].Stdout)
		(*procs)[i].Stderr, pc.streams[i].stderr = streams.NewTee((*procs)[i].Stderr)
	}

	return procs
}
