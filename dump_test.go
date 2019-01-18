package gini

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDumpReqAndResp(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.Use(DumpReqAndResp())
	r.POST("/echo", func(c *gin.Context) {
		var data map[string]interface{}
		c.Bind(&data)
		c.JSON(200, data)
	})
	r.POST("/empty", func(c *gin.Context) {
	})
	r.GET("/hello", func(c *gin.Context) {
		data := map[string]interface{}{}
		c.JSON(200, data)
	})
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()
	body := strings.NewReader("{\"name\":\"yeeuu\"}")
	req1, _ := http.NewRequest("POST", "/echo", body)
	req2, _ := http.NewRequest("POST", "/empty", body)
	req3, _ := http.NewRequest("GET", "/hello", nil)
	req1.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w1, req1)
	r.ServeHTTP(w2, req2)
	r.ServeHTTP(w3, req3)
	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, "{\"name\":\"yeeuu\"}", w1.Body.String())
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, 200, w3.Code)
}
