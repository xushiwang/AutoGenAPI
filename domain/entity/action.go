package entity

import (
	"errors"
	"fmt"
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

type storeEngine interface {
	Save(m interface{}) error
	Delete(id string) error
	Update(id string, m interface{}) error
	FetchOne(id string) (interface{}, error)
	FetchAll(limit, offset int, sort string, opt ...map[string]string) (int64, interface{}, error)
}

type store struct {
	model interface{}
	storeEngine
	x *xorm.Engine
}

func (s store) Save(m interface{}) error {
	_, err := s.x.Insert(m)
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

func (s store) FetchAll(pageSize, page int, orderBy string, sort string, opts map[string]interface{}) (total int64, data interface{}, err error) {
	objTypeSlice := reflect.New(reflect.SliceOf(reflect.TypeOf(s.model)))
	data = objTypeSlice.Interface()
	offset := pageSize * (page - 1)
	objType := reflect.New(reflect.TypeOf(s.model).Elem())
	obj := objType.Interface()
	if len(opts) > 0 { // query not null
		// todo check query if right
		query := ""
		i := 0
		for k, v := range opts {
			if i == 0 {
				query += fmt.Sprintf("%s='%s'", k, v)
			} else {
				query += fmt.Sprintf(" and %s='%s'", k, v)
			}
			i++
		}
		total, err = s.x.Where(query).Count(obj)
		if orderBy != "" {
			if sort == "asc" {
				err = s.x.Where(query).Asc(orderBy).Limit(pageSize, offset).OrderBy(orderBy).Find(data)
			} else {
				err = s.x.Where(query).Desc(orderBy).Limit(pageSize, offset).OrderBy(orderBy).Find(data)
			}
		} else {
			err = s.x.Where(query).Limit(pageSize, offset).Find(data)
		}
		if err != nil {
			return 0, nil, err
		}
	} else { // query is null
		total, err = s.x.Where("1 = ?", 1).Count(obj)
		if orderBy != "" {
			if sort == "asc" {
				err = s.x.Limit(pageSize, offset).Asc(orderBy).OrderBy(orderBy).Find(data)
			} else {
				err = s.x.Limit(pageSize, offset).Desc(orderBy).OrderBy(orderBy).Find(data)
			}
		} else {
			err = s.x.Limit(pageSize, offset).Find(data)
		}
		if err != nil {
			return 0, nil, err
		}
	}
	return total, data, nil
}
func (s store) FetchOne(id string) (interface{}, error) {
	objType := reflect.New(reflect.TypeOf(s.model).Elem())
	data := objType.Interface()
	has, err := s.x.ID(id).Get(data)
	if err != nil {
		return nil, err
	}
	if has {
		return data, nil
	}
	return nil, errors.New("NOT EXIST")
}
func (s store) Delete(id string) error {
	objType := reflect.New(reflect.TypeOf(s.model).Elem())
	data := objType.Interface()
	_, err := s.x.ID(id).Delete(data)
	return err
}

func (s store) Update(id string, m interface{}) error {
	_, err := s.x.Id(id).Update(m)
	return err
}
