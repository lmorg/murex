package cachedb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/utils/consts"
)

const (
	sqlCreateTable = `CREATE TABLE IF NOT EXISTS '%s' (key STRING PRIMARY KEY, value STRING, ttl DATETIME KEY);`
	sqlRead        = `SELECT value FROM '%s' WHERE key == ? AND ttl > unixepoch();`
	sqlWrite       = `INSERT OR REPLACE INTO '%s' (key, value, ttl) VALUES (?, ?, ?);`
)

var (
	Disabled bool   = true
	path     string = os.TempDir() + "/murex-temp-cache.db" // allows tests to run without contaminating regular cachedb
)

func dbConnect() *sql.DB {
	db, err := sql.Open(driverName, fmt.Sprintf("file:%s?cache=shared", path))
	if err != nil {
		dbFailed("opening cache database", err)
		return nil
	}

	db.SetMaxOpenConns(1)

	Disabled = false
	return db
}

func CreateTable(namespace string) {
	db := dbConnect()
	defer db.Close()

	if Disabled {
		return
	}

	_, err := db.Exec(fmt.Sprintf(sqlCreateTable, namespace))
	if err != nil {
		dbFailed("creating table "+namespace, err)
	}
}

func dbFailed(message string, err error) {
	if debug.Enabled {
		//panic(fmt.Sprintf("%s: %s", message, err.Error()))
		os.Stderr.WriteString(fmt.Sprintf("Error %s: %s: '%s'\n%s\n", message, err.Error(), path, consts.IssueTrackerURL))
		os.Stderr.WriteString("!!! Disabling persistent cache !!!\n")
	}
	Disabled = true
}

func Read(namespace string, key string, ptr any) bool {
	db := dbConnect()
	defer db.Close()

	if Disabled || ptr == nil {
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
		dbFailed(fmt.Sprintf("unmarshalling cache in %s: %T (%s)", namespace, ptr, s), err)
		return false
	}

	return true
}

func Write(namespace string, key string, value any, ttl time.Time) {
	db := dbConnect()
	defer db.Close()

	if Disabled || value == nil {
		return
	}

	b, err := json.Marshal(value)
	if err != nil {
		dbFailed(fmt.Sprintf("marshalling cache in %s: %T (%v)", namespace, value, value), err)
		return
	}

	_, err = db.Exec(fmt.Sprintf(sqlWrite, namespace), key, string(b), ttl.Unix())
	if err != nil {
		dbFailed("writing to cache in "+namespace, err)
		return
	}
}

func SetPath(newPath string) {
	path = newPath
}

func GetPath() string {
	return path
}
