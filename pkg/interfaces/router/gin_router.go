package router

import (
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	domain_short "github.com/n-creativesystem/short-url/pkg/domain/short"
	"github.com/n-creativesystem/short-url/pkg/domain/social"
	"github.com/n-creativesystem/short-url/pkg/domain/tx"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/interfaces"
	"github.com/n-creativesystem/short-url/pkg/interfaces/middleware"
	oauth2client "github.com/n-creativesystem/short-url/pkg/service/oauth2_client"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

type RouterInput struct {
	OAuth2Token         oauth2.TokenStore
	OAuth2ClientService oauth2client.Service
	ShortRepository     domain_short.Repository
	Beginner            tx.ContextBeginner
	AppConfig           *config.Application
	SocialRepo          social.UserRepository

	SessionStore scs.Store
}

func NewRouterInput(
	shortRepository domain_short.Repository,
	oauth2Store oauth2.TokenStore,
	oauth2ClientService oauth2client.Service,
	beginner tx.ContextBeginner,
	socialRepo social.UserRepository,
) *RouterInput {
	return &RouterInput{
		ShortRepository:     shortRepository,
		OAuth2Token:         oauth2Store,
		OAuth2ClientService: oauth2ClientService,
		Beginner:            beginner,
		SocialRepo:          socialRepo,
	}
}

func newGinRouter() *gin.Engine {
	route := gin.New()
	route.Use(middleware.Logger("/healthz"), gin.Recovery())
	route.GET("/healthz", func(c *gin.Context) {
		ctx := c.Request.Context()
		db := interfaces.GetPing(interfaces.RDB)
		if err := db.PingContext(ctx); err != nil {
			msg := "Health check failed."
			slog.With(logging.WithErr(err)).ErrorContext(ctx, msg)
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "ng"})
		} else {
			c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		}
	})
	return route
}
