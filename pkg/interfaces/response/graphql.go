package response

import (
	oauth2client "github.com/n-creativesystem/short-url/pkg/domain/oauth2_client"
	"github.com/n-creativesystem/short-url/pkg/interfaces/handler/graphql/models"
)

func OAuth2ApplicationResponseModel(value oauth2client.Client) models.OAuthApplication {
	return models.OAuthApplication{
		ID:     value.GetID(),
		Name:   value.GetAppName(),
		Secret: value.GetSecret(),
		Domain: value.GetDomain(),
	}
}
