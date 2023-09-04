//go:build !cgo_sqlite3
// +build !cgo_sqlite3

package sqlselect

import (
	_ "modernc.org/sqlite"
)
