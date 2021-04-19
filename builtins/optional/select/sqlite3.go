package sqlselect

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqlCreateTable = `CREATE TABLE IF NOT EXISTS main (%s);`

	sqlInsertRecord = `INSERT INTO main VALUES (%s);`

	sqlQueryWhere = `SELECT * FROM main WHERE %s;`

	sqlQueryNoWhere = `SELECT %s FROM main %s;`
)

func open(headings []string) (*sql.DB, *sql.Tx, error) {
	if len(headings) == 0 {
		return nil, nil, fmt.Errorf("Cannot create table: no titles supplied")
	}

	var sHeadings string
	for i := range headings {
		sHeadings += fmt.Sprintf(`"%s" NUMERIC,`, headings[i])
	}
	sHeadings = sHeadings[:len(sHeadings)-1]

	db, err := sql.Open("sqlite3", ":memory:" /*"file:debug.db"*/)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not open database: %s", err.Error())
	}

	query := fmt.Sprintf(sqlCreateTable, sHeadings)
	_, err = db.Exec(query)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not create table: %s\n%s", err.Error(), query)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not create transaction: %s", err.Error())
	}

	return db, tx, nil
}

func insertRecords(tx *sql.Tx, records []interface{}) error {
	if len(records) == 0 {
		return fmt.Errorf("No records to insert into transaction")
	}

	values, err := createValues(len(records))
	if err != nil {
		return fmt.Errorf("Cannot insert records into transaction: %s", err.Error())
	}

	_, err = tx.Exec(fmt.Sprintf(sqlInsertRecord, values), records...)
	if err != nil {
		return fmt.Errorf("Cannot insert records into transaction: %s", err.Error())
	}

	return nil
}

func createValues(length int) (string, error) {
	if length == 0 {
		return "", fmt.Errorf("No records to insert")
	}

	values := strings.Repeat("?,", length)
	values = values[:len(values)-1]

	return values, nil
}
