package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
)

// WebUIManifest
//
// @Summary Manifestの取得
// @Tags UI
// @Accept json
// @Produce json
// @Success 200 {object} response.WebUIManifest
// @Router /manifest [get]
// @ID UIManifest
func WebUIManifest(cfg *config.WebUI) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "")
		defer span.End()
		*c.Request = *c.Request.WithContext(ctx)
		var manifest response.WebUIManifest
		manifest.CsrfTokenBase = cfg.CSRF.TokenBase
		manifest.HeaderName = cfg.CSRF.HeaderName
		c.JSON(http.StatusOK, &manifest)
	}
}
