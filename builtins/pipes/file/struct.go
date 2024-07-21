package file

import (
	"context"
	"io"
	"os"
	"sync"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

// File Io interface
type File struct {
	mutex      sync.Mutex
	bWritten   uint64
	dependents int32
	file       *os.File
}

// Read is an empty method because file devices are write only
func (f *File) Read([]byte) (int, error) { return 0, io.EOF }

// ReadLine is an empty method because file devices are write only
func (f *File) ReadLine(func([]byte)) error { return nil }

// ReadArray is an empty method because file devices are write only
func (f *File) ReadArray(context.Context, func([]byte)) error { return nil }

// ReadArrayWithType is an empty method because file devices are write only
func (f *File) ReadArrayWithType(context.Context, func(interface{}, string)) error { return nil }

// ReadMap is an empty method because file devices are write only
func (f *File) ReadMap(*config.Config, func(*stdio.Map)) error { return nil }

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
