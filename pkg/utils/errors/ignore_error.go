package errors

import "errors"

type IgnoreError struct {
	err error
}

var ()

func (e *IgnoreError) Error() string {
	return e.err.Error()
}

func (e *IgnoreError) Unwrap() error {
	return e.err
}

func NewIgnoreErrorWithErr(err error) error {
	return &IgnoreError{err: err}
}

func NewIgnoreError(msg string) error {
	return NewIgnoreErrorWithErr(errors.New(msg))
}
