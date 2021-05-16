package file

import (
	"os"

	"github.com/lmorg/murex/lang/proc/stdio"
)

func init() {
	stdio.RegisterPipe("file", NewFile)
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
func NewFile(name string) (_ stdio.Io, err error) {
	f := new(File)
	f.file, err = os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return nil, err
	}
	f.dependants++
	return f, err
}

// Open file writer
func (f *File) Open() {
	f.mutex.Lock()
	f.dependants++
	f.mutex.Unlock()
}

// Close file writer
func (f *File) Close() {
	f.mutex.Lock()

	f.dependants--
	if f.dependants == 0 {
		f.file.Close()
	}

	if f.dependants < 0 {
		panic("more closed dependants than open")
	}

	f.mutex.Unlock()
}

// ForceClose forces the stream.Io interface to close.
func (f *File) ForceClose() {}
