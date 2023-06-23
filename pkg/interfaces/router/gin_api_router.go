package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n-creativesystem/short-url/pkg/interfaces/handler"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
	_ "github.com/n-creativesystem/short-url/pkg/interfaces/router/swagger/api/docs"
	"github.com/n-creativesystem/short-url/pkg/service/short"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Short url
// @version 1.0
// @description 短縮URL生成 API
// @license.name nozomi.nishinohara
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenurl /api/v1/oauth2/token
// @in header
// @name OAuth2Application
// @description エンドポイントを保護します
func NewAPI(input *RouterInput) *gin.Engine {
	mainRoute := newGinRouter()
	api := mainRoute.Group("/api/v1")
	oauth2Handler := handler.NewOAuthHandler(input.OAuth2ClientService, input.OAuth2Token, input.AppConfig)
	oauth2Handler.Router(api)

	shortService := short.NewService(input.ShortRepository, input.AppConfig, input.Beginner)
	shortHandler := handler.NewShortHandler(shortService)
	shortHandler.APIRouter(api, oauth2Handler.ValidationBearerToken())
	swag := mainRoute.Group("/swagger")
	{
		// index.htmlでswagger uiを表示することができる
		swag.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	mainRoute.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, response.NewErrorsWithMessage("Not found."))
	})
	return mainRoute
}
