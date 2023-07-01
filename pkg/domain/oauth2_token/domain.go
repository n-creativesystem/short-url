package oauth2token

import (
	"encoding/json"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
)

type TokenInfo interface {
	oauth2.TokenInfo
	Encode() string
}

type Token struct {
	oauth2.TokenInfo
}

func NewToken(token oauth2.TokenInfo) *Token {
	return &Token{TokenInfo: token}
}

func (t *Token) Encode() string {
	buf, _ := json.Marshal(t.TokenInfo)
	return crypto.MustEncrypt(string(buf))
}

func Decode(buf string) (TokenInfo, error) {
	return toToken(crypto.MustDecrypt(buf))
}

func toToken(data string) (TokenInfo, error) {
	var token models.Token
	if err := json.Unmarshal([]byte(data), &token); err != nil {
		return nil, repository.ErrRecordNotFound
	}
	return NewToken(&token), nil
}

type InnerToken struct {
	oauth2.TokenInfo
	UserID credentials.MaskedStringer
}

func (t *InnerToken) GetUserID() string {
	return t.UserID.UnmaskedString()
}

func (t *InnerToken) SetUserID(userID string) {
	t.UserID = credentials.NewMaskedString(userID)
}

func NewInnerToken(token oauth2.TokenInfo) *InnerToken {
	userId := credentials.NewEncryptStringWithMustDecrypt(token.GetUserID())
	return &InnerToken{
		TokenInfo: token,
		UserID:    &userId,
	}
}
