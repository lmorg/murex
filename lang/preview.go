package lang

import (
	"fmt"
	"strings"
	"sync"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/expressions/functions"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils/lists"
	"github.com/lmorg/murex/utils/parser"
)

var previewCache = new(previewCacheT)

type previewCacheT struct {
	mutex sync.Mutex
	err   error
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

func PreviewInit() {
	previewCache = new(previewCacheT)
}

const errPressF9 = "for your safety, press [f9] to confirm preview reload"

// compare returns:
// -1     == new / no match
// len(s) == all matched
func (pc *previewCacheT) compare(s []string) (int, error) {
	if len(pc.raw) > 0 {

		if len(s) != len(pc.raw) {
			return 0, fmt.Errorf("new commands added to the command line: %s", errPressF9)
		}

	}

	sLen, rLen := len(s), len(pc.raw)

	var i int
	for ; i < sLen && i < rLen; i++ {
		if s[i] != pc.raw[i] {
			break
		}
	}

	return i - 1, nil
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

func (pc *previewCacheT) compile(tree *[]functions.FunctionT, procs *[]Process) error {
	pc.mutex.Lock()
	defer pc.mutex.Unlock()

	if pc.err != nil {
		return pc.err
	}

	s := make([]string, len(*tree))
	for i := range s {
		s[i] = string((*tree)[i].Raw)
	}

	offset, err := pc.compare(s)
	if err != nil {
		pc.err = err
		return err
	}

	safe := parser.GetSafeCmds()
	for i := offset + 1; i < len(*tree); i++ {
		cmd := string((*tree)[i].CommandName())
		switch {
		case !strings.HasPrefix(s[i], cmd):
			return fmt.Errorf("a command executable has changed name: %s", errPressF9)
		case !lists.Match(safe, cmd):
			return fmt.Errorf("a command line change has been made prior to potentially unsafe commands: %s", errPressF9)
		}
	}

	pc.grow(s)

	for i := range *procs {
		(*procs)[i].cache = new(cacheT)
		(*procs)[i].cache.b = &pc.cache[i]
		(*procs)[i].cache.dt = &pc.dt[i]
		(*procs)[i].Stdout, (*procs)[i].cache.tee.stdout = streams.NewTee((*procs)[i].Stdout)
		(*procs)[i].Stderr, (*procs)[i].cache.tee.stderr = streams.NewTee((*procs)[i].Stderr)

		if i <= offset {
			(*procs)[i].cache.use = true
		}
	}

	return nil
}
