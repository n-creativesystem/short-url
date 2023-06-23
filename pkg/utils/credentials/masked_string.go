package credentials

import (
	"crypto/subtle"
	"database/sql"
	"database/sql/driver"
	"encoding"
	"fmt"
)

const maskedValue = "******"

type MaskedStringer interface {
	fmt.Stringer
	fmt.GoStringer
	encoding.TextUnmarshaler
	sql.Scanner
	driver.Valuer
	UnmaskedString() string
	Equal(value string) bool
}

func NewMaskedString(sensitiveData string) *MaskedString {
	return &MaskedString{sensitiveData: sensitiveData}
}

var (
	_ MaskedStringer = (*MaskedString)(nil)
)

type MaskedString struct {
	sensitiveData string
}

func (m MaskedString) String() string {
	return maskedValue
}

func (m MaskedString) GoString() string {
	return maskedValue
}

func (m MaskedString) UnmaskedString() string {
	return string(m.sensitiveData)
}

func (m *MaskedString) Equal(value string) bool {
	return MaskedStringEqual(m, NewMaskedString(value))
}

func (m *MaskedString) UnmarshalText(text []byte) error {
	m.sensitiveData = string(text)
	return nil
}

func (m MaskedString) Value() (driver.Value, error) {
	return m.UnmaskedString(), nil
}

func (m *MaskedString) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch v := src.(type) {
	case string:
		m.sensitiveData = v
	case []byte:
		m.sensitiveData = string(v)
	default:
		return fmt.Errorf("Unsupported type: %T", v)
	}
	return nil
}

func MaskedStringEqual(a, b MaskedStringer) bool {
	return subtle.ConstantTimeCompare([]byte(a.UnmaskedString()), []byte(b.UnmaskedString())) == 1
}
