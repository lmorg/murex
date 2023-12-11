package cachedb

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	sqlTrimRead   = `SELECT key FROM %s WHERE ttl < unixepoch();`
	sqlTrimDelete = `DELETE FROM %s WHERE ttl < unixepoch();`
)

func Trim(ctx context.Context, namespace string) ([]string, error) {
	opts := new(sql.TxOptions)
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, fmt.Sprintf(sqlTrimRead, namespace))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		slice []string
		key   string
	)

	for rows.Next() {
		err = rows.Scan(&key)
		if err != nil {
			return slice, err
		}
		slice = append(slice, key)
	}

	_, err = tx.ExecContext(ctx, fmt.Sprintf(sqlTrimDelete, namespace))
	if err != nil {
		return slice, err
	}
	return slice, tx.Commit()
}
