package main

import (
	"DDD/application"
	"DDD/domain/entity"
	"DDD/router"
)

func main() {
	r := router.NewRouter()
	app := entity.NewEntity("applications", new(application.Applications))
	r.Register(app)
	r.Run("localhost:8080")
}
