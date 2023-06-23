package types

import (
	"errors"
	"strings"

	"golang.org/x/exp/constraints"
)

var (
	ErrInvalid = errors.New("Invalid value.")
)

type Required[T constraints.Ordered] interface {
	isRequired() bool
	Value() (T, error)
}

type String string

var (
	_ Required[string] = (*String)(nil)
)

func (v String) isRequired() bool {
	return strings.TrimSpace(string(v)) != ""
}

func (v String) Value() (string, error) {
	if !v.isRequired() {
		return "", ErrInvalid
	}
	return string(v), nil
}

type numeric[T constraints.Integer | constraints.Float] struct {
	value T
}

func (v numeric[T]) isRequired() bool {
	return v.value > 0
}

func (v numeric[T]) Value() (T, error) {
	if !v.isRequired() {
		return 0, ErrInvalid
	}
	return v.value, nil
}

func Numeric[T constraints.Integer | constraints.Float](value T) Required[T] {
	return numeric[T]{
		value: value,
	}
}
