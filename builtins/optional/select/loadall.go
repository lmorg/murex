package sqlselect

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/lmorg/murex/builtins/core/open"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/humannumbers"
)

func loadAll(p *lang.Process, fromFile string, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings, confPrintHeadings bool, parameters string) error {
	var (
		dt  string
		v   interface{}
		err error
	)

	if p.IsMethod {
		dt = p.Stdin.GetDataType()

		v, err = lang.UnmarshalData(p, dt)
		if err != nil {
			return fmt.Errorf("unable to unmarshal STDIN: %s", err.Error())
		}

	} else {
		closers, err := open.OpenFile(p, &fromFile, &dt)
		if err != nil {
			return err
		}

		f, err := os.Open(fromFile)
		if err != nil {
			return err
		}

		fork := p.Fork(0)
		fork.Process.Stdin = streams.NewReadCloser(f)

		v, err = lang.UnmarshalData(fork.Process, dt)
		if err != nil {
			return fmt.Errorf("unable to unmarshal %s: %s", fromFile, err.Error())
		}

		err = open.CloseFiles(closers)
		if err != nil {
			return err
		}
	}

	p.Stdout.SetDataType(dt)

	switch v := v.(type) {
	case [][]string:
		return sliceSliceString(p, v, dt, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings, confPrintHeadings, parameters)

	/*case [][]interface{}:
	return sliceSliceInterface(p, v.([][]interface{}, dt, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings)*/

	default:
		return fmt.Errorf("unable to convert the following data structure into a table: %T", v)
	}
}

func sliceSliceString(p *lang.Process, v [][]string, dt string, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings, confPrintHeadings bool, parameters string) error {
	if len(v) == 0 {
		return fmt.Errorf("no table found")
	}

	var (
		db       *sql.DB
		tx       *sql.Tx
		err      error
		headings []string
		nRow     int
	)

	if confTableIncHeadings {
		headings = make([]string, len(v[0]))
		for i := range headings {
			headings[i] = fmt.Sprint(v[0][i])
		}
		db, tx, err = openDb(headings)
		if err != nil {
			return err
		}
		nRow = 1

	} else {
		headings = make([]string, len(v[0]))
		for i := range headings {
			headings[i] = humannumbers.ColumnLetter(i)
		}
		db, tx, err = openDb(headings)
		if err != nil {
			return err
		}

		slice := stringToInterfaceTrim(v[0], len(v))
		err = insertRecords(tx, slice)
		if err != nil {
			return fmt.Errorf("unable to insert headings into sqlite3: %s", err.Error())
		}
		nRow = 1
	}

	for ; nRow < len(v); nRow++ {
		if p.HasCancelled() {
			return nil
		}

		if len(v[nRow]) != len(headings) && confFailColMismatch {
			return fmt.Errorf("table rows contain a different number of columns to table headings\n%d: %s", nRow, v[nRow])
		}

		if confMergeTrailingColumns {
			slice := stringToInterfaceMerge(v[nRow], len(headings))
			err = insertRecords(tx, slice)
			if err != nil {
				return fmt.Errorf("%s\n%d: %s", err.Error(), nRow, v[nRow])
			}
		} else {
			slice := stringToInterfaceTrim(v[nRow], len(headings))
			err = insertRecords(tx, slice)
			if err != nil {
				return fmt.Errorf("%s\n%d: %s", err.Error(), nRow, v[nRow][:len(headings)-1])
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("unable to commit sqlite3 transaction: %s", err.Error())
	}

	query := createQueryString(parameters)
	debug.Log(query)

	rows, err := db.QueryContext(p.Context, query)
	if err != nil {
		return fmt.Errorf("cannot query table: %s\nSQL: %s", err.Error(), query)
	}

	r, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("cannot query rows: %s", err.Error())
	}

	var table [][]string
	nRow = 0

	if confPrintHeadings {
		table = [][]string{r}
		nRow++
	}

	for rows.Next() {
		table = append(table, make([]string, len(r)))
		slice := stringToInterfacePtr(&table[nRow], len(r))

		err = rows.Scan(slice...)
		if err != nil {
			return err
		}

		nRow++
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("cannot retrieve rows: %s", err.Error())
	}

	b, err := lang.MarshalData(p, dt, table)
	if err != nil {
		return fmt.Errorf("unable to marshal STDOUT: %s", err.Error())
	}

	_, err = p.Stdout.Write(b)
	return err
}
