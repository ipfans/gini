package gini

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	defaultMapValue = map[string]string{}
	errMap          = map[error]errDesc{}
	errRWLock       sync.RWMutex
)

type errDesc struct {
	Status int
	Msg    string
}

// RenderBody json render layout.
type JSONRenderBody struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// RegisterError register error to RenderWrap
func RegisterError(err error, status int, desc string) {
	errRWLock.Lock()
	errMap[err] = errDesc{
		status,
		desc,
	}
	errRWLock.Unlock()
}

// JSONRenderWrap to write valid response.
func JSONRenderWrap(next func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := next(c)
		if err != nil {
			resp := JSONRenderBody{Data: defaultMapValue}

			errRWLock.RLock()
			v, ok := errMap[err]
			errRWLock.RUnlock()
			if !ok {
				resp.Status = http.StatusInternalServerError
				resp.Msg = err.Error()
			} else {
				resp.Status = v.Status
				resp.Msg = v.Msg
			}
			c.JSON(http.StatusOK, resp)
			return
		}
		data, _ := c.Get("data")
		resp := JSONRenderBody{
			Status: http.StatusOK,
			Data:   data,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}
