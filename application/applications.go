package application

import (
	"fmt"
	"time"
)

type Applications struct {
	Id      int64     `json:"id,pk"`
	Created time.Time `json:"created" xorm:"created"`
	Updated time.Time `json:"updated" xorm:"updated"`
	Deleted time.Time `json:"deleted" xorm:"deleted"`

	Name string
	Url  string
}

func (a *Applications) BeforeCreate() {
	a.Url = "https://github.com"
}

func (a *Applications) AfterCreate() {
	fmt.Println("create success")
}
