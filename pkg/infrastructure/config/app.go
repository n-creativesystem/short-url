package config

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v3"
)

type Decoder interface {
	Decode(v interface{}) (err error)
}

type NewDecoder func(r io.Reader) Decoder

func fileOrReader(filename string, r io.Reader, callback func(r io.ReadCloser) error) error {
	var reader io.ReadCloser
	if filename == "" {
		reader = io.NopCloser(r)
	} else {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		reader = f
	}
	defer reader.Close()
	return callback(reader)
}

func readCloserToStruct(r io.Reader, d NewDecoder, v interface{}) error {
	return d(r).Decode(v)
}

func setEnv(r io.Reader, overload bool) error {
	envMap, err := godotenv.Parse(r)
	if err != nil {
		return err
	}
	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		if !currentEnv[key] || overload {
			_ = os.Setenv(key, value)
		}
	}
	return nil
}

type applicationImpl struct {
	app appConfig
}

func NewApplication() config.ApplicationRepository {
	return &applicationImpl{
		app: appConfig{
			RetryGenerateCount: 10,
		},
	}
}

func (a *applicationImpl) Get(ctx context.Context, opts ...config.OptionFunc) (*config.Application, error) {
	opt := &config.Options{
		LoadTypes:         []config.LoadType{},
		EnvConfigLookuper: envconfig.OsLookuper(),
	}
	config.OptionApply(opt, opts...)
	app := a.app
	for _, typ := range opt.LoadTypes {
		switch typ {
		case config.Env:
			if err := fileOrReader(opt.EnvFile, opt.Reader, func(r io.ReadCloser) error {
				return setEnv(r, opt.EnvOverload)
			}); err != nil {
				return nil, err
			}
			if err := envconfig.ProcessWith(ctx, &app, opt.EnvConfigLookuper); err != nil {
				return nil, err
			}
		case config.Yaml:
			if err := fileOrReader(opt.YamlFile, opt.Reader, func(r io.ReadCloser) error {
				return readCloserToStruct(r, func(r io.Reader) Decoder { return yaml.NewDecoder(r) }, &app)
			}); err != nil {
				return nil, err
			}
		case config.Json:
			if err := fileOrReader(opt.JsonFile, opt.Reader, func(r io.ReadCloser) error {
				return readCloserToStruct(r, func(r io.Reader) Decoder { return json.NewDecoder(r) }, &app)
			}); err != nil {
				return nil, err
			}
		}
	}
	return app.ToModel(), nil
}

type appConfig struct {
	RetryGenerateCount int `env:"RETRY_GENERATE_COUNT,overwrite" yaml:"RETRY_GENERATE_COUNT,omitempty" json:"RETRY_GENERATE_COUNT,omitempty"`
}

func (a *appConfig) ToModel() *config.Application {
	return &config.Application{
		RetryGenerateCount: a.RetryGenerateCount,
	}
}
