package gini

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBind(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.POST("/echo", func(c *gin.Context) {
		var data map[string]interface{}
		err := Bind(c, &data)
		if err != nil {
			panic(err)
		}
		c.JSON(200, data)
	})
	w := httptest.NewRecorder()
	body := strings.NewReader("{\"name\":\"yeeuu\"}")
	req, _ := http.NewRequest("POST", "/echo", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"name\":\"yeeuu\"}", w.Body.String())
}

func TestMockBind(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.POST("/echo", func(c *gin.Context) {
		var data map[string]interface{}
		err := Bind(c, &data)
		if err != nil {
			panic(err)
		}
		c.JSON(200, data)
	})
	mb := MockJSONBinder{}
	mb.Body("{\"name\":\"sam\"}")
	SetBinder(&mb)
	w := httptest.NewRecorder()
	body := strings.NewReader("{\"name\":\"yeeuu\"}")
	req, _ := http.NewRequest("POST", "/echo", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"name\":\"sam\"}", w.Body.String())
}
