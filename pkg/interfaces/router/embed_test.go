package router

import (
	"embed"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//go:embed tests
var testEmbed embed.FS

func TestEmbed(t *testing.T) {
	embedDir := newEmbedDir(testEmbed, "tests", "")
	assert.True(t, embedDir.Exists("", "test1.txt"))
	assert.True(t, embedDir.Exists("", "sub/test2.txt"))
	assert.False(t, embedDir.Exists("", "notfound.txt"))
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	embedDir.GetIndex(c)
	assert.Equal(t, "This is test.\n", w.Body.String())
}
