package gini

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRenderWrap(t *testing.T) {
	r := gin.New()
	err := errors.New("whoooo")
	gin.SetMode(gin.ReleaseMode)
	r.GET("/", JSONRenderWrap(func(c *gin.Context) error {
		c.Set("data", "hello")
		return nil
	}))
	r.GET("/err", JSONRenderWrap(func(c *gin.Context) error {
		return err
	}))
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/", nil)
	req2, _ := http.NewRequest("GET", "/err", nil)
	r.ServeHTTP(w1, req1)
	r.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, "{\"status\":200,\"msg\":\"\",\"data\":\"hello\"}", w1.Body.String())
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, "{\"status\":500,\"msg\":\"whoooo\",\"data\":{}}", w2.Body.String())
	RegisterError(err, 400, "hahaha")
	req3, _ := http.NewRequest("GET", "/err", nil)
	r.ServeHTTP(w3, req3)
	assert.Equal(t, 200, w3.Code)
	assert.Equal(t, "{\"status\":400,\"msg\":\"hahaha\",\"data\":{}}", w3.Body.String())
	r.ServeHTTP(w3, req3)
}
