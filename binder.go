package gini

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Binder to parse request body to struct.
type Binder interface {
	Bind(*gin.Context, interface{}) error
}

var (
	binder Binder
)

func init() {
	binder = &defaultBinder{}
}

type defaultBinder struct {
}

func (db *defaultBinder) Bind(ctx *gin.Context, v interface{}) error {
	b := binding.Default(ctx.Request.Method, ctx.ContentType())
	return ctx.MustBindWith(v, b)
}

// SetBinder set Binder instance.
func SetBinder(b Binder) {
	binder = b
}

// Bind to parse request body to struct.
func Bind(ctx *gin.Context, v interface{}) error {
	return binder.Bind(ctx, v)
}

// MockJSONBinder json binder for mock.
type MockJSONBinder struct {
	body string
}

// Body set request body
func (mb *MockJSONBinder) Body(body string) {
	mb.body = body
}

// Bind for Binder interface.
func (mb *MockJSONBinder) Bind(c *gin.Context, v interface{}) error {
	return json.Unmarshal([]byte(mb.body), v)
}
