package unicode

import (
	"github.com/lmorg/murex/lang/stdio"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return nil
	}

	r := []rune(string(b))
	for _, i := range r {
		callback([]byte(string([]rune{i})))
	}

	return nil
}

type arrayWriter struct {
	writer stdio.Io
}

func newArrayWriter(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriter{writer: writer}
	return w, nil
}

func (w *arrayWriter) Write(b []byte) error {
	_, err := w.writer.Write(b)
	return err
}

func (w *arrayWriter) WriteString(s string) error {
	_, err := w.writer.Write([]byte(s))
	return err
}

func (w *arrayWriter) Close() error { return nil }
