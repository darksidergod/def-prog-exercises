package safesql

import (
	"context"
	"database/sql"
)

type DB struct {
	db *sql.DB
}

type Rows = sql.Rows

func (db *DB) QueryContext(
	ctx context.Context,
	query TrustedSQL,
	args ...any) (*Rows, error) {
	r, err := db.db.QueryContext(ctx, query.s, args...)
	return r, err
}

type Result = sql.Result

func (db *DB) ExecContext(ctx context.Context, query TrustedSQL, args ...any) (*Result, error) {
	r, err := db.db.ExecContext(ctx, query.s, args...)
	return &r, err
}

func Open(driverName, dataSourceName string) (*DB, error) {
	d, err := sql.Open(driverName, dataSourceName)
	return &DB{d}, err
}
