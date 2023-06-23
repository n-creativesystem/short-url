package filekms

import "github.com/google/tink/go/tink"

type masterKeyAEAD struct {
}

var (
	_ tink.AEAD = (*masterKeyAEAD)(nil)
)

func NewMasterKey() *masterKeyAEAD {
	return &masterKeyAEAD{}
}

func (x masterKeyAEAD) Encrypt(plaintext, associatedData []byte) ([]byte, error) {
	return plaintext, nil
}

func (x masterKeyAEAD) Decrypt(ciphertext, associatedData []byte) ([]byte, error) {
	return ciphertext, nil
}
