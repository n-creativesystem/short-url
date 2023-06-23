package router

import (
	"embed"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

var (
	//go:embed static/*
	Static embed.FS

	//go:embed static/app
	webUI embed.FS
)

const INDEX = "index.html"

type dir struct {
	http.FileSystem
	index string
}

func (d *dir) Exists(prefix string, filepath string) bool {
	path := strings.TrimPrefix(filepath, prefix)
	_, err := d.Open(path)
	return err == nil
}

func (d *dir) GetIndex(c *gin.Context) {
	f, _ := d.Open(d.index)
	buf, _ := io.ReadAll(f)
	tpl, _ := template.New("index").Parse(string(buf))
	c.Render(http.StatusOK, render.HTML{
		Template: tpl,
	})
}

func newEmbedDir(emfs embed.FS, path, index string) *dir {
	sub, _ := fs.Sub(emfs, path)
	if index == "" {
		index = INDEX
	}
	return &dir{
		FileSystem: http.FS(sub),
		index:      index,
	}
}
