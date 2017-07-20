package streams

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"os"
	"sync"
)

// This structure exists as a wrapper around os.Stdout and os.Stderr so they can be easily interchanged with this
// shells streams (which has a larger array of methods to enable easier writing of builtin shell functions.

type term struct {
	sync.Mutex
	//debug.Mutex
	bWritten uint64
	bRead    uint64
	lastChar byte
}

func (t *term) MakeParent()                                              {}
func (t *term) UnmakeParent()                                            {}
func (t *term) MakePipe()                                                {}
func (t *term) Read([]byte) (int, error)                                 { return 0, io.EOF }
func (t *term) ReadLine(func([]byte)) error                              { return nil }
func (t *term) ReadArray(func([]byte)) error                             { return nil }
func (t *term) ReadMap(*config.Config, func(string, string, bool)) error { return nil }
func (t *term) ReadAll() ([]byte, error)                                 { return []byte{}, nil }
func (t *term) WriteTo(io.Writer) (int64, error)                         { return 0, io.EOF }
func (t *term) GetDataType() string                                      { return types.Null }
func (t *term) SetDataType(string)                                       {}
func (t *term) DefaultDataType(bool)                                     {}

//func (t *term) Close()                           {}

func (t *term) Stats() (bytesWritten, bytesRead uint64) {
	t.Lock()
	bytesWritten = t.bWritten
	bytesRead = t.bRead
	t.Unlock()
	return
}

// Terminal: Standard Out

type TermOut struct {
	term
}

func (t *TermOut) Write(b []byte) (i int, err error) {
	t.Lock()
	t.bWritten += uint64(len(b))
	i, err = os.Stdout.Write(b)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	} else if len(b) > 0 {
		t.lastChar = b[len(b)-1]
	}
	t.Unlock()
	return
}

func (t *TermOut) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	return t.Write(line)
}

func (t *TermOut) Close() {
	// Since this is a terminal we'll append \n if none present (for better readability)
	if t.lastChar != '\n' && t.bWritten > 0 {
		t.Write(utils.NewLineByte)
	}
}

// Terminal: Standard Error

type TermErr struct {
	term
}

func (t *TermErr) Write(b []byte) (i int, err error) {
	t.Lock()
	t.bWritten += uint64(len(b))
	i, err = os.Stderr.Write(b)
	if err != nil {
		os.Stdout.WriteString(err.Error())
	} else if len(b) > 0 {
		t.lastChar = b[len(b)-1]
	}
	t.Unlock()
	return
}

func (t *TermErr) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	return t.Write(line)
}

func (t *TermErr) Close() {
	// Since this is a terminal we'll append \n if none present (for better readability)
	if t.lastChar != '\n' && t.bWritten > 0 {
		t.Write(utils.NewLineByte)
	}
}
