//go:build !dynamodb

package short

import (
	"github.com/n-creativesystem/short-url/pkg/domain/short"
)

func newShortImpl() *shortImpl {
	return &shortImpl{}
}

func NewRepository() short.Repository {
	return newShortImpl()
}
