package noop

import (
	"context"

	"github.com/n-creativesystem/short-url/pkg/domain/tx"
)

var (
	_ tx.Beginner = NewBeginner
)

func NewBeginner(_ ...tx.OptionFunc) (tx.ContextBeginner, error) {
	return &beginner{}, nil
}

type beginner struct{}

func (*beginner) BeginTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}
