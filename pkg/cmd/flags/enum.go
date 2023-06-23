package flags

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

type EnumFlags struct {
	target  string
	options []string
}

var (
	_ pflag.Value = (*EnumFlags)(nil)
)

func New(options ...string) *EnumFlags {
	target := options[0]
	return &EnumFlags{
		target:  target,
		options: options,
	}
}

func (e *EnumFlags) String() string {
	return e.target
}

func (e *EnumFlags) Set(value string) error {
	for _, v := range e.options {
		if v == value {
			e.target = value
			return nil
		}
	}
	return fmt.Errorf("expected one of %q", e.options)
}

func (e *EnumFlags) Type() string {
	return "enum"
}

func (e *EnumFlags) ChangeTarget(idx int) {
	e.target = e.options[idx]
}

func (e *EnumFlags) ChangeTargetByName(name string) {
	for idx, opt := range e.options {
		if strings.EqualFold(opt, name) {
			e.ChangeTarget(idx)
			return
		}
	}
}

func (e *EnumFlags) Clone() *EnumFlags {
	return New(e.options...)
}
