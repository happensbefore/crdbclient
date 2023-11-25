package main

import (
	"context"

	"example/db"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db db.DB
}

func (r *Repository) Save(ctx context.Context, data string) error {
	q := `INSERT INTO test_table (data) VALUES (@data)`

	return r.db.Exec(ctx, q, pgx.NamedArgs{"data": data})
}
