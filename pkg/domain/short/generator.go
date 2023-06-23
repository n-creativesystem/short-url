package short

import (
	"strings"

	"github.com/jaevor/go-nanoid"
)

type Generator func(length ...int) string

var (
	generateKey        Generator
	defaultGenerateKey = NanoIdGenerateKey()
)

func GenerateKey(length ...int) string {
	if generateKey == nil {
		generateKey = defaultGenerateKey
	}
	return generateKey(length...)
}

func NanoIdGenerateKey() Generator {
	var (
		alphabet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numeric      = "0123456789"
		alphaNumeric = strings.ToLower(alphabet) + strings.ToUpper(alphabet) + numeric
	)
	return func(length ...int) string {
		value := 5
		if len(length) > 0 {
			value = length[0]
		}
		generator, _ := nanoid.CustomASCII(alphaNumeric+"_-", value)
		return generator()
	}
}

func SetGenerator(fn Generator) {
	generateKey = fn
}
