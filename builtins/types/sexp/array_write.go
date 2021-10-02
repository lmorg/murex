package sexp

import (
	"github.com/abesto/sexp"
	"github.com/lmorg/murex/lang/stdio"
)

type arrayWriter struct {
	array     []string
	writer    stdio.Io
	canonical bool
}

func newArrayWriterC(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriter{
		writer:    writer,
		canonical: true,
	}
	return w, nil
}

func newArrayWriterS(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriter{
		writer:    writer,
		canonical: false,
	}
	return w, nil
}

func (w *arrayWriter) Write(b []byte) error {
	w.array = append(w.array, string(b))
	return nil
}

func (w *arrayWriter) WriteString(s string) error {
	w.array = append(w.array, s)
	return nil
}

func (w *arrayWriter) Close() error {
	b, err := sexp.Marshal(w.array, w.canonical)
	if err != nil {
		return err
	}

	_, err = w.writer.Write(b)
	return err
}
