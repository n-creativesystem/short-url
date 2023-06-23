package config

import (
	"io"

	"github.com/sethvargo/go-envconfig"
)

type LoadType int

const (
	Env LoadType = iota
	Yaml
	Json
)

type OptionFunc interface {
	apply(*Options)
}

func OptionApply(o *Options, opts ...OptionFunc) {
	for _, opt := range opts {
		opt.apply(o)
	}
}

type optionFn func(*Options)

func (fn optionFn) apply(o *Options) {
	fn(o)
}

type Options struct {
	LoadTypes         []LoadType
	EnvConfigLookuper envconfig.Lookuper
	EnvFile           string
	EnvOverload       bool
	YamlFile          string
	JsonFile          string
	Reader            io.Reader
}

func WithLoadType(typ LoadType) OptionFunc {
	return optionFn(func(o *Options) {
		o.LoadTypes = append(o.LoadTypes, typ)
	})
}

func WithYamlFile(filename string) OptionFunc {
	return optionFn(func(o *Options) {
		o.YamlFile = filename
	})
}

func WithJsonFile(filename string) OptionFunc {
	return optionFn(func(o *Options) {
		o.JsonFile = filename
	})
}

func WithEnvConfigLookuper(lookuper envconfig.Lookuper) OptionFunc {
	return optionFn(func(o *Options) {
		o.EnvConfigLookuper = lookuper
	})
}

func WithEnvFile(filename string) OptionFunc {
	return optionFn(func(o *Options) {
		o.EnvFile = filename
	})
}

func WithEnvOverload() OptionFunc {
	return optionFn(func(o *Options) {
		o.EnvOverload = true
	})
}

func WithReader(r io.Reader) OptionFunc {
	return optionFn(func(o *Options) {
		o.Reader = r
	})
}
