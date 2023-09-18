package secretbox

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"sync"

	"golang.org/x/crypto/nacl/secretbox"
)

var (
	// シングルトンにする
	secretBox        *SecretBox
	setSecretBoxOnce sync.Once
	secretBoxMu      sync.Mutex
)

func Init(encryptKey string) error {
	if encryptKey == "" {
		encryptKey = os.Getenv("ENCRYPT_KEY")
	}
	if encryptKey == "" {
		return errors.New("encryptKey is empty")
	}
	if v, err := NewSecretBox(encryptKey); err == nil {
		SetOnceSecretBox(v)
		return nil
	} else {
		return err
	}
}

func SetOnceSecretBox(v *SecretBox) {
	setSecretBoxOnce.Do(func() {
		secretBox = v
	})
}

func MustSecretBox(key string) {
	v, err := NewSecretBox(key)
	if err != nil {
		panic(err)
	}
	SetOnceSecretBox(v)
}

type SecretBox struct {
	key *[32]byte
}

func NewSecretBox(key string) (*SecretBox, error) {
	secretKeyBytes, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)
	return &SecretBox{
		key: &secretKey,
	}, nil
}

func (s *SecretBox) Encrypt(plainText string) (string, error) {
	var nonce [24]byte
	_, _ = rand.Read(nonce[:])
	buf := secretbox.Seal(nonce[:], []byte(plainText), &nonce, s.key)
	return hex.EncodeToString(buf), nil
}

func (s *SecretBox) Decrypt(cipherText string) (string, error) {
	sealed, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	var decryptNonce [24]byte
	copy(decryptNonce[:], sealed[:24])
	decrypted, ok := secretbox.Open(nil, sealed[24:], &decryptNonce, s.key)
	if !ok {
		return "", errors.New("decryption error")
	}
	return string(decrypted), nil
}

func Encrypt(plainText string) (string, error) {
	secretBoxMu.Lock()
	defer secretBoxMu.Unlock()
	return secretBox.Encrypt(plainText)
}

func Decrypt(cipherText string) (string, error) {
	secretBoxMu.Lock()
	defer secretBoxMu.Unlock()
	return secretBox.Decrypt(cipherText)
}

func MustEncrypt(plainText string) string {
	v, err := secretBox.Encrypt(plainText)
	if err != nil {
		panic(err)
	}
	return v
}

func MustDecrypt(cipherText string) string {
	v, err := secretBox.Decrypt(cipherText)
	if err != nil {
		panic(err)
	}
	return v
}
