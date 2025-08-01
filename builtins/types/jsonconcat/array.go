package jsonconcat

import (
	"context"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	cb := func(b []byte) {
		callback(b)
	}

	return parse(b, cb)
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	cb := func(b []byte) {
		callback(b, types.Json)
	}

	return parse(b, cb)
}

type arrayWriter struct {
	writer stdio.Io
}

func newArrayWriter(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriter{writer: writer}
	return w, nil
}

func (w *arrayWriter) Write(b []byte) (err error) {
	_, err = w.writer.Writeln(b)
	return
}

func (w *arrayWriter) WriteString(s string) (err error) {
	return w.Write([]byte(s))
}

func (w *arrayWriter) Close() error {
	return nil
}
