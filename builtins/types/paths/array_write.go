package paths

import (
	"github.com/lmorg/murex/lang/stdio"
)

// path

/*type arrayWriterPath struct {
	array  []string
	writer stdio.Io
}

func newArrayWriterPath(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriterPaths{writer: writer}
	return w, nil
}

func (w *arrayWriterPath) Write(b []byte) error {
	w.array = append(w.array, string(b))
	return nil
}

func (w *arrayWriterPath) WriteString(s string) error {
	w.array = append(w.array, s)
	return nil
}

func (w *arrayWriterPath) Close() error {
	b, err := marshalPaths(nil, w.array)
	if err != nil {
		return err
	}

	_, err = w.writer.Write(b)
	return err
}*/

// paths

type arrayWriterPaths struct {
	array  []string
	writer stdio.Io
}

func newArrayWriterPaths(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriterPaths{writer: writer}
	return w, nil
}

func (w *arrayWriterPaths) Write(b []byte) error {
	w.array = append(w.array, string(b))
	return nil
}

func (w *arrayWriterPaths) WriteString(s string) error {
	w.array = append(w.array, s)
	return nil
}

func (w *arrayWriterPaths) Close() error {
	b, err := marshalPaths(nil, w.array)
	if err != nil {
		return err
	}

	_, err = w.writer.Write(b)
	return err
}
