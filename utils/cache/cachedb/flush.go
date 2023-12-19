package cachedb

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	sqlFlushRead   = `SELECT key FROM %s;`
	sqlFlushDelete = `DELETE FROM %s;`
)

func Flush(ctx context.Context, namespace string) ([]string, error) {
	opts := new(sql.TxOptions)
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, fmt.Sprintf(sqlFlushRead, namespace))
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

	_, err = tx.ExecContext(ctx, fmt.Sprintf(sqlFlushDelete, namespace))
	if err != nil {
		return slice, err
	}
	return slice, tx.Commit()
}
