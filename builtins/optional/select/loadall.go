package sqlselect

import (
	"database/sql"
	"fmt"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/humannumbers"
)

func loadAll(p *lang.Process, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings, confPrintHeadings bool) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	v, err := lang.UnmarshalData(p, dt)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal STDIN: %s", err.Error())
	}
	switch v.(type) {
	case [][]string:
		return sliceSliceString(p, v.([][]string), dt, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings, confPrintHeadings)

	/*case [][]interface{}:
	return sliceSliceInterface(p, v.([][]interface{}), dt, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings)*/

	default:
		return fmt.Errorf("Not a table") // TODO: better error message please
	}
}

func sliceSliceString(p *lang.Process, v [][]string, dt string, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings, confPrintHeadings bool) error {
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
		db, tx, err = open(headings)
		if err != nil {
			return err
		}
		nRow = 1

	} else {
		headings = make([]string, len(v[0]))
		for i := range headings {
			headings[i] = humannumbers.ColumnLetter(i)
		}
		db, tx, err = open(headings)
		if err != nil {
			return err
		}

		slice := stringToInterfaceTrim(v[0], len(v))
		err = insertRecords(tx, slice)
		nRow = 1
	}

	for ; nRow < len(v); nRow++ {
		if p.HasCancelled() {
			return nil
		}

		if len(v[nRow]) != len(headings) && confFailColMismatch {
			return fmt.Errorf("Table rows contain a different number of columns to table headings\n%d: %s", nRow, v[nRow])
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
		return err
	}

	query := createQueryString(p.Parameters.StringAll())
	debug.Log(query)

	rows, err := db.QueryContext(p.Context, query)
	//rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("Cannot query table: %s\nSQL: %s", err.Error(), query)
	}

	r, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("Cannot query rows: %s", err.Error())
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
		return fmt.Errorf("Cannot retrieve rows: %s", err.Error())
	}

	b, err := lang.MarshalData(p, dt, table)
	if err != nil {
		return fmt.Errorf("Unable to marshal STDOUT: %s", err.Error())
	}

	_, err = p.Stdout.Write(b)
	return err
}
