package oauth2token

import (
	"testing"

	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/stretchr/testify/require"
)

func TestDomain(t *testing.T) {
	r := require.New(t)
	tokenInfo := models.NewToken()
	tokenInfo.SetClientID("Client id")
	token := NewToken(tokenInfo)
	encode := token.Encode()
	info, err := Decode(encode)
	r.NoError(err)
	r.Equal(tokenInfo.GetClientID(), info.GetClientID())
}
