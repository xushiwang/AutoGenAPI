package application

import "time"

type Users struct {
	Id       int64
	Created  time.Time `xorm:"CREATED"`
	Updated  time.Time `xorm:"UPDATED"`
	Username string
	Password string
}
