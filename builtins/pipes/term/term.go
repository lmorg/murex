package term

import (
	"context"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"

	"io"
	"sync"
)

// We don't register these pipes because we don't want users creating them
// adhoc inside murex
/*func init() {
	stdio.RegisterPipe("term-out", func(string) (stdio.Io, error) {
		return nil, errors.New("`term-out` is a system device and cannot be created")
	})

	stdio.RegisterPipe("term-err", func(string) (stdio.Io, error) {
		return nil, errors.New("`term-err` is a system device and cannot be created")
	})
}*/

// NewErr returns either Err or ErrRed depending on whether colourised output
// was defined via `colorise`
func NewErr(colourise bool) stdio.Io {
	if colourise {
		return new(ErrRed)
	}
	return new(Err)
}

// term structure exists as a wrapper around tty.Stdout and tty.Stderr so they
// can be easily interchanged with this shells streams (which has a larger
// array of methods to enable easier writing of builtin shell functions.
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
func (t *term) ReadArray(context.Context, func([]byte)) error { return nil }

// ReadArray is a null method because the term interface is write-only
func (t *term) ReadArrayWithType(context.Context, func(any, string)) error { return nil }

// ReadMap is a null method because the term interface is write-only
func (t *term) ReadMap(*config.Config, func(*stdio.Map)) error { return nil }

// ReadAll is a null method because the term interface is write-only
func (t *term) ReadAll() ([]byte, error) { return []byte{}, nil }

// WriteTo is a null method because the term interface is write-only
func (t *term) WriteTo(io.Writer) (int64, error) { return 0, io.EOF }

// GetDataType is a null method because the term interface is write-only
func (t *term) GetDataType() string { return types.Null }

// DefaultDataType is a null method because the term interface is write-only
func (t *term) DefaultDataType(bool) {}

// Open is a null method because the OS standard streams shouldn't be closed
// thus we don't need to track how many times they've been opened
func (t *term) Open() {}

// Close is a null method because the OS standard streams shouldn't be closed
func (t *term) Close() {}

// ForceClose is a null method because the OS standard streams shouldn't be closed
func (t *term) ForceClose() {}

// IsTTY always returns `true` because you are writing to a TTY. All over
// stream.Io interfaces should return `false`
func (t *term) IsTTY() bool { return true }

// Stats returns the bytes written and bytes read from the term interface
func (t *term) Stats() (bytesWritten, bytesRead uint64) {
	//t.mutex.RLock()
	t.mutex.Lock()
	bytesWritten = t.bWritten
	bytesRead = t.bRead
	//t.mutex.RUnlock()
	t.mutex.Unlock()
	return
}
