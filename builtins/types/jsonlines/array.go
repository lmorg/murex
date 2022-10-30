package jsonlines

import (
	"bufio"
	"bytes"
	"context"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(bytes.TrimSpace(scanner.Bytes()))
		}
	}

	return scanner.Err()
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func([]byte, string)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(bytes.TrimSpace(scanner.Bytes()), types.Json)
		}
	}

	return scanner.Err()
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
