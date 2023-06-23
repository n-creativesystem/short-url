package config

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrCSRFSetting = errors.New("CSRF setting has not been configured.")
)

func checkCustomHeaderName(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return fmt.Errorf("No support type: %T", value)
	}
	if !strings.HasPrefix(strings.ToLower(v), "x-") {
		return errors.New("Start with `X-` or `x-`")
	}
	return nil
}
