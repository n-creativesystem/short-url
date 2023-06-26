package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/interfaces/handler"
	"github.com/n-creativesystem/short-url/pkg/interfaces/handler/graphql"
	"github.com/n-creativesystem/short-url/pkg/interfaces/middleware"
	"github.com/n-creativesystem/short-url/pkg/interfaces/middleware/session"
	"github.com/n-creativesystem/short-url/pkg/service/short"
)

// @title Web UI
// @version 1.0
// @description 短縮URL管理WEB UI
// @BasePath /api
// @schemes http https
// @license.name nozomi.nishinohara
// @scope
// @description エンドポイントを保護します
func NewWebUI(input *RouterInput, cfg *config.WebUI) *gin.Engine {
	router := newGinRouter()
	route := router.Group(cfg.Prefix)
	route.Use(cors.New(cfg.Cors.ToCorsConfig()))
	route.Use(middleware.Session(session.WithSessionStore(cfg.Store)))
	var csrfMiddleware gin.HandlerFunc
	if cfg.CSRF.TokenBase {
		csrfHandler := handler.NewCSRFTokenHandler()
		route.GET("/csrf_token", csrfHandler.GetToken())
		csrfMiddleware = csrfHandler.Middleware()
	} else {
		csrfMiddleware = func(ctx *gin.Context) {
			ctx.Next()
		}
	}
	route.GET("/manifest", handler.WebUIManifest(cfg))

	shortService := short.NewService(input.ShortRepository, input.AppConfig, input.Beginner)
	socialHandler := handler.NewSocialHandler(cfg.Providers, cfg.LoginSuccessURL, cfg.LogoutSuccessURL)
	socialHandler.Router(route, middleware.Protected())
	// Graphql handler
	gqlResolver := graphql.NewResolver(input.OAuth2ClientService, shortService)
	route.POST("/graphql", middleware.Protected(), csrfMiddleware, graphql.GraphQLHandler(gqlResolver))
	route.GET("/playground", graphql.GraphQLPlayGroundHandler("/api/graphql"))

	if cfg.IsUI {
		// static file(UI)
		webUIEmbedDir := newEmbedDir(webUI, "static/app", INDEX)
		router.Use(static.Serve("", webUIEmbedDir))
		// router.GET("/*file", gin.WrapH(webUIEmbedServer))
		router.NoRoute(func(c *gin.Context) {
			// HTMLを返すために用意
			webUIEmbedDir.GetIndex(c)
		})
	}
	return router
}
