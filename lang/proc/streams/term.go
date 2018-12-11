package streams

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"

	"io"
	"os"
	"sync"
)

// NewTermErr returns either TermErr or TermErrRed depending on whether colourised output was defined via `colorise`
func NewTermErr(colourise bool) stdio.Io {
	if colourise {
		return new(TermErrRed)
	}
	return new(TermErr)
}

// term structure exists as a wrapper around os.Stdout and os.Stderr so they can be easily interchanged with this
// shells streams (which has a larger array of methods to enable easier writing of builtin shell functions.
type term struct {
	mutex    sync.Mutex
	bWritten uint64
	bRead    uint64
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

// Open is a null method because the OS standard streams shouldn't be closed
// thus we don't need to track how many times they've been opened
func (t *term) Open() {}

// Close is a null method because the OS standard streams shouldn't be closed
func (t *term) Close() {}

// ForceClose is a null method because the OS standard streams shouldn't be closed
func (t *term) ForceClose() {}

// IsTTY always returns `true` because you are writing to a TTY. All over stream.Io interfaces should return `false`.
func (t *term) IsTTY() bool { return true }

// MakePipe sets the isParent flag but probably isn't needed since terminals cannot be closed
func (t *term) MakePipe() {
	//t.MakeParent()
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
	t.mutex.Unlock()

	i, err = os.Stdout.Write(b)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	} else if len(b) > 0 {
		CrLf.set(b[len(b)-1])
	}

	return
}

// Writeln writes an OS-specific terminated line to the stdout
func (t *TermOut) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	return t.Write(line)
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *TermOut) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return writeArray(t, dataType)
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

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *TermErr) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return writeArray(t, dataType)
}

// Terminal: Standard Error - Coloured Red

// TermErrRed is the Stderr interface for term - with output coloured red
type TermErrRed struct {
	term
}

const (
	reset = "\x1b[0m"
	fgRed = "\x1b[31m"
)

// Write is the io.Writer() interface for term
func (t *TermErrRed) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	i, err = os.Stderr.WriteString(fgRed + string(b) + reset)
	if err != nil {
		os.Stdout.WriteString(fgRed + err.Error() + reset)
	} else if len(b) > 0 {
		CrLf.set(b[len(b)-1])
	}
	t.mutex.Unlock()
	return i - 9, err
}

// Writeln writes an OS-specific terminated line to the stderr
func (t *TermErrRed) Writeln(b []byte) (int, error) {
	line := append(b, utils.NewLineByte...)
	return t.Write(line)
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *TermErrRed) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return writeArray(t, dataType)
}
