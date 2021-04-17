package stdio

import (
	"io"

	"github.com/lmorg/murex/config"
)

// Io is the interface that's used for the shell functions pipelining of data eg via standard in, out and err.
// It is written to be compatible with Go Reader and Writer interfaces however does expand upon them with additional
// helper methods to enable easier writing of builtin shell functions.
type Io interface {
	Stats() (uint64, uint64)

	GetDataType() string
	SetDataType(string)
	DefaultDataType(bool)
	IsTTY() bool

	Read([]byte) (int, error)
	ReadLine(callback func([]byte)) error
	ReadArray(callback func([]byte)) error
	ReadMap(*config.Config, func(string, string, bool)) error
	ReadAll() ([]byte, error)

	Write([]byte) (int, error)
	Writeln([]byte) (int, error)
	WriteArray(string) (ArrayWriter, error)

	//ReadFrom(r io.Reader) (n int64, err error)
	WriteTo(io.Writer) (int64, error)

	Open()
	Close()
	ForceClose()
}
