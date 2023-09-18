package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestIgnoreError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	_ = c.Error(ErrAuthorize)
	require := require.New(t)
	require.True(IsIgnoreError(c.Errors))
}
