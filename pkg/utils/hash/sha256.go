package hash

import (
	"crypto/sha256"
	"database/sql/driver"
	"encoding/hex"
	"fmt"
)

func Sum(value ...[]byte) string {
	h := sha256.New()
	for _, v := range value {
		h.Write(v)
	}
	return hex.EncodeToString(h.Sum(nil))
}

type Hash struct {
	hashValue     string
	OriginalValue string
}

func (m Hash) Value() (driver.Value, error) {
	v := Sum([]byte(m.OriginalValue))
	if !m.isHash(v) {
		return v, nil
	}
	return m.hashValue, nil
}

func (m *Hash) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch v := src.(type) {
	case string:
		m.hashValue = v
	case []byte:
		m.hashValue = string(v)
	default:
		return fmt.Errorf("Unsupported type: %T", v)
	}
	m.OriginalValue = m.hashValue
	return nil
}

func (m *Hash) isHash(value string) bool {
	return m.hashValue == value
}

func NewHash(value string) Hash {
	return Hash{
		OriginalValue: value,
	}
}
