package cachedb

import (
	"context"
	"fmt"
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
	rows, err := db.QueryContext(ctx, fmt.Sprintf(sqlTrimRead, namespace))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		slice []listT
		row   listT
	)

	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			return slice, err
		}
		slice = append(slice, row)
	}

	return slice, rows.Err()
}
