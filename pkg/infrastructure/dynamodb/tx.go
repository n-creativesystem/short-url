package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/n-creativesystem/short-url/pkg/domain/tx"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

type txKey struct{}

var (
	txK txKey
)

func GetExecutor(ctx context.Context) *dynamodb.Client {
	if v, ok := getTx(ctx); ok {
		return v
	}
	return GetDB()
}

func begin(ctx context.Context, opt *tx.Options) (context.Context, error) {
	// 既にトランザクションが開始されているかつ
	// ネストしたトランザクションを許可していない場合は新たに開始しない
	if _, ok := getTx(ctx); ok {
		return ctx, nil
	}
	tx := GetExecutor(ctx)
	return context.WithValue(ctx, txK, tx), nil
}
func getTx(ctx context.Context) (*dynamodb.Client, bool) {
	v, ok := ctx.Value(txK).(*dynamodb.Client)
	return v, ok
}

func transaction(ctx context.Context, fn func(ctx context.Context) error, opt *tx.Options) error {
	ctx, err := begin(ctx, opt)
	if err != nil {
		return err
	}
	return fn(ctx)
}

type txImpl struct {
	opt *tx.Options
}

var (
	_ tx.ContextBeginner = (*txImpl)(nil)
	_ tx.Beginner        = New
)

func New(opts ...tx.OptionFunc) (tx.ContextBeginner, error) {
	opt := &tx.Options{}
	tx.OptionApply(opt, opts...)
	return &txImpl{
		opt: opt,
	}, nil
}

func (impl *txImpl) BeginTx(ctx context.Context, fn func(ctx context.Context) error) error {
	logging.Default().Warn("`dynamodb` does not support transactions.")
	return transaction(ctx, fn, impl.opt)
}
