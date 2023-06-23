package db

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/n-creativesystem/short-url/pkg/domain/config"
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

func MustGetDirectory() string {
	d, err := GetDirectory()
	if err != nil {
		panic(err)
	}
	return d
}

func getSubDirectory(dirname string) (string, error) {
	dir, err := GetDirectory()
	if err != nil {
		return "", err
	}
	filename := filepath.Join(dir, dirname)
	if stat, err := os.Stat(filename); err != nil {
		return "", err
	} else {
		if !stat.IsDir() {
			return "", fmt.Errorf("%s is not a directory.", dirname)
		}
	}
	return filename, nil
}

func GetConfigDirectory() (string, error) {
	dir, err := getSubDirectory("config")
	if err != nil {
		return "", err
	}
	return dir, nil
}

func MustGetConfigDirectory() string {
	dir, err := GetConfigDirectory()
	if err != nil {
		panic(err)
	}
	return dir
}

func GetMigrationsDirectory() (string, error) {
	subDir := ""
	switch config.GetDriver() {
	case config.MySQL:
		subDir = "mysql"
	case config.PostgreSQL:
		subDir = "postgres"
	case config.DynamoDB:
		subDir = "dynamodb"
	}
	dir, err := getSubDirectory(filepath.Join("migrations", subDir))
	if err != nil {
		return "", err
	}
	return dir, nil
}

func MustGetMigrationsDirectory() string {
	dir, err := GetMigrationsDirectory()
	if err != nil {
		panic(err)
	}
	return dir
}
