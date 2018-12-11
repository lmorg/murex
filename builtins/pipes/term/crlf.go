package term

import (
	"os"
	"sync"

	"github.com/lmorg/murex/utils"
)

// This function is just a way for readline to guarantee that it will always start on a new line.
// (it actually only works for murex builtins so i'll need to get clevererer with readline at some point)

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
	lf.char = '\n'
	lf.mutex.Unlock()
	if b != '\n' {
		os.Stderr.Write(utils.NewLineByte)
	}
}

// CrLf function to append a line feed character at the end of text piped to the terminal to aid readability.
var CrLf = appendCrLf{char: '\n'}
