package main

import (
	"github.com/xushiwang/AutoGenAPI/application"
	"github.com/xushiwang/AutoGenAPI/domain/entity"
	"github.com/xushiwang/AutoGenAPI/router"
)

func main() {
	r := router.NewRouter()
	app := entity.NewEntity("applications", new(application.Applications))
	r.Register(app)
	r.Run("localhost:8080")
}
