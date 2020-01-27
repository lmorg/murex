package generic

import (
	"bufio"

	"github.com/lmorg/murex/lang/proc/stdio"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(scanner.Bytes())
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

func (w *arrayWriter) Write(b []byte) error {
	_, err := w.writer.Writeln(b)
	return err
}

func (w *arrayWriter) WriteString(s string) error {
	_, err := w.writer.Writeln([]byte(s))
	return err
}

func (w *arrayWriter) Close() error { return nil }
