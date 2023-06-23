package utils

import (
	"os"
	"strconv"
	"strings"
)

func AppEnv() string {
	return os.Getenv("APP_ENV")
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
	v, _ := strconv.ParseBool(os.Getenv("CI"))
	return v
}

func IsCIorTest() bool {
	return IsCI() || IsTest()
}

func IsDevOrCIorTest() bool {
	return IsDev() || IsCI() || IsTest()
}
