package lang

import (
	"sync"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/expressions/functions"
	"github.com/lmorg/murex/lang/stdio"
)

type previewCacheT struct {
	mutex sync.Mutex
	raw   []string
	cache []cacheBytesT
	dt    []cacheDataTypeT
}

type cacheT struct {
	b   *cacheBytesT
	dt  *cacheDataTypeT
	tee cacheStreamT
	use bool
}

type cacheStreamT struct {
	stdout stdio.Io
	stderr stdio.Io
}

type cacheBytesT struct {
	stdout []byte
	stderr []byte
}

type cacheDataTypeT struct {
	stdout string
	stderr string
}

var previewCache = new(previewCacheT)

// compare returns index+1 for last match, 0 for no match, or -1 for no change
func (pc *previewCacheT) compare(s []string) int {
	if len(pc.raw) < len(s) {
		l := len(pc.raw)
		pc.grow(s)
		return l
	}

	for i := range pc.raw {
		if pc.raw[i] != s[i] {
			return i
		}
	}

	if len(pc.raw) > len(s) {
		pc.shrink(s)
		return len(s)
	}

	return -1
}

func (pc *previewCacheT) grow(s []string) {
	pc.raw = s

	cache := make([]cacheBytesT, len(s))
	copy(cache, pc.cache)
	pc.cache = cache

	dt := make([]cacheDataTypeT, len(s))
	copy(dt, pc.dt)
	pc.dt = dt
}

func (pc *previewCacheT) shrink(s []string) {
	pc.raw = s

	pc.cache = pc.cache[:len(s)]
	pc.dt = pc.dt[:len(s)]
}

func (pc *previewCacheT) compile(tree *[]functions.FunctionT, procs *[]Process) int {
	s := make([]string, len(*tree))
	for i := range s {
		s[i] = string((*tree)[i].Raw)
	}

	offset := pc.compare(s)
	switch offset {
	case -1:
		i := len(*procs) - 1
		(*procs)[i].cache = new(cacheT)
		(*procs)[i].cache.b = &pc.cache[i]
		(*procs)[i].cache.dt = &pc.dt[i]
		(*procs)[i].cache.use = true
		// we don't want to create any tee's because there's been no change to the pipeline
		return offset

	case 0:
		// do nothing. Basically we need to run the entire pipeline

	default:
		i := offset - 1
		(*procs)[i].cache = new(cacheT)
		(*procs)[i].cache.b = &pc.cache[i]
		(*procs)[i].cache.dt = &pc.dt[i]
		(*procs)[i].cache.use = true
	}

	for i := 0; i < offset; i++ {
		(*procs)[i].Stdout.Close()
		(*procs)[i].Stderr.Close()
		(*procs)[i].hasTerminatedM.Lock()
		(*procs)[i].hasTerminatedV = true
		(*procs)[i].hasTerminatedM.Unlock()
	}

	for i := offset; i < len(*procs); i++ {
		(*procs)[i].cache = new(cacheT)
		(*procs)[i].cache.b = &pc.cache[i]
		(*procs)[i].cache.dt = &pc.dt[i]
		(*procs)[i].Stdout, (*procs)[i].cache.tee.stdout = streams.NewTee((*procs)[i].Stdout)
		(*procs)[i].Stderr, (*procs)[i].cache.tee.stderr = streams.NewTee((*procs)[i].Stderr)
	}

	return offset
}
