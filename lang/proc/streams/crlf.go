package streams

import (
	"github.com/lmorg/murex/utils"
	"os"
	"sync"
)

// This function is just a way for readline to guarantee that it will always start on a new line.

type appendCrLf struct {
	mutex sync.Mutex
	char  byte
}

func (lf *appendCrLf) set(b byte) {
	lf.mutex.Lock()
	lf.char = b
	lf.mutex.Unlock()
}

func (lf *appendCrLf) Write() {
	lf.mutex.Lock()
	b := lf.char
	lf.mutex.Unlock()
	if b != '\n' {
		os.Stderr.Write(utils.NewLineByte)
	}
}

var CrLf appendCrLf
