package request

import (
	"crypto/rand"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortValid(t *testing.T) {
	// URL必須エラー
	v := GenerateShortURL{}
	err := v.Valid()
	require.NotNil(t, err)
	assert.Len(t, err.Errors, 1)
	assert.Equal(t, "url", err.Errors[0].Fields)
	assert.Equal(t, "URL is a required field.", err.Errors[0].Message)

	// URL形式エラー
	v = GenerateShortURL{
		URL: "http://localhost.%31",
	}
	err = v.Valid()
	assert.NotNil(t, err)
	assert.Len(t, err.Errors, 1)
	assert.Equal(t, "url", err.Errors[0].Fields)
	assert.Equal(t, "Please enter in URL format.", err.Errors[0].Message)

	// 桁数エラー
	buf := make([]byte, 256)
	_, _ = rand.Read(buf)
	v = GenerateShortURL{
		URL: "http://localhost",
		Key: hex.EncodeToString(buf),
	}
	err = v.Valid()
	assert.NotNil(t, err)
	assert.Len(t, err.Errors, 1)
	assert.Equal(t, "key", err.Errors[0].Fields)
	assert.Equal(t, "The key range is 1 to 255.", err.Errors[0].Message)

	// 複数エラー
	v = GenerateShortURL{
		URL: "http://localhost.%31",
		Key: hex.EncodeToString(buf),
	}
	err = v.Valid()
	assert.NotNil(t, err)
	assert.Len(t, err.Errors, 2)
	assert.Equal(t, "url", err.Errors[0].Fields)
	assert.Equal(t, "Please enter in URL format.", err.Errors[0].Message)
	assert.Equal(t, "key", err.Errors[1].Fields)
	assert.Equal(t, "The key range is 1 to 255.", err.Errors[1].Message)

	// エラーなし
	v = GenerateShortURL{
		URL: "http://localhost",
		Key: "",
	}
	err = v.Valid()
	assert.Nil(t, err)
}
