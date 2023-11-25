package db

import (
	"context"
)

type Tx interface {
	DB
	Rollback(ctx context.Context)
	Commit(ctx context.Context) error
}
