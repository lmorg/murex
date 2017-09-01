package streams

import (
	"errors"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"os"
	"sync"
)

// File Io interface
type File struct {
	mutex    sync.Mutex
	closed   bool
	bWritten uint64
	isParent bool
	file     *os.File
}

// Read is an empty method because file devices are write only
func (f *File) Read([]byte) (int, error) { return 0, io.EOF }

// ReadLine is an empty method because file devices are write only
func (f *File) ReadLine(func([]byte)) error { return nil }

// ReadArray is an empty method because file devices are write only
func (f *File) ReadArray(func([]byte)) error { return nil }

// ReadMap is an empty method because file devices are write only
func (f *File) ReadMap(*config.Config, func(string, string, bool)) error { return nil }

// ReadAll is an empty method because file devices are write only
func (f *File) ReadAll() ([]byte, error) { return []byte{}, nil }

// WriteTo is an empty method because file devices are write only
func (f *File) WriteTo(io.Writer) (int64, error) { return 0, io.EOF }

// GetDataType is an empty method because file devices are write only
func (f *File) GetDataType() string { return types.Null }

// SetDataType is an empty method because file devices are write only
func (f *File) SetDataType(string) {}

// DefaultDataType is an empty method because file devices are write only
func (f *File) DefaultDataType(bool) {}

// IsTTY returns false because the file writer is not a pseudo-TTY
func (f *File) IsTTY() bool { return false }

// MakePipe turns the stream.Io interface into a named pipe
func (f *File) MakePipe() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.isParent = true
}

// MakeParent prevents the stream.Io interface from being casually closed
func (f *File) MakeParent() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.isParent = true
}

// UnmakeParent allows the stream.Io interface to be closed
func (f *File) UnmakeParent() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.isParent = false
}

// Write is the io.Writer interface
func (f *File) Write(b []byte) (int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f == nil || f.file == nil {
		return 0, errors.New("No file open.")
	}

	if f.closed {
		return 0, io.ErrClosedPipe
	}

	f.mutex.Unlock()
	i, err := f.file.Write(b)

	f.mutex.Lock()
	f.bWritten += uint64(i)

	return i, err
}

// Writeln is the io.Writeln interface
func (f *File) Writeln(b []byte) (int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f == nil || f.file == nil {
		return 0, errors.New("No file open.")
	}

	if f.closed {
		return 0, io.ErrClosedPipe
	}

	f.mutex.Unlock()
	i, err := f.file.Write(append(b, utils.NewLineByte...))

	f.mutex.Lock()
	f.bWritten += uint64(i)

	return i, err
}

// Stats returns bytes written and read. As File is a write-only interface bytes read will always equal 0
func (f *File) Stats() (bytesWritten, bytesRead uint64) {
	f.mutex.Lock()
	bytesWritten = f.bWritten
	bytesRead = 0
	f.mutex.Unlock()
	return
}

// NewFile writer stream.Io pipe
func NewFile(name string) (f *File, err error) {
	f = new(File)
	f.file, err = os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return nil, err
	}
	return
}

// Close file writer
func (f *File) Close() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f.isParent {
		return
	}

	f.file.Close()
}
