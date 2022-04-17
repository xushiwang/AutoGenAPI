package main

import (
	"DDD/application"
	"DDD/domain/entity"
	"DDD/infrastruter"
	"DDD/router"
)

func main() {
	db := infrastruter.GetEngine()
	err := db.Ping()
	if err != nil {
		panic(any(err))
	}
	r := router.NewRouter()
	a := entity.NewEntity("applications", new(application.Applications))
	b := entity.NewEntity("users", new(application.Users))
	r.Register(a)
	r.Register(b)
	r.Run("localhost:8080")
}
