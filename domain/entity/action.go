package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"reflect"
)

type Action interface {
	Name() string
	New(ctx *gin.Context)
	Del(ctx *gin.Context)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type StoreEngine interface {
	Save(m interface{}) error
	Delete(id string) error
	Update(id string, m interface{}) error
	FetchOne(id string) (interface{}, error)
	FetchAll(limit, offset int, sort string, opt ...map[string]string) (int64, interface{}, error)
}

type X struct {
	model interface{}
	StoreEngine
	x *xorm.Engine
}

func (x X) Save(m interface{}) error {
	_, err := x.x.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

type Applications struct {
	Id   int
	Name string
	Url  string
}

func (x X) FetchAll(pageSize, page int, sort string, opt ...map[string]string) (int64, interface{}, error) {
	objTypeSlice := reflect.New(reflect.SliceOf(reflect.TypeOf(x.model)))
	data := objTypeSlice.Interface()
	offset := pageSize * (page - 1)

	objType := reflect.New(reflect.TypeOf(x.model).Elem())
	obj := objType.Interface()
	total, err := x.x.Where("Id > ?", -1).Count(obj)
	if err != nil {
		return 0, nil, err
	}
	if sort != "" {
		err = x.x.Limit(pageSize, offset).OrderBy(sort).Find(data)
	} else {
		err = x.x.Limit(pageSize, offset).Find(data)
	}
	if err != nil {
		return 0, nil, err
	}
	return total, data, nil
}
func (x X) FetchOne(id string) (interface{}, error) {
	objType := reflect.New(reflect.TypeOf(x.model).Elem())
	data := objType.Interface()
	_, err := x.x.ID(id).Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (x X) Delete(id string) error {
	objType := reflect.New(reflect.TypeOf(x.model).Elem())
	data := objType.Interface()
	_, err := x.x.ID(id).Delete(data)
	return err
}

func (x X) Update(id string, m interface{}) error {
	_, err := x.x.Id(id).Update(m)
	return err
}
