package tests

import (
	"os"
	"strings"
)

func EnvSetup() func() {
	values := os.Environ()
	mp := make(map[string]string, len(values))
	for _, value := range values {
		key, value, _ := strings.Cut(value, "=")
		mp[key] = value
	}
	return func() {
		os.Clearenv()
		for key, value := range mp {
			os.Setenv(key, value)
		}
	}
}
