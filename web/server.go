package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

type RContext struct {
	echo.Context
	redis redis.Client
}

func (r *RContext) BodyForm() map[string]interface{} {
	by, _ := ioutil.ReadAll(r.Request().Body)
	maps := make(map[string]interface{})
	_ = json.Unmarshal(by, &maps)
	return maps
}

func (r *RContext) Ok(str interface{}) error {
	maps := make(map[string]interface{})
	maps["code"] = 200
	maps["data"] = str
	return r.JSON(http.StatusOK, maps)
}

func (r *RContext) Fail(code int, str interface{}) error {
	maps := make(map[string]interface{})
	maps["code"] = code
	maps["msg"] = str
	return r.JSON(http.StatusOK, maps)
}

func Start(client *redis.Client) {
	e := echo.New()
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(&RContext{
				c,
				*client,
			})
		}
	})
	e.POST("/admin/login", login)
	e.GET("/admin/logout", logout)
	e.GET("/admin/router", menuList)
	e.GET("/admin/user", getUser)
	e.GET("/admin/system/info", sysInfo)
	e.GET("/admin/system/statistics", statistics)
	e.GET("/admin/user/list", userList)
	e.POST("/admin/user/setting", userAdd)
	e.DELETE("/admin/user/del", userDel)
	e.Logger.Fatal(e.Start(":1323"))
}
