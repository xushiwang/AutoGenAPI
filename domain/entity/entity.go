package entity

import (
	"DDD/infrastruter"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

type entity struct {
	model      interface{}
	action     Action
	engine     X
	entityName string
}

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func NewEntity(name string, model interface{}) *entity {
	if err := infrastruter.GetEngine().Sync2(model); err != nil {
		panic(any(err))
	}
	return &entity{
		engine:     X{x: infrastruter.GetEngine(), model: model},
		entityName: name,
		model:      model,
	}
}

func (b *entity) New(ctx *gin.Context) {
	m := reflect.New(reflect.TypeOf(b.model).Elem())
	data := m.Interface()
	ctx.BindJSON(data)
	err := b.engine.Save(data)
	if err != nil {
		FailWithMessage(err.Error(), ctx)
		return
	}
	OkWithMessage("INSERT OK", ctx)
}

func (b *entity) Del(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		FailWithMessage("DELETE METHOD MUST WITH AN ID", ctx)
		return
	}
	err := b.engine.Delete(id)
	if err != nil {
		FailWithMessage(err.Error(), ctx)
		return
	}
	OkWithMessage("DELETE OK", ctx)
}

func (b *entity) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		FailWithMessage("UPDATE METHOD MUST WITH AN ID", ctx)
		return
	}
	_, err := strconv.Atoi(id)
	if err != nil {
		FailWithMessage("ID ERROR", ctx)
		return
	}
	m := reflect.New(reflect.TypeOf(b.model).Elem())
	data := m.Interface()
	err = ctx.BindJSON(data)
	if err != nil {
		FailWithMessage(err.Error(), ctx)
		return
	}
	err = b.engine.Update(id, data)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	OkWithMessage("UPDATE OK", ctx)
}

func (b *entity) Name() string {
	return b.entityName
}

func (b *entity) Get(c *gin.Context) {
	id := c.Param("id")
	var data interface{}
	m := make(map[string]interface{}, 2)
	pageSize, page := 10, 1
	var sort string
	var total int64
	_, err := strconv.Atoi(id)
	if err != nil {
		for key, _ := range c.Request.URL.Query() {
			switch key {
			case "pageSize":
				ps := c.Request.URL.Query().Get(key)
				val, err := strconv.Atoi(ps)
				if err != nil {
					FailWithMessage(err.Error(), c)
					return
				}
				pageSize = val
			case "page":
				pa, err := strconv.Atoi(c.Request.URL.Query().Get(key))
				if err != nil {
					FailWithMessage(err.Error(), c)
					return
				}
				page = pa
			case "sort":
				sort = c.Request.URL.Query().Get(key)
			default:
				m[key] = c.Request.URL.Query().Get(key)
			}
		}
		total, data, err = b.engine.FetchAll(pageSize, page, sort, m)
		if err != nil {
			FailWithMessage(err.Error(), c)
			return
		}
		OkWithData(gin.H{"data": data, "total": total}, c)
		return
	} else {
		data, err = b.engine.FetchOne(id)
		if err != nil {
			FailWithMessage(err.Error(), c)
			return
		}
		OkWithData(data, c)
		return
	}
	//infrastruter.GetEngine().Asc("id").Find(&appdata)
	//id := c.Param("id")
	//fmt.Println(id)
	//if id == "" { // get all data
	//	//query := make(map[string]string)
	//	//limit, offset, sort := 10, 0, ""
	//	//var err error
	//	for key, val := range c.Request.URL.Query() {
	//		fmt.Println(key,val)
	//		//switch key {
	//		//case "limit":
	//		//	limit, err = strconv.Atoi(key)
	//		//	if err != nil {
	//		//		FailWithDetailed(http.StatusBadRequest, "LIMIT ERROR", c)
	//		//	}
	//		//case "offset":
	//		//
	//		//default:
	//		//	query[key] = c.Query(key)
	//		//}
	//
	//	}
	//	//b.engine.FetchAll()
	//	OkWithData("", c)
	//	//c.AbortWithError(200,errors.New("sdf"))
	//} else { // get one data
	//	data, err := b.engine.FetchOne(id)
	//	if err != nil {
	//		FailWithMessage(err.Error(), c)
	//		return
	//	}
	//	if data == nil {
	//		FailWithMessage(NOTFOUNDERROR, c)
	//	}
	//	c.JSON(http.StatusOK, response{Code: http.StatusOK, Msg: "FETCHONE"})
	//}
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, response{
		code,
		data,
		msg,
	})
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
