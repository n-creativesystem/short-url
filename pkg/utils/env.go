package utils

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type AppMode = int

const (
	Service AppMode = iota
	API
	UI
)

var (
	appMode     AppMode
	initAppMode sync.Once

	appEnv     string
	initAppEnv sync.Once

	isCI     bool
	initIsCI sync.Once
)

func setAppMode(mode AppMode) {
	initAppMode.Do(func() {
		appMode = mode
	})
}

func RunService() {
	setAppMode(Service)
}

func IsService() bool {
	return appMode == Service
}

func RunAPI() {
	setAppMode(API)
}

func IsAPI() bool {
	return appMode == API
}

func RunUI() {
	setAppMode(UI)
}

func IsUI() bool {
	return appMode == UI
}

func AppEnv() string {
	initAppEnv.Do(func() {
		appEnv = os.Getenv("APP_ENV")
	})
	return appEnv
}

func IsProduction() bool {
	return strings.EqualFold(AppEnv(), "production")
}

func IsStaging() bool {
	return strings.EqualFold(AppEnv(), "staging")
}

func IsDev() bool {
	return strings.EqualFold(AppEnv(), "dev")
}

func IsTest() bool {
	return strings.EqualFold(AppEnv(), "test")
}

func IsCI() bool {
	initIsCI.Do(func() {
		v, _ := strconv.ParseBool(os.Getenv("CI"))
		isCI = v
	})
	return isCI
}

func IsCIorTest() bool {
	return IsCI() || IsTest()
}

func IsDevOrCIorTest() bool {
	return IsDev() || IsCI() || IsTest()
}

func Getenv(key, default_ string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return default_
}

func GetBoolEnv(key string) bool {
	e := os.Getenv(key)
	b, err := strconv.ParseBool(e)
	return err == nil && b
}
