package stdio

import (
	"context"
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
	IsTTY() bool

	Read([]byte) (int, error)
	ReadLine(callback func([]byte)) error
	ReadArray(ctx context.Context, callback func([]byte)) error
	ReadArrayWithType(ctx context.Context, callback func([]byte, string)) error
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
