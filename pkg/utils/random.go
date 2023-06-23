package utils

import (
	"crypto/rand"
	"encoding/hex"
)

type TRandFunc func([]byte) (int, error)

var (
	RandFunc TRandFunc = rand.Read
)

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := RandFunc(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
