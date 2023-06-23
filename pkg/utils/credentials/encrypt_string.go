package credentials

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

type EncryptStringer interface {
	MaskedStringer
	json.Marshaler
	json.Unmarshaler
}

var (
	_ EncryptStringer = (*EncryptString)(nil)
)

type EncryptString struct {
	*MaskedString
}

func NewEncryptString(value string) EncryptString {
	return EncryptString{
		MaskedString: NewMaskedString(value),
	}
}

func NewEncryptStringWithMustDecrypt(value string) EncryptString {
	v, err := crypto.Decrypt(value)
	if err != nil {
		logging.Default().Warn(err)
		v = ""
	}
	return EncryptString{
		MaskedString: NewMaskedString(v),
	}
}

func (m EncryptString) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.sensitiveData)
}

func (m *EncryptString) UnmarshalJSON(value []byte) error {
	var v string
	if err := json.Unmarshal(value, &v); err != nil {
		return err
	}
	*m = EncryptString{MaskedString: NewMaskedString(v)}
	return nil
}

func (m EncryptString) Value() (driver.Value, error) {
	return m.encrypt()
}

func (m *EncryptString) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	var (
		value string
		err   error
	)
	switch v := src.(type) {
	case string:
		value, err = crypto.Decrypt(v)
	case []byte:
		value, err = crypto.Decrypt(string(v))
	default:
		return fmt.Errorf("Unsupported type: %T", v)
	}
	if err != nil {
		return err
	}
	if m.MaskedString == nil {
		m.MaskedString = NewMaskedString(value)
	} else {
		m.MaskedString.sensitiveData = value
	}
	return nil
}

func (m EncryptString) MustEncrypt() string {
	v, err := m.encrypt()
	if err != nil {
		logging.Default().Warn(err)
		return ""
	}
	return v
}

func (m EncryptString) encrypt() (string, error) {
	return crypto.Encrypt(m.UnmaskedString())
}
