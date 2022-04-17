package main

import (
	"DDD/domain/entity"
	"DDD/infrastruter"
	"DDD/router"
	"time"
)

type Applications struct {
	Id      int64
	Created time.Time `xorm:"CREATED"`
	Updated time.Time `xorm:"UPDATED"`
	Name    string
	Url     string
}

func main() {
	db := infrastruter.GetEngine()
	err := db.Ping()
	if err != nil {
		panic(any(err))
	}
	r := router.NewRouter()
	a := entity.NewEntity("applications", new(Applications))
	r.Register(a)
	r.Run("localhost:8080")
}
