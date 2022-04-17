package infrastruter

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	err    error
	engine *xorm.Engine
)

func initEngine() *xorm.Engine {
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		panic(any(err))
	}
	return engine
}

func GetEngine() *xorm.Engine {
	if engine == nil {
		return initEngine()
	}
	return engine
}
