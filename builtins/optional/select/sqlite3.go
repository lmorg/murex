package sqlselect

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
)

const (
	sqlCreateTable = `CREATE TABLE IF NOT EXISTS %s (%s);`

	sqlInsertRecord = `INSERT INTO %s VALUES (%s);`

	sqlQuery = `SELECT %s FROM %s %s %s;`
)

var (
	rxQuery      = regexp.MustCompile(`(?i)\s+(WHERE|GROUP BY|ORDER BY)\s+`)
	rxCheckFrom  = regexp.MustCompile(`(?iU)(\s+)?FROM\s+(\P{C})+($|\s+(WHERE|GROUP BY|ORDER BY)[\s]+)`)
	rxPipesMatch = regexp.MustCompile(`^(<[a-zA-Z0-9]+>[\s,]*)+$`)
	rxVarsMatch  = regexp.MustCompile(`^(\$[-_a-zA-Z0-9]+[\s,]*)+$`)
	rxPipesSplit = regexp.MustCompile(`[\s,]+`)
)

func createDb() (*sql.DB, error) {
	db, err := sql.Open(driverName, ":memory:" /*"file:debug.db"*/)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %s", err.Error())
	}

	return db, nil
}

func openTable(db *sql.DB, name string, headings []string) (*sql.Tx, error) {
	var err error

	if len(headings) == 0 {
		return nil, fmt.Errorf("cannot create table '%s': no titles supplied", name)
	}

	var sHeadings string
	for i := range headings {
		sHeadings += fmt.Sprintf(`"%s" NUMERIC,`, headings[i])
	}
	sHeadings = sHeadings[:len(sHeadings)-1]

	query := fmt.Sprintf(sqlCreateTable, name, sHeadings)
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("could not create table '%s': %s\n%s", name, err.Error(), query)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("could not create transaction: %s", err.Error())
	}

	return tx, nil
}

func insertRecords(tx *sql.Tx, name string, records []interface{}) error {
	if len(records) == 0 {
		return fmt.Errorf("no records to insert into transaction on table %s", name)
	}

	values, err := createValues(len(records))
	if err != nil {
		return fmt.Errorf("cannot insert records into transaction on table %s: %s", name, err.Error())
	}

	_, err = tx.Exec(fmt.Sprintf(sqlInsertRecord, name, values), records...)
	if err != nil {
		return fmt.Errorf("cannot insert records into transaction on table %s: %s", name, err.Error())
	}

	return nil
}

func createValues(length int) (string, error) {
	if length == 0 {
		return "", fmt.Errorf("no records to insert")
	}

	values := strings.Repeat("?,", length)
	values = values[:len(values)-1]

	return values, nil
}

func createQueryString(pipes []string, parameters string) string {
	split := rxQuery.Split(parameters, 2)
	match := rxQuery.FindString(parameters)

	switch len(split) {
	case 1:
		return fmt.Sprintf(sqlQuery, split[0], "main", match, "")

	case 2:
		if len(pipes) > 0 {
			return fmt.Sprintf(sqlQuery, split[0], strings.Join(pipes, ", "), match, split[1])
		}

		return fmt.Sprintf(sqlQuery, split[0], "main", match, split[1])

	default:
		panic("unexpected length of split")
	}
}
