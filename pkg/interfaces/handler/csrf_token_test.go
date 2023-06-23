package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func testCookie() *gin.Engine {
	gin.SetMode(gin.TestMode)
	handler := NewCSRFTokenHandler()
	router := gin.New()
	router.GET("/", handler.GetToken())
	router.GET("/ignore", handler.Middleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.POST("/protected", handler.Middleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	return router
}

func testCSRFHandler(t *testing.T, router *gin.Engine, w *httptest.ResponseRecorder) string {
	require := require.New(t)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	respBody := map[string]string{}
	err := json.NewDecoder(w.Body).Decode(&respBody)
	require.NoError(err)
	v, ok := respBody["csrf_token"]
	require.True(ok)
	require.NotEmpty(t, v)
	cookies := w.Result().Cookies()
	mpCookies := make(map[string]*http.Cookie, len(cookies))
	for _, cookie := range cookies {
		c := *cookie
		mpCookies[cookie.Name] = &c
	}
	vc, ok := mpCookies["csrf_token"]
	require.True(ok)
	require.NotEmpty(t, vc.Value)
	return v
}

func TestCSRFHandler(t *testing.T) {
	router := testCookie()
	w := httptest.NewRecorder()
	testCSRFHandler(t, router, w)
}

func TestCSRFProtectCheck(t *testing.T) {
	var (
		w          *httptest.ResponseRecorder
		req        *http.Request
		err        error
		respBodyFn = func() map[string]interface{} { return map[string]interface{}{} }
		respBody   map[string]interface{}
	)
	require := require.New(t)
	router := testCookie()

	// Ignore path
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/ignore", nil)
	router.ServeHTTP(w, req)
	respBody = respBodyFn()
	err = json.NewDecoder(w.Body).Decode(&respBody)
	require.NoError(err)
	v, ok := respBody["status"].(string)
	require.True(ok)
	require.Equal(v, "ok")

	// Protect path
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/protected", nil)
	router.ServeHTTP(w, req)
	require.Equal(w.Result().StatusCode, http.StatusForbidden)
	respBody = respBodyFn()
	err = json.NewDecoder(w.Body).Decode(&respBody)
	require.NoError(err)
	errs, ok := respBody["errors"].([]interface{})
	require.True(ok)
	v, ok = errs[0].(map[string]interface{})["message"].(string)
	require.True(ok)
	require.Equal(v, "Invalid csrf token")

	w = httptest.NewRecorder()
	token := testCSRFHandler(t, router, w)
	cookies := w.Result().Cookies()
	tests := []struct {
		name   string
		header string
		form   string
		query  string
	}{
		{
			name:   "Header X-CSRF-Token",
			header: "X-CSRF-Token",
		},
		{
			name:   "Header X-XSRF-TOKEN",
			header: "X-XSRF-TOKEN",
		},
		{
			name: "From",
			form: "_csrf",
		},
		{
			name:  "Query",
			query: "_csrf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w = httptest.NewRecorder()
			req = httptest.NewRequest(http.MethodPost, "/protected", nil)
			for _, cookie := range cookies {
				req.AddCookie(cookie)
			}
			if tt.header != "" {
				req.Header.Add(tt.header, token)
			} else if tt.form != "" {
				req.PostForm.Add(tt.form, token)
			} else if tt.query != "" {
				q := req.URL.Query()
				q.Add(tt.query, token)
				req.URL.RawQuery = q.Encode()
			}
			router.ServeHTTP(w, req)
			respBody = respBodyFn()
			err = json.NewDecoder(w.Body).Decode(&respBody)
			require.NoError(err)
			v, ok = respBody["status"].(string)
			require.True(ok)
			require.Equal(v, "ok")
		})
	}

}
