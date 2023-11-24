package cachelib

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
	disabled bool
	dbPath   string  = os.TempDir() + "/murex-temp-cache.db" // allows tests to run without contaminating regular cachedb
	db       *sql.DB = dbConnect()
)

func init() {
	if os.Getenv("MUREX_DEV") == "true" {
		fmt.Printf("cache DB: %s\n", dbPath)
	}
}

func dbConnect() *sql.DB {
	db, err := sql.Open(driverName, "file:"+dbPath)
	if err != nil {
		dbFailed("opening cache database", err)
		return nil
	}

	return db
}

func CreateTable(namespace string) {
	_, err := db.Exec(fmt.Sprintf(sqlCreateTable, namespace))
	if err != nil {
		dbFailed("creating table "+namespace, err)
	}
}

func dbFailed(message string, err error) {
	os.Stderr.WriteString(fmt.Sprintf(
		"Error %s: %s: '%s'\n!!! Disabling persistent cache !!!", message, err.Error(), dbPath,
	))
	disabled = true
}

func CloseDb() {
	_ = db.Close()
}

func Read(namespace string, key string, ptr any) bool {
	if disabled {
		return false
	}

	var s string

	rows, err := db.Query(fmt.Sprintf(sqlRead, namespace), key)
	if err != nil {
		dbFailed("querying cache in "+namespace, err)
		return false
	}

	ok := rows.Next()
	if !ok {
		return false
	}

	err = rows.Scan(&s)
	if err != nil {
		dbFailed("reading cache in "+namespace, err)
		return false
	}

	if err = rows.Close(); err != nil {
		dbFailed("closing cache post read in "+namespace, err)
		return false
	}

	if err = json.Unmarshal([]byte(s), ptr); err != nil {
		dbFailed("unmarshalling cache in "+namespace, err)
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
		dbFailed("marshalling cache in "+namespace, err)
		return
	}

	_, err = db.Exec(fmt.Sprintf(sqlWrite, namespace), key, string(b), ttl.Unix())
	if err != nil {
		dbFailed("writing to cache in "+namespace, err)
		return
	}
}
