package compare

import "crypto/subtle"

func ConstantTimeCompare(x, y string) bool {
	return ConstantTimeCompareWithByte([]byte(x), []byte(y))
}

func ConstantTimeCompareWithByte(x, y []byte) bool {
	return subtle.ConstantTimeCompare(x, y) == 1
}
