package generic

import (
	"fmt"
	"text/tabwriter"

	"github.com/lmorg/murex/lang/stdio"
)

type arrayWriter struct {
	tabwriter *tabwriter.Writer
}

func newArrayWriter(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriter{
		tabwriter: tabwriter.NewWriter(writer, twMinWidth, twTabWidth, twPadding, twPadChar, twFlags),
	}
	return w, nil
}

func (w *arrayWriter) Write(b []byte) error {
	_, err := fmt.Fprintln(w.tabwriter, string(b))
	return err
}

func (w *arrayWriter) WriteString(s string) error {
	_, err := fmt.Fprintln(w.tabwriter, s)
	return err
}

func (w *arrayWriter) Close() error {
	return w.tabwriter.Flush()
}
