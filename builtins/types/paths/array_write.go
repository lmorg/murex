package paths

/*
type arrayWriter struct {
	array  []string
	writer stdio.Io
}

func newArrayWriter(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriter{writer: writer}
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
	b, err := json.Marshal(w.array, w.writer.IsTTY())
	if err != nil {
		return err
	}

	_, err = w.writer.Write(b)
	return err
}
*/
