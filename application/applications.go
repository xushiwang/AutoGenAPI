package application

import "time"

type Applications struct {
	Id      int64
	Created time.Time `xorm:"CREATED"`
	Updated time.Time `xorm:"UPDATED"`
	Name    string
	Url     string
}
