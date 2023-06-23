package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
)

type contentType int

const (
	jsonType contentType = iota
	formType
	otherType
)

func convertContentTypeToEnum(cType string) contentType {
	v := strings.ToLower(strings.TrimSpace(cType))
	if strings.HasPrefix(v, "application/json") {
		return jsonType
	}
	if strings.HasPrefix(v, "application/x-www-form-urlencoded") {
		return formType
	}
	return otherType
}

func convertContentTypeToEnumWithRequest(r *http.Request) contentType {
	v := strings.ToLower(strings.TrimSpace(r.Header.Get("Content-Type")))
	return convertContentTypeToEnum(v)
}

func getContext[T any](c *gin.Context, name string) (T, bool) {
	v, ok := c.Get(name)
	if !ok {
		var value T
		return value, false
	}
	value, ok := v.(T)
	return value, ok
}

func generateRandomBytes(length int) []byte {
	buf := make([]byte, length)
	_, _ = rand.Read(buf)
	return buf
}

func randomString(length int) string {
	buf := generateRandomBytes(32)
	return base64.StdEncoding.EncodeToString(buf)
}

func bindJSON[T any](c *gin.Context, request *T) bool {
	if err := c.ShouldBindJSON(request); err != nil {
		errRes := response.NewErrorsWithMessage("Invalid request")
		c.AbortWithStatusJSON(http.StatusBadRequest, errRes)
		return false
	}
	return true
}

func validation(c *gin.Context, req interface{ Valid() *response.Errors }) bool {
	if errRes := req.Valid(); errRes != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errRes)
		return false
	}
	return true
}
