package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xushiwang/AutoGenAPI/domain/entity"
)

type router struct {
	*gin.Engine
}

func NewRouter() router {
	r := router{gin.Default()}
	return r
}
func (r *router) Register(e entity.Action) {
	route := r.Group("/" + e.Name())
	route.POST("/", e.New)
	route.PUT("/:id", e.Update)
	route.DELETE("/:id", e.Del)
	route.GET("/:id", e.Get)
	route.GET("/", e.Get)
}
