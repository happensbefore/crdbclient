package db

import "context"

type Rows interface {
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type DB interface {
	Exec(ctx context.Context, query string, queryArgs ...interface{}) error
	Query(ctx context.Context, query string, queryArgs ...interface{}) (Rows, error)
}

type Closer interface {
	Close()
}
