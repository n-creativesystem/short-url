package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomString(t *testing.T) {
	length := 20
	randomString, err := GenerateRandomString(length)
	require.NoError(t, err)
	assert.Len(t, randomString, length*2)

	anotherRandomString, err := GenerateRandomString(length)
	require.NoError(t, err)
	assert.NotEqual(t, randomString, anotherRandomString)

	dummyRandFunc := func(b []byte) (int, error) {
		return 0, fmt.Errorf("dummy error")
	}

	RandFunc = dummyRandFunc

	_, err = GenerateRandomString(length)
	assert.Error(t, err)
}
