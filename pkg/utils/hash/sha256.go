package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sum(value ...[]byte) string {
	h := sha256.New()
	for _, v := range value {
		h.Write(v)
	}
	return hex.EncodeToString(h.Sum(nil))
}
