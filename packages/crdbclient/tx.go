package crdbclient

import (
	"context"

	"example/db"
	"github.com/jackc/pgx/v5"
)

type Tx struct {
	tx pgx.Tx
}

func (pt Tx) Rollback(ctx context.Context) {
	_ = pt.tx.Rollback(ctx)
}

func (pt Tx) Commit(ctx context.Context) error {
	return pt.tx.Commit(ctx)
}

func (pt Tx) Exec(ctx context.Context, query string, queryArgs ...interface{}) (err error) {
	_, err = pt.tx.Exec(ctx, query, queryArgs...)

	return
}

func (pt Tx) Query(ctx context.Context, query string, queryArgs ...interface{}) (db.Rows, error) {
	return pt.tx.Query(ctx, query, queryArgs...)
}

var _ db.DB = Tx{}
