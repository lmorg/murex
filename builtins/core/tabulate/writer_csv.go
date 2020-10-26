package tabulate

import (
	"encoding/csv"
	"errors"

	"github.com/lmorg/murex/lang/proc/stdio"
)

type csvWriter struct {
	table       [][]string
	out         *csv.Writer
	err         error
	joiner      string
	columnWraps bool
}

func newCsvWriter(writer stdio.Io, joiner string, columnWraps bool) *csvWriter {
	return &csvWriter{
		out:         csv.NewWriter(writer),
		joiner:      joiner,
		columnWraps: columnWraps,
	}
}

func (cw *csvWriter) Write(array []string) error {
	if cw.columnWraps {
		cw.table = append(cw.table, array)
		return nil

	} else {
		return cw.out.Write(array)
	}
}

func (cw *csvWriter) Merge(last, s string) error {
	if !cw.columnWraps {
		return errors.New("Merge shouldn't be invoked with columnWrap. This is a murex bug, please raise an issue at https://github.com/lmorg/murex")
	}

	l := len(cw.table) - 1
	cw.table[l][len(cw.table[l])-1] += cw.joiner + s
	return nil
}

func (cw *csvWriter) Flush() {
	if cw.columnWraps {
		cw.err = cw.out.WriteAll(cw.table)
		cw.out.Flush()

	} else {
		cw.out.Flush()
	}
}

func (cw *csvWriter) Error() error {
	if cw.err != nil {
		return cw.err
	}
	return cw.out.Error()
}
