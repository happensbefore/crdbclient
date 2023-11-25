package crdbclient

import (
	"context"
	"fmt"
	"time"

	"example/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	db.DB
	db.Closer

	NewTxContext(parentCtx context.Context) (db.Context, error)
}

var _txKey = struct{}{}

type client struct {
	pool *pgxpool.Pool

	timeout time.Duration
}

func New(ctx context.Context, cfg *Config) (Client, error) {
	conf, err := pgxpool.ParseConfig(cfg.ToConnString())

	conf.AfterConnect = func(ctx context.Context, p *pgx.Conn) error {
		_, err = p.Exec(ctx, "SET TIMEZONE TO 'UTC'")
		return err
	}

	conf.MaxConnLifetime = cfg.MaxConnectionLifeTime
	conf.MaxConnIdleTime = cfg.MaxConnectionIdleTime
	conf.MaxConns = int32(cfg.MaxConnections)

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("[NewConn] can't create db conn: %w", err)
	}

	return &client{pool, cfg.QueryTimeout}, nil
}

func (c *client) NewTxContext(parentCtx context.Context) (db.Context, error) {
	return db.WithTx(parentCtx, _txKey, c)
}

func (c *client) Begin(ctx context.Context) (db.Tx, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return Tx{tx}, nil
}

func (c *client) Exec(ctx context.Context, sql string, args ...interface{}) (err error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	ctxConn, ok := ctx.Value(_txKey).(db.DB)
	if ok {
		return ctxConn.Exec(ctx, sql, args...)
	}

	_, err = c.pool.Exec(ctx, sql, args...)

	return
}

func (c *client) Query(ctx context.Context, sql string, args ...interface{}) (db.Rows, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	ctxConn, ok := ctx.Value(_txKey).(db.DB)
	if ok {
		return ctxConn.Query(ctx, sql, args...)
	}

	return c.pool.Query(ctx, sql, args...)
}

func (c *client) Close() {
	if c.pool != nil {
		c.pool.Close()
	}
}
