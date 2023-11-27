package cachedb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	sqlCreateTable = `CREATE TABLE IF NOT EXISTS %s (key STRING PRIMARY KEY, value STRING, ttl DATETIME KEY);`

	sqlRead = `SELECT value FROM %s WHERE key == ? AND ttl > unixepoch();`

	sqlWrite = `INSERT OR REPLACE INTO %s (key, value, ttl) VALUES (?, ?, ?);`
)

var (
	disabled bool   = true
	Path     string = os.TempDir() + "/murex-temp-cache.db" // allows tests to run without contaminating regular cachedb
	db       *sql.DB
)

func dbConnect() *sql.DB {
	if os.Getenv("MUREX_DEV") == "true" {
		fmt.Printf("cache DB: %s\n", Path)
	}

	db, err := sql.Open(driverName, "file:"+Path)
	if err != nil {
		dbFailed("opening cache database", err)
		return nil
	}

	disabled = false
	return db
}

func CreateTable(namespace string) {
	if db == nil {
		db = dbConnect()
	}

	if disabled {
		return
	}

	_, err := db.Exec(fmt.Sprintf(sqlCreateTable, namespace))
	if err != nil {
		dbFailed("creating table "+namespace, err)
	}
}

func dbFailed(message string, err error) {
	os.Stderr.WriteString(fmt.Sprintf("Error %s: %s: '%s'\n", message, err.Error(), Path))

	if os.Getenv("MUREX_DEV") != "true" {
		os.Stderr.WriteString("!!! Disabling persistent cache !!!\n")
		disabled = true
	}
}

func CloseDb() {
	if db != nil {
		_ = db.Close()
	}
}

func Read(namespace string, key string, ptr any) bool {
	if disabled || ptr == nil {
		return false
	}

	rows, err := db.Query(fmt.Sprintf(sqlRead, namespace), key)
	if err != nil {
		dbFailed("querying cache in "+namespace, err)
		return false
	}

	ok := rows.Next()
	if !ok {
		return false
	}

	var s string
	err = rows.Scan(&s)
	if err != nil {
		dbFailed("reading cache in "+namespace, err)
		return false
	}

	if err = rows.Close(); err != nil {
		dbFailed("closing cache post read in "+namespace, err)
		return false
	}

	if len(s) == 0 { // nothing returned
		return false
	}

	if err := json.Unmarshal([]byte(s), ptr); err != nil {
		dbFailed("unmarshalling cache in "+namespace, err)
		return false
	}

	return true
}

func Write(namespace string, key string, value any, ttl time.Time) {
	if disabled || value == nil {
		return
	}

	b, err := json.Marshal(value)
	if err != nil {
		dbFailed("marshalling cache in "+namespace, err)
		return
	}

	_, err = db.Exec(fmt.Sprintf(sqlWrite, namespace), key, string(b), ttl.Unix())
	if err != nil {
		dbFailed("writing to cache in "+namespace, err)
		return
	}
}
