package entity

import (
	"github.com/gin-gonic/gin"
)

// Action 为实体的公开行为类型
type Action interface {
	Name() string
	New(ctx *gin.Context)
	Del(ctx *gin.Context)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
}
