package utils

import (
	"os"
	"testing"

	"github.com/n-creativesystem/short-url/pkg/utils/tests"
	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	tearDown := tests.EnvSetup()
	defer tearDown()
	envTests := []struct {
		name          string
		value         string
		production    bool
		staging       bool
		dev           bool
		test          bool
		ci            bool
		ciOrTest      bool
		devOrTestOrCI bool
	}{
		{
			name:       "Prod",
			value:      "PRODUCTION",
			production: true,
		},
		{
			name:    "Staging",
			value:   "STAGING",
			staging: true,
		},
		{
			name:          "Dev",
			value:         "DEV",
			dev:           true,
			devOrTestOrCI: true,
		},
		{
			name:          "Test",
			value:         "TEST",
			test:          true,
			ciOrTest:      true,
			devOrTestOrCI: true,
		},
	}
	t.Parallel()
	for _, tt := range envTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("APP_ENV", tt.value)
			assert.Equal(t, tt.production, IsProduction())
			assert.Equal(t, tt.staging, IsStaging())
			assert.Equal(t, tt.dev, IsDev())
			assert.Equal(t, tt.test, IsTest())
			assert.Equal(t, tt.ci, IsCI())
			assert.Equal(t, tt.ciOrTest, IsCIorTest())
			assert.Equal(t, tt.devOrTestOrCI, IsDevOrCIorTest())
		})
	}
	for _, value := range []string{"1", "t", "T", "true", "TRUE", "True"} {
		os.Setenv("CI", value)
		assert.True(t, IsCI())
		assert.True(t, IsCIorTest())
		assert.True(t, IsDevOrCIorTest())
	}

}
