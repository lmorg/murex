package cachelib

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lmorg/murex/debug"
)

const (
	sqlCreateTable = `CREATE TABLE IF NOT EXISTS %s (key STRING PRIMARY KEY, value STRING, ttl DATETIME KEY);`

	sqlRead = `SELECT value FROM %s WHERE key == ? AND ttl > unixepoch();`

	sqlWrite = `INSERT OR REPLACE INTO %s (key, value, ttl) VALUES (?, ?, ?);`
)

var (
	disabled bool
	dbPath   string  = "deleteme.db" // TODO: change me
	db       *sql.DB = dbConnect()
)

func dbConnect() *sql.DB {
	db, err := sql.Open(driverName, "file:"+dbPath)
	if err != nil {
		dbFailed("cannot open cache database")
		return nil
	}

	return db
}

func CreateTable(namespace string) {
	_, err := db.Exec(fmt.Sprintf(sqlCreateTable, namespace))
	if err != nil {
		dbFailed(err.Error())
	}
}

func dbFailed(message string) {
	if debug.Enabled {
		panic(message)
	}
	os.Stderr.WriteString(fmt.Sprintf(
		"Error: %s: '%s'\nD!!! Disabling persistent cache !!!", message, dbPath,
	))
	disabled = true
}

func CloseDb() {
	err := db.Close()
	if err != nil {
		dbFailed(err.Error())
		return
	}
}

func Read(namespace string, key string, ptr any) bool {
	if disabled {
		return false
	}

	var s string

	rows, err := db.Query(fmt.Sprintf(sqlRead, namespace), key)
	if err != nil {
		dbFailed(err.Error())
		return false
	}

	ok := rows.Next()
	if !ok {
		return false
	}

	err = rows.Scan(&s)
	if err != nil {
		dbFailed(err.Error())
		return false
	}

	if err = rows.Close(); err != nil {
		dbFailed(err.Error())
		return false
	}

	if err = json.Unmarshal([]byte(s), ptr); err != nil {
		dbFailed(err.Error())
		return false
	}

	return true
}

func Write(namespace string, key string, value any, ttl time.Time) {
	if disabled {
		return
	}

	b, err := json.Marshal(value)
	if err != nil {
		dbFailed(err.Error())
		return
	}

	_, err = db.Exec(fmt.Sprintf(sqlWrite, namespace), key, string(b), ttl.Unix())
	if err != nil {
		dbFailed(err.Error())
		return
	}
}
