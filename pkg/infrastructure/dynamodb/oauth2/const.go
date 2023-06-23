package oauth2

import (
	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb/tables"
)

const (
	authorizationTable = tables.OAuth2Authorization
	accessTokenTable   = tables.OAuth2AccessToken
	refreshTokenTable  = tables.OAuth2RefreshToken
	clientTable        = tables.OAuth2Client
)

var tokenColumns = tables.OAuth2TokenColumns

var clientColumns = tables.OAuth2ClientColumns
