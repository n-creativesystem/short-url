package csrf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	token := generateToken()
	assert.Len(t, token, tokenLength)
}

func TestVerifiesMaskedTokenCorrectly(t *testing.T) {
	realToken := []byte("qwertyuiopasdfghjklzxcvbnm123456")
	sentToken := []byte("qwertyuiopasdfghjklzxcvbnm123456" +
		"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
		"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	require.True(t, verifyToken(realToken, sentToken))
	realToken[0] = 'x'
	require.False(t, verifyToken(realToken, sentToken))
}

func TestMaskTokenAndVerify(t *testing.T) {
	token := generateToken()
	require.Len(t, token, tokenLength)
	maskToken := maskToken(token)
	require.Len(t, maskToken, 2*tokenLength)
	require.True(t, verifyToken(token, maskToken))
}

func TestGenerateTokenAndVerify(t *testing.T) {
	token, maskToken := GenerateToken()
	require.True(t, VerifyToken(token, maskToken))
}
