package credentials

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskedString(t *testing.T) {
	origMessage := "testMessage"
	maskedMessage := NewMaskedString(origMessage)
	assert.Equal(t, maskedValue, fmt.Sprint(maskedMessage))
	assert.Equal(t, maskedValue, fmt.Sprintf("%v", maskedMessage))
	assert.Equal(t, maskedValue, fmt.Sprintf("%#v", maskedMessage))
	assert.Equal(t, origMessage, maskedMessage.UnmaskedString())
	assert.True(t, maskedMessage.Equal(origMessage))
	assert.False(t, maskedMessage.Equal("testMessages"))
}
