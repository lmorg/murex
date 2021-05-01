package sqlselect

import (
	"bytes"
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

var rxWhitespace = regexp.MustCompile(`[\s\t]+`)

// loadSeq is highly experimental!
func loadSeq(p *lang.Process, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings bool) error {
	p.Stdout.SetDataType(types.Generic)

	var (
		db       *sql.DB
		tx       *sql.Tx
		headings []string
		err      error
		innerErr error
		nRow     int
	)

	dt := p.Stdin.GetDataType()

	// TODO: confTableIncHeadings

	err = p.Stdin.ReadArray(func(b []byte) {
		b = bytes.TrimSpace(b)
		switch dt {
		case types.Generic, types.String, "generic", "string":
			row := string(b)
			records := rxWhitespace.Split(row, -1)

			if nRow == 0 {
				headings = records
				db, tx, innerErr = open(headings)
				if innerErr != nil {
					p.Done()
				}
				nRow++
				return
			}

			slice := make([]interface{}, len(headings))

			if len(records) != len(headings) {
				if confFailColMismatch {
					innerErr = fmt.Errorf("Table rows contain a different number of columns to table headings\n%d: %s", nRow, records)
					p.Done()
					return
				} //else {

				for i := len(headings); i >= len(records); i-- {
					slice[i-1] = ""
				}
				//}
			}

			for i := 0; i < len(headings); i++ {
				slice[i] = records[i]
			}

			err = insertRecords(tx, slice)
			if err != nil {
				innerErr = fmt.Errorf("%s\n%d: %s", err.Error(), nRow, records)
				p.Done()
			}

		default:
			panic("todo")
		}

		nRow++
		// todo transaction if i>highNumber
	})

	if innerErr != nil {
		return innerErr
	}
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	query := createQueryString(p.Parameters.StringAll())
	debug.Log(query)

	//rows, err := db.QueryContext(p.Context, query)
	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("Cannot query table: %s", err.Error())
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	err = aw.WriteString(strings.Join(headings, " , "))
	if err != nil {
		return err
	}

	records := make([]string, len(headings))    // TODO: should use rows.Columns instead of headings
	slice := make([]interface{}, len(headings)) // TODO: should use rows.Columns instead of headings
	for i := range slice {
		slice[i] = &records[i]
	}
	for rows.Next() {
		err = rows.Scan(slice...)
		if err != nil {
			return err
		}

		err = aw.WriteString(strings.Join(records, " , "))
		if err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("Cannot retrieve rows: %s", err.Error())
	}

	err = db.Close()
	if err != nil {
		return fmt.Errorf("Cannot close db: %s", err.Error())
	}

	return aw.Close()
}
