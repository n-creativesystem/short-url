package crypto

import (
	"encoding/base64"
	"sync"

	"github.com/google/tink/go/tink"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

var (
	aeadCipher tink.AEAD
	aeadMu     sync.Mutex
)

func Encrypt(plaintext string) (string, error) {
	aeadMu.Lock()
	defer aeadMu.Unlock()
	v, err := aeadCipher.Encrypt([]byte(plaintext), []byte{})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(v), nil
}

func MustEncrypt(plaintext string) string {
	v, err := Encrypt(plaintext)
	if err != nil {
		logging.Default().Warn(err)
		return ""
	}
	return v
}

func Decrypt(ciphertext string) (string, error) {
	buf, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	aeadMu.Lock()
	defer aeadMu.Unlock()
	v, err := aeadCipher.Decrypt(buf, []byte{})
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func MustDecrypt(ciphertext string) string {
	v, err := Decrypt(ciphertext)
	if err != nil {
		logging.Default().Warn(err)
		return ""
	}
	return string(v)
}
