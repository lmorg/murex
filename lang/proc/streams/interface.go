package streams

import "io"

// This is the stream interface that's used for the shell functions streaming of data via standard in, out and err.
// It's written to be compatible with Go Reader and Writer interfaces however does expand upon then with additional
// helper methods  to enable easier writing of builtin shell functions.

type Io interface {
	MakeParent()
	UnmakeParent()

	Stats() (uint64, uint64)

	GetDataType() string
	SetDataType(string)
	DefaultDataType(bool)

	Read([]byte) (int, error)
	//ReaderFunc(callback func([]byte))
	ReadLine(callback func([]byte))
	ReadArray(callback func([]byte))
	ReadAll() []byte

	Write([]byte) (int, error)
	Writeln([]byte) (int, error)

	//ReadFrom(r io.Reader) (n int64, err error)
	WriteTo(io.Writer) (int64, error)

	Close()
}
