# AUTO GEN RESTFUL API
> 根据结构体自动生成符合restful风格的API
```
func main(){
    r := router.NewRouter()
    app := entity.NewEntity("applications", new(application.Applications))
    r.Register(app)
    r.Run("localhost:8080")
}

Output:
[GIN-debug] POST   /applications/            --> DDD/domain/entity.Action.New-fm (3 handlers)
[GIN-debug] PUT    /applications/:id         --> DDD/domain/entity.Action.Update-fm (3 handlers)
[GIN-debug] DELETE /applications/:id         --> DDD/domain/entity.Action.Del-fm (3 handlers)
[GIN-debug] GET    /applications/:id         --> DDD/domain/entity.Action.Get-fm (3 handlers)
[GIN-debug] GET    /applications/            --> DDD/domain/entity.Action.Get-fm (3 handlers)
[GIN-debug] Listening and serving HTTP on localhost:8080
```