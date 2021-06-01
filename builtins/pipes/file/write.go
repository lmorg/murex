package file

import (
	"errors"
	"io"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils"
)

// Write is the io.Writer interface
func (f *File) Write(b []byte) (int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f == nil || f.file == nil {
		return 0, errors.New("No file open")
	}

	if f.dependants < 1 {
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
		return 0, errors.New("No file open")
	}

	if f.dependants < 1 {
		return 0, io.ErrClosedPipe
	}

	f.mutex.Unlock()
	i, err := f.file.Write(append(b, utils.NewLineByte...))

	f.mutex.Lock()
	f.bWritten += uint64(i)

	return i, err
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (f *File) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(f, dataType)
}
