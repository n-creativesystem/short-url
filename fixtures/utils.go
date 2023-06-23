package fixtures

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetDirectory() (string, error) {
	_, filePath, _, _ := runtime.Caller(0)
	fixtureDir := filepath.Dir(filePath)
	_, err := os.Stat(fixtureDir)
	if err != nil {
		return "", err
	}
	return fixtureDir, err
}
