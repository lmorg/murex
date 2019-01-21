package string

import (
	"github.com/lmorg/murex/lang/proc/stdio"
)

type arrayWriter struct {
	writer stdio.Io
}

func newArrayWriter(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriter{writer: writer}
	return w, nil
}

func (w *arrayWriter) Write(_ []byte) error {
	return nil
}

func (w *arrayWriter) WriteString(_ string) error {
	return nil
}

func (w *arrayWriter) Close() error { return nil }
