package cachedb

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	sqlClearRead   = `SELECT key FROM %s;`
	sqlClearDelete = `DELETE FROM %s;`
)

func Clear(ctx context.Context, namespace string) ([]string, error) {
	opts := new(sql.TxOptions)
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, fmt.Sprintf(sqlClearRead, namespace))
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

	_, err = tx.ExecContext(ctx, fmt.Sprintf(sqlClearDelete, namespace))
	if err != nil {
		return slice, err
	}
	return slice, tx.Commit()
}
