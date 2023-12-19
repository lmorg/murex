package cachedb

import (
	"context"
	"fmt"
	"time"
)

const (
	sqlList = `SELECT * FROM %s WHERE ttl > unixepoch();`
)

type listT struct {
	Key   string
	Value string
	TTL   string
}

func List(ctx context.Context, namespace string) ([]listT, error) {
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
