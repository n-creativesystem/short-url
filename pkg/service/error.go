package service

import (
	"github.com/friendsofgo/errors"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
)

var (
	ErrNotFound     = repository.ErrRecordNotFound
	ErrKeyDuplicate = errors.New("Cannot be used because a duplicate key is specified.")
)

func Wrap(err error, message string) error {
	if err == nil {
		return err
	}
	if errors.Is(err, repository.ErrRecordNotFound) {
		return ErrNotFound
	}
	return errors.Wrap(err, message)
}

type ClientError struct {
	err error
}

func NewClientError(err error) error {
	return &ClientError{err: err}
}

func (e *ClientError) Error() string {
	return e.err.Error()
}
