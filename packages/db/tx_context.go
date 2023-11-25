package db

import (
	"context"
	"time"
)

type txContext struct {
	parentCtx context.Context
	txKey     any
	txConn    Tx
}

var _ context.Context = txContext{}

func (t txContext) Deadline() (deadline time.Time, ok bool) {
	return t.parentCtx.Deadline()
}

func (t txContext) Done() <-chan struct{} {
	return t.parentCtx.Done()
}

func (t txContext) Err() error {
	return t.parentCtx.Err()
}

func (t txContext) Value(key any) any {
	if key == t.txKey {
		return t.txConn
	}

	return t.parentCtx.Value(key)
}

func (t txContext) Rollback() {
	t.txConn.Rollback(t.parentCtx)
}

func (t txContext) Commit() error {
	return t.txConn.Commit(t.parentCtx)
}

type Context interface {
	context.Context

	Rollback()
	Commit() error
}

type TxProvider interface {
	Begin(ctx context.Context) (Tx, error)
}

func WithTx(parentCtx context.Context, txKey any, txProvider TxProvider) (Context, error) {
	txConn, err := txProvider.Begin(parentCtx)
	if err != nil {
		return nil, err
	}

	return txContext{parentCtx, txKey, txConn}, nil
}
