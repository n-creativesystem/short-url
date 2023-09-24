package rdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/n-creativesystem/short-url/pkg/domain/tx"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
)

func BeginTx(ctx context.Context, opts *sql.TxOptions) (*ent.Tx, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	client := GetClient()
	if client == nil {
		panic("database does not support context-aware transactions")
	}
	return client.BeginTx(ctx, opts)
}

func begin(ctx context.Context, opt *tx.Options) (context.Context, bool, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	// 既にトランザクションが開始されているかつ
	// ネストしたトランザクションを許可していない場合は新たに開始しない
	if _, ok := getTx(ctx); ok {
		return ctx, false, nil
	}

	tx, err := BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return nil, false, err
	}
	ctx = ent.NewTxContext(ctx, tx)
	return ctx, true, nil
}

func GetExecutor(ctx context.Context) *ent.Client {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	if v, ok := getTx(ctx); ok {
		return v.Client()
	}
	return GetClient()
}

func getTx(ctx context.Context) (*ent.Tx, bool) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	v := ent.TxFromContext(ctx)
	ok := v != nil
	return v, ok
}

func transaction(ctx context.Context, fn func(ctx context.Context) error, opt *tx.Options) error {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	ctx, isTxStart, err := begin(ctx, opt)
	if err != nil {
		return err
	}
	err = fn(ctx)
	tx, ok := getTx(ctx)
	if ok {
		if isTxStart {
			if err == nil {
				return tx.Commit()
			} else {
				return errors.Join(err, tx.Rollback())
			}
		} else {
			// ネストしたトランザクションの場合は何もしない
			return nil
		}
	}
	return errors.New("Context is not set to the connection on which the transaction was initiated.")
}

var (
	_ tx.Beginner = NewBeginner
)

func NewBeginner(opts ...tx.OptionFunc) (tx.ContextBeginner, error) {
	opt := &tx.Options{}
	tx.OptionApply(opt, opts...)

	return &beginner{
		opt: opt,
	}, nil
}

type beginner struct {
	opt *tx.Options
}

func (b *beginner) BeginTx(ctx context.Context, fn func(ctx context.Context) error) error {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	return transaction(ctx, fn, b.opt)
}
