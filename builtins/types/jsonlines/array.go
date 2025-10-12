package jsonlines

import (
	"bufio"
	"bytes"
	"context"
	"fmt"

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

	err := scanner.Err()
	if err != nil {
		return fmt.Errorf("error while reading a %s array: %s", types.JsonLines, err.Error())
	}
	return nil
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(bytes.TrimSpace(scanner.Bytes()), types.Json)
		}
	}

	err := scanner.Err()
	if err != nil {
		return fmt.Errorf("error while reading a %s array: %s", types.JsonLines, err.Error())
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
