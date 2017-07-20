package streams

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"os"
	"sync"
)

type File struct {
	mutex    sync.Mutex
	closed   bool
	bWritten uint64
	isParent bool
	file     *os.File
}

// These are null because file devices are write only
func (f *File) Read([]byte) (int, error)                                 { return 0, io.EOF }
func (f *File) ReadLine(func([]byte)) error                              { return nil }
func (f *File) ReadArray(func([]byte)) error                             { return nil }
func (f *File) ReadMap(*config.Config, func(string, string, bool)) error { return nil }
func (f *File) ReadAll() ([]byte, error)                                 { return []byte{}, nil }
func (f *File) WriteTo(io.Writer) (int64, error)                         { return 0, io.EOF }
func (f *File) GetDataType() string                                      { return types.Null }
func (f *File) SetDataType(string)                                       {}
func (f *File) DefaultDataType(bool)                                     {}

func (f *File) MakePipe() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.isParent = true
}

func (f *File) MakeParent() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.isParent = true
}

func (f *File) UnmakeParent() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.isParent = false
}

func (f *File) Write(b []byte) (int, error) {
	i, err := f.file.Write(b)

	f.mutex.Lock()
	f.bWritten += uint64(i)
	f.mutex.Unlock()

	return i, err
}

func (f *File) Writeln(b []byte) (int, error) {
	i, err := f.file.Write(append(b, utils.NewLineByte...))

	f.mutex.Lock()
	f.bWritten += uint64(i)
	f.mutex.Unlock()

	return i, err
}
func (f *File) Stats() (bytesWritten, bytesRead uint64) {
	f.mutex.Lock()
	bytesWritten = f.bWritten
	bytesRead = 0
	f.mutex.Unlock()
	return
}

func NewFile(name string) (f *File, err error) {
	f = new(File)
	f.file, err = os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return nil, err
	}
	return
}

func (f *File) Close() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f.isParent {
		return
	}

	f.file.Close()
}
