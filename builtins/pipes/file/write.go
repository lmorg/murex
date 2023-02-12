package file

import (
	"errors"
	"io"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils"
)

// Write is the io.Writer interface
func (f *File) Write(b []byte) (int, error) {
	if f.file == nil {
		return 0, errors.New("no file open")
	}

	f.mutex.Lock()
	dep := f.dependents
	f.mutex.Unlock()

	if dep < 1 {
		return 0, io.ErrClosedPipe
	}

	i, err := f.file.Write(b)

	f.mutex.Lock()
	f.bWritten += uint64(i)
	f.mutex.Unlock()

	return i, err
}

// Writeln is the io.Writeln interface
func (f *File) Writeln(b []byte) (int, error) {
	return f.Write(append(b, utils.NewLineByte...))
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (f *File) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(f, dataType)
}
