package streams

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"os"
	"sync"
)

// term structure exists as a wrapper around os.Stdout and os.Stderr so they can be easily interchanged with this
// shells streams (which has a larger array of methods to enable easier writing of builtin shell functions.
type term struct {
	mutex sync.Mutex
	//debug.Mutex
	bWritten uint64
	bRead    uint64
	isParent bool
}

// Read is a null method because the term interface is write-only
func (t *term) Read([]byte) (int, error) { return 0, io.EOF }

// ReadLine is a null method because the term interface is write-only
func (t *term) ReadLine(func([]byte)) error { return nil }

// ReadArray is a null method because the term interface is write-only
func (t *term) ReadArray(func([]byte)) error { return nil }

// ReadMap is a null method because the term interface is write-only
func (t *term) ReadMap(*config.Config, func(string, string, bool)) error { return nil }

// ReadAll is a null method because the term interface is write-only
func (t *term) ReadAll() ([]byte, error) { return []byte{}, nil }

// WriteTo is a null method because the term interface is write-only
func (t *term) WriteTo(io.Writer) (int64, error) { return 0, io.EOF }

// GetDataType is a null method because the term interface is write-only
func (t *term) GetDataType() string { return types.Null }

// SetDataType is a null method because the term interface is write-only
func (t *term) SetDataType(string) {}

// DefaultDataType is a null method because the term interface is write-only
func (t *term) DefaultDataType(bool) {}

// Close is a null method because the OS standard streams shouldn't be closed
func (t *term) Close() {}

// IsTTY always returns `true` because you are writing to a TTY. All over stream.Io interfaces should return `false`.
func (t *term) IsTTY() bool { return true }

// MakeParent sets the isParent flag but probably isn't needed since terminals cannot be closed
func (t *term) MakeParent() {
	t.mutex.Lock()
	t.isParent = true
	t.mutex.Unlock()
}

// MakeParent unsets the isParent flag but probably isn't needed since terminals cannot be closed
func (t *term) UnmakeParent() {
	t.mutex.Lock()
	t.isParent = false
	t.mutex.Unlock()
}

// MakePipe sets the isParent flag but probably isn't needed since terminals cannot be closed
func (t *term) MakePipe() {
	t.MakeParent()
}

// Stats returns the bytes written and bytes read from the term interface
func (t *term) Stats() (bytesWritten, bytesRead uint64) {
	t.mutex.Lock()
	bytesWritten = t.bWritten
	bytesRead = t.bRead
	t.mutex.Unlock()
	return
}

// Terminal: Standard Out

// TermOut is the Stdout interface for term
type TermOut struct {
	term
}

// Write is the io.Writer() interface for term
func (t *TermOut) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	i, err = os.Stdout.Write(b)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	} else if len(b) > 0 {
		CrLf.set(b[len(b)-1])
	}
	t.mutex.Unlock()
	return
}

// Writeln writes an OS-specific terminated line to the stdout
func (t *TermOut) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	return t.Write(line)
}

// Terminal: Standard Error

// TermErr is the Stderr interface for term
type TermErr struct {
	term
}

// Write is the io.Writer() interface for term
func (t *TermErr) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	i, err = os.Stderr.Write(b)
	if err != nil {
		os.Stdout.WriteString(err.Error())
	} else if len(b) > 0 {
		CrLf.set(b[len(b)-1])
	}
	t.mutex.Unlock()
	return
}

// Writeln writes an OS-specific terminated line to the stderr
func (t *TermErr) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	return t.Write(line)
}
