package infrastruter

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
)

var (
	err    error
	engine *xorm.Engine
)

func initEngine() *xorm.Engine {
	engine, err = xorm.NewEngine("sqlite3", "./data.db")
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

type StoreEngine struct {
	M interface{}
	X *xorm.Engine
}

func (s StoreEngine) Save(m interface{}) error {
	_, err := s.X.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func (s StoreEngine) FetchAll(pageSize, page int, orderBy string, sort string, opts map[string]interface{}) (total int64, data interface{}, err error) {
	objTypeSlice := reflect.New(reflect.SliceOf(reflect.TypeOf(s.M)))
	data = objTypeSlice.Interface()
	offset := pageSize * (page - 1)
	objType := reflect.New(reflect.TypeOf(s.M).Elem())
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
		total, err = s.X.Where(query).Count(obj)
		if orderBy != "" {
			if sort == "asc" {
				err = s.X.Where(query).Asc(orderBy).Limit(pageSize, offset).OrderBy(orderBy).Find(data)
			} else {
				err = s.X.Where(query).Desc(orderBy).Limit(pageSize, offset).OrderBy(orderBy).Find(data)
			}
		} else {
			err = s.X.Where(query).Limit(pageSize, offset).Find(data)
		}
		if err != nil {
			return 0, nil, err
		}
	} else { // query is null
		total, err = s.X.Where("1 = ?", 1).Count(obj)
		if orderBy != "" {
			if sort == "asc" {
				err = s.X.Limit(pageSize, offset).Asc(orderBy).OrderBy(orderBy).Find(data)
			} else {
				err = s.X.Limit(pageSize, offset).Desc(orderBy).OrderBy(orderBy).Find(data)
			}
		} else {
			err = s.X.Limit(pageSize, offset).Find(data)
		}
		if err != nil {
			return 0, nil, err
		}
	}
	return total, data, nil
}
func (s StoreEngine) FetchOne(id string) (interface{}, error) {
	objType := reflect.New(reflect.TypeOf(s.M).Elem())
	data := objType.Interface()
	has, err := s.X.ID(id).Get(data)
	if err != nil {
		return nil, err
	}
	if has {
		return data, nil
	}
	return nil, errors.New("NOT EXIST")
}
func (s StoreEngine) Delete(id string) error {
	objType := reflect.New(reflect.TypeOf(s.M).Elem())
	data := objType.Interface()
	_, err := s.X.ID(id).Delete(data)
	return err
}

func (s StoreEngine) Update(id string, m interface{}) error {
	_, err := s.X.Id(id).Update(m)
	return err
}
