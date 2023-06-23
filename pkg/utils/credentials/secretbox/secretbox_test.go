package secretbox

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecretBox(t *testing.T) {
	r := require.New(t)
	const key = "63285a53c8daf4aa0b3f3de46a1e6f4e89094b2f335286d1cc8e84391a77830d"
	MustSecretBox(key)
	plainText := "テストです"
	encrypt, err := Encrypt(plainText)
	r.NoError(err)
	decryptText, err := Decrypt(encrypt)
	r.NoError(err)
	r.Equal(plainText, decryptText)
}
