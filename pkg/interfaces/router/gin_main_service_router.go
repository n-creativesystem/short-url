package router

import (
	"html/template"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/n-creativesystem/short-url/pkg/interfaces/handler"
	"github.com/n-creativesystem/short-url/pkg/service/short"
)

func NewMainService(input *RouterInput) *gin.Engine {
	route := newGinRouter()
	shortService := short.NewService(input.ShortRepository, input.AppConfig, input.Beginner)
	shortHandler := handler.NewShortHandler(shortService)
	shortHandler.ServiceRouter(route)

	embedDir := newEmbedDir(Static, "static", "404.html")
	route.StaticFS("/static", embedDir)
	route.GET("/notfound", func(c *gin.Context) {
		f, _ := embedDir.Open("404.html")
		buf, _ := io.ReadAll(f)
		tpl, _ := template.New("404").Parse(string(buf))
		c.Render(http.StatusNotFound, render.HTML{
			Template: tpl,
		})
	})
	route.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/notfound")
	})
	return route
}
