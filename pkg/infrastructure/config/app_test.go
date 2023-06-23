package config

import (
	"bytes"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/utils/tests"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name    string
	typ     config.LoadType
	buf     []byte
	opts    []config.OptionFunc
	want    config.Application
	wantErr error
}

func TestAppConfigDefault(t *testing.T) {
	tearDown := tests.EnvSetup()
	want := config.Application{
		RetryGenerateCount: 10,
	}
	repo := NewApplication()
	tests := []testTable{
		{
			name:    "Yaml",
			typ:     config.Yaml,
			buf:     []byte(`{}`),
			want:    want,
			wantErr: nil,
		},
		{
			name:    "Json",
			typ:     config.Json,
			buf:     []byte(`{}`),
			want:    want,
			wantErr: nil,
		},
		{
			name:    "Env",
			typ:     config.Env,
			buf:     []byte(``),
			want:    want,
			wantErr: nil,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			buf := bytes.NewReader(tt.buf)
			tt.opts = append(tt.opts, config.WithLoadType(tt.typ), config.WithReader(buf))
			app, err := repo.Get(context.Background(), tt.opts...)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
			assert.Equal(t, tt.want, *app)
		})
	}
	tearDown()
}

func TestAppConfigOverwrite(t *testing.T) {
	tearDown := tests.EnvSetup()
	want := config.Application{
		RetryGenerateCount: 2,
	}
	repo := NewApplication()
	tests := []testTable{
		{
			name:    "Yaml",
			typ:     config.Yaml,
			buf:     []byte(`RETRY_GENERATE_COUNT: 2`),
			want:    want,
			wantErr: nil,
		},
		{
			name:    "Json",
			typ:     config.Json,
			buf:     []byte(`{"RETRY_GENERATE_COUNT": 2}`),
			want:    want,
			wantErr: nil,
		},
		{
			name:    "Env",
			typ:     config.Env,
			buf:     []byte(`RETRY_GENERATE_COUNT=2`),
			want:    want,
			wantErr: nil,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewReader(tt.buf)
			tt.opts = append(tt.opts, config.WithLoadType(tt.typ), config.WithReader(buf))
			app, err := repo.Get(context.Background(), tt.opts...)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
			assert.Equal(t, tt.want, *app)
		})
	}
	tearDown()
}

func TestAppConfigForFile(t *testing.T) {
	tearDown := tests.EnvSetup()
	filename := "app_config"
	want := config.Application{
		RetryGenerateCount: 2,
	}
	repo := NewApplication()
	tests := []testTable{
		{
			name:    "Yaml",
			typ:     config.Yaml,
			buf:     []byte(`RETRY_GENERATE_COUNT: 2`),
			want:    want,
			wantErr: nil,
		},
		{
			name:    "Json",
			typ:     config.Json,
			buf:     []byte(`{"RETRY_GENERATE_COUNT": 2}`),
			want:    want,
			wantErr: nil,
		},
		{
			name:    "Env",
			typ:     config.Env,
			buf:     []byte(`RETRY_GENERATE_COUNT=2`),
			want:    want,
			wantErr: nil,
		},
		{
			name:    "not found error",
			typ:     config.Yaml,
			buf:     []byte(`{RETRY_GENERATE_COUNT}`),
			want:    want,
			wantErr: errors.New("open app_config: no such file or directory"),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := os.WriteFile(filename, tt.buf, 0644); err != nil {
				t.Fatal(err.Error())
			}
			var fileOpt config.OptionFunc
			switch tt.typ {
			case config.Yaml:
				if tt.wantErr != nil {
					_ = os.RemoveAll(filename)
				}
				fileOpt = config.WithYamlFile(filename)
			case config.Json:
				fileOpt = config.WithJsonFile(filename)
			case config.Env:
				fileOpt = config.WithEnvFile(filename)
			}
			tt.opts = append(tt.opts, config.WithLoadType(tt.typ), fileOpt)
			app, err := repo.Get(context.Background(), tt.opts...)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
			if err == nil {
				assert.Equal(t, tt.want, *app)
			} else {
				assert.Nil(t, app)
			}
		})
	}
	tearDown()
}

func TestAppConfigError(t *testing.T) {
	tearDown := tests.EnvSetup()
	want := config.Application{
		RetryGenerateCount: 10,
	}
	repo := NewApplication()
	tests := []testTable{
		{
			name:    "Yaml",
			typ:     config.Yaml,
			buf:     []byte(`RETRY_GENERATE_COUNT`),
			want:    want,
			wantErr: errors.New("yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `RETRY_G...` into config.appConfig"),
		},
		{
			name:    "Json",
			typ:     config.Json,
			buf:     []byte(`RETRY_GENERATE_COUNT`),
			want:    want,
			wantErr: errors.New("invalid character 'R' looking for beginning of value"),
		},
		{
			name:    "Env",
			typ:     config.Env,
			buf:     []byte(`あああ`),
			want:    want,
			wantErr: errors.New("unexpected character \"\\u0081\" in variable name near \"あああ\""),
		},
		{
			name:    "Env",
			typ:     config.Env,
			buf:     []byte(`RETRY_GENERATE_COUNT=10`),
			opts:    []config.OptionFunc{config.WithEnvConfigLookuper(nil)},
			want:    want,
			wantErr: errors.New("lookuper cannot be nil"),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.opts = append(tt.opts, config.WithLoadType(tt.typ), config.WithReader(bytes.NewReader(tt.buf)))
			app, err := repo.Get(context.Background(), tt.opts...)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
			assert.Nil(t, app)
		})
	}
	tearDown()
}
