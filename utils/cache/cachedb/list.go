package cachedb

import (
	"context"
	"fmt"
	"time"
)

const (
	sqlList           = `SELECT * FROM '%s' WHERE ttl > unixepoch();`
	sqlListNamespaces = `SELECT name FROM sqlite_schema WHERE type ='table' AND name NOT LIKE 'sqlite_%';`
)

type listT struct {
	Key   string
	Value string
	TTL   string
}

func List(ctx context.Context, namespace string) ([]listT, error) {
	db := dbConnect()
	defer db.Close()

	rows, err := db.QueryContext(ctx, fmt.Sprintf(sqlList, namespace))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		slice []listT
		key   string
		value string
		ttl   int64
	)

	for rows.Next() {
		err = rows.Scan(&key, &value, &ttl)
		if err != nil {
			return slice, err
		}
		slice = append(slice, listT{
			Key:   key,
			Value: value,
			TTL:   time.Unix(ttl, 0).Format(time.UnixDate),
		})
	}

	return slice, rows.Err()
}

func ListNamespaces() []string {
	db := dbConnect()
	defer db.Close()

	namespaces := []string{}

	if Disabled {
		return namespaces
	}

	rows, err := db.Query(sqlListNamespaces)
	if err != nil {
		dbFailed("listing namespaces", err)
		return namespaces
	}

	var s string
	for rows.Next() {
		err = rows.Scan(&s)
		if err != nil {
			dbFailed("reading namespaces", err)
			return namespaces
		}
		namespaces = append(namespaces, s)
	}

	if err = rows.Close(); err != nil {
		dbFailed("closing cache post read of namespaces", err)
		return namespaces
	}

	return namespaces
}
