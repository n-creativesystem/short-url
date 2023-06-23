package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	// Wrap
	err := errors.New("Error")
	wrapErr := Wrap(err, "Wrap")
	assert.Error(t, wrapErr)
	assert.Equal(t, "Wrap: Error", wrapErr.Error())

	err = nil
	wrapErr = Wrap(err, "Wrap")
	assert.NoError(t, wrapErr)
	assert.Nil(t, wrapErr)

	err = NewClientError(errors.New("client error"))
	assert.Error(t, err)
	assert.Equal(t, "client error", err.Error())
}
