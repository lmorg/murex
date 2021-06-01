package generic

import (
	"github.com/lmorg/murex/lang/stdio"
)

type arrayWriter struct {
	writer stdio.Io
}

func newArrayWriter(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriter{writer: writer}
	return w, nil
}

func (w *arrayWriter) Write(b []byte) error {
	_, err := w.writer.Writeln(b)
	return err
}

func (w *arrayWriter) WriteString(s string) error {
	_, err := w.writer.Writeln([]byte(s))
	return err
}

func (w *arrayWriter) Close() error { return nil }
