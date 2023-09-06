package sqlselect

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/lmorg/murex/builtins/core/open"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/humannumbers"
)

func loadTables(p *lang.Process, fromFile string, pipes, vars []string, parameters string, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings, confPrintHeadings bool, confDataType string) error {
	var (
		v      interface{}
		dt     string
		err    error
		tables []string
	)

	db, err := createDb()
	if err != nil {
		return err
	}

	switch {
	case len(pipes) > 0:
		dt = confDataType
		debug.Json("select pipes", pipes)
		debug.Log(fromFile, parameters)
		tables = pipes
		for i := range pipes {
			v, err = readPipe(p, pipes[i])
			if err != nil {
				return err
			}

			err = createTable(p, db, pipes[i], v, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings)
			if err != nil {
				return err
			}
		}

	case len(vars) > 0:
		dt = confDataType
		debug.Json("select vars", vars)
		debug.Log(fromFile, parameters)
		tables = vars
		for i := range vars {
			v, err = readVariable(p, vars[i])
			if err != nil {
				return err
			}

			err = createTable(p, db, vars[i], v, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings)
			if err != nil {
				return err
			}
		}

	default:
		v, dt, err = readFile(p, fromFile)
		if err != nil {
			return err
		}

		err = createTable(p, db, "main", v, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings)
		if err != nil {
			return err
		}
	}

	p.Stdout.SetDataType(dt)
	return runQuery(p, db, dt, tables, parameters, confPrintHeadings)
}

func readPipe(p *lang.Process, name string) (interface{}, error) {
	pipe, err := lang.GlobalPipes.Get(name)
	if err != nil {
		return nil, err
	}

	fork := p.Fork(0)
	fork.Process.Stdin = pipe

	dt := pipe.GetDataType()
	v, err := lang.UnmarshalData(fork.Process, dt)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal named pipe '%s': %s", name, err.Error())
	}

	return v, nil
}

func readVariable(p *lang.Process, name string) (interface{}, error) {
	s, err := p.Variables.GetString(name)
	if err != nil {
		return nil, err
	}
	dt := p.Variables.GetDataType(name)

	fork := p.Fork(lang.F_CREATE_STDIN)
	fork.Process.Stdin.SetDataType(dt)
	fork.Process.Stdin.Write([]byte(s))

	v, err := lang.UnmarshalData(fork.Process, dt)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal variable '%s': %s", name, err.Error())
	}

	return v, nil
}

func readFile(p *lang.Process, fromFile string) (interface{}, string, error) {
	var (
		v   interface{}
		dt  string
		err error
	)

	if p.IsMethod {
		dt = p.Stdin.GetDataType()

		v, err = lang.UnmarshalData(p, dt)
		if err != nil {
			return nil, "", fmt.Errorf("unable to unmarshal STDIN: %s", err.Error())
		}

	} else {
		closers, err := open.OpenFile(p, &fromFile, &dt)
		if err != nil {
			return nil, "", err
		}

		f, err := os.Open(fromFile)
		if err != nil {
			return nil, "", err
		}

		fork := p.Fork(0)
		fork.Process.Stdin = streams.NewReadCloser(f)

		v, err = lang.UnmarshalData(fork.Process, dt)
		if err != nil {
			return nil, "", fmt.Errorf("unable to unmarshal %s: %s", fromFile, err.Error())
		}

		err = open.CloseFiles(closers)
		if err != nil {
			return nil, "", err
		}
	}

	return v, dt, nil
}

func createTable(p *lang.Process, db *sql.DB, name string, v interface{}, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings bool) error {
	debug.Log("Creating table:", name)
	switch v := v.(type) {
	case [][]string:
		return createTable_SliceSliceString(p, db, name, v, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings)

	case []interface{}:
		table := make([][]string, len(v)+1)
		i := 1
		err := types.MapToTable(v, func(s []string) error {
			table[i] = s
			i++
			return nil
		})
		if err != nil {
			return err
		}
		return createTable_SliceSliceString(p, db, name, table, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings)

	default:
		return fmt.Errorf("unable to convert the following data structure into a table '%s': %T", name, v)
	}
}

func createTable_SliceSliceString(p *lang.Process, db *sql.DB, name string, v [][]string, confFailColMismatch, confMergeTrailingColumns, confTableIncHeadings bool) error {
	if len(v) == 0 {
		return fmt.Errorf("no table found")
	}

	var (
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
		tx, err = openTable(db, name, headings)
		if err != nil {
			return err
		}
		nRow = 1

	} else {
		headings = make([]string, len(v[0]))
		for i := range headings {
			headings[i] = humannumbers.ColumnLetter(i)
		}
		tx, err = openTable(db, name, headings)
		if err != nil {
			return err
		}

		slice := stringToInterfaceTrim(v[0], len(v))
		err = insertRecords(tx, name, slice)
		if err != nil {
			return fmt.Errorf("unable to insert headings into sqlite3: %s", err.Error())
		}
		nRow = 1
	}

	for ; nRow < len(v); nRow++ {
		if p.HasCancelled() {
			return fmt.Errorf("cancelled")
		}

		if len(v[nRow]) != len(headings) && confFailColMismatch {
			return fmt.Errorf("table rows contain a different number of columns to table headings\n%d: %s", nRow, v[nRow])
		}

		if confMergeTrailingColumns {
			slice := stringToInterfaceMerge(v[nRow], len(headings))
			err = insertRecords(tx, name, slice)
			if err != nil {
				return fmt.Errorf("%s\n%d: %s", err.Error(), nRow, v[nRow])
			}
		} else {
			slice := stringToInterfaceTrim(v[nRow], len(headings))
			err = insertRecords(tx, name, slice)
			if err != nil {
				return fmt.Errorf("%s\n%d: %s", err.Error(), nRow, v[nRow][:len(headings)-1])
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("unable to commit sqlite3 transaction: %s", err.Error())
	}

	return nil
}

func runQuery(p *lang.Process, db *sql.DB, dt string, tables []string, parameters string, confPrintHeadings bool) error {
	query := createQueryString(tables, parameters)
	debug.Log(query)

	rows, err := db.QueryContext(p.Context, query)
	if err != nil {
		return fmt.Errorf("cannot query table: %s\nSQL: %s", err.Error(), query)
	}

	r, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("cannot query rows: %s", err.Error())
	}

	var (
		table [][]string
		nRow  int
	)

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
