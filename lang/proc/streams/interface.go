package streams

// This is the stream interface that's used for the shell functions streaming of data via standard in, out and err.
// It's written to be compatible with Go Reader and Writer interfaces however does expand upon then with additional
// helper methods  to enable easier writing of builtin shell functions.

type Io interface {
	MakeParent()
	UnmakeParent()

	Stats() (bytesWritten, bytesRead uint64)

	Read(p []byte) (i int, err error)
	ReadLine(line *string) (more bool)
	ReadData() (b []byte, more bool)
	ReaderFunc(callback func([]byte))
	ReadLineFunc(callback func([]byte))
	ReadAll() (b []byte)

	Write(p []byte) (int, error)
	Writeln(p []byte) (int, error)

	Close()
}
