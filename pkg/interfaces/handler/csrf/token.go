// @refs: https://github.com/justinas/nosurf
package csrf

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
)

const tokenLength = 32

func xorToken(data, key []byte) {
	n := len(data)
	for i := 0; i < n; i++ {
		data[i] ^= key[i]
	}
}

func maskToken(data []byte) []byte {
	if len(data) != tokenLength {
		return nil
	}

	result := make([]byte, 2*tokenLength)
	key := result[:tokenLength]
	token := result[tokenLength:]
	copy(token, data)
	_, _ = rand.Read(key)

	xorToken(token, key)
	return result
}

func unmaskToken(data []byte) []byte {
	if len(data) != tokenLength*2 {
		return nil
	}

	key := data[:tokenLength]
	token := data[tokenLength:]
	xorToken(token, key)

	return token
}

func generateToken() []byte {
	bytes := make([]byte, tokenLength)
	_, _ = rand.Read(bytes)
	return bytes
}

func b64encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func b64decode(data string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil
	}
	return decoded
}

func verifyMasked(realToken, sentToken []byte) bool {
	sentPlain := unmaskToken(sentToken)
	return subtle.ConstantTimeCompare(realToken, sentPlain) == 1
}

func GenerateToken() (string, string) {
	token := generateToken()
	return b64encode(token), b64encode(maskToken(token))
}

func verifyToken(realToken, sentToken []byte) bool {
	realN := len(realToken)
	sentN := len(sentToken)

	if realN == tokenLength && sentN == 2*tokenLength {
		return verifyMasked(realToken, sentToken)
	} else {
		return false
	}
}

func VerifyToken(realToken, sentToken string) bool {
	decodeRealToken := b64decode(realToken)
	decodeSentToken := b64decode(sentToken)

	return verifyToken(decodeRealToken, decodeSentToken)
}
