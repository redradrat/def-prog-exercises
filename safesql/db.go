package safesql

import (
	"context"
	"database/sql"
)

type DB struct {
	db *sql.DB
}

func (db *DB) QueryContext(ctx context.Context, query TrustedSQL, args ...any) (*Rows, error) {
	r, err := db.db.QueryContext(ctx, query.s, args...)
	return r, err
}

func (db *DB) ExecContext(ctx context.Context, query TrustedSQL, args ...any) (Result, error) {
	r, err := db.db.ExecContext(ctx, query.s, args...)
	return r, err
}

func Open(driverName string, dataSourceName string) (*DB, error) {
	d, err := sql.Open(driverName, dataSourceName)
	return &DB{db: d}, err
}

type (
	Rows   = sql.Rows
	Result = sql.Result
)
