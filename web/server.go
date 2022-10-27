package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type RContext struct {
	echo.Context
	db *sqlx.DB
}

type Res struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data,omitempty"`
	Msg       interface{} `json:"msg,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

type Page struct {
	Size  int         `json:"size"`
	Page  int         `json:"page"`
	Total int         `json:"total"`
	Data  interface{} `json:"data,omitempty"`
}

func (r *RContext) BodyForm() map[string]interface{} {
	by, _ := ioutil.ReadAll(r.Request().Body)
	maps := make(map[string]interface{})
	_ = json.Unmarshal(by, &maps)
	return maps
}

func (r *RContext) Ok(str interface{}) error {
	res := &Res{
		Code:      200,
		Data:      str,
		Timestamp: time.Now().UnixMilli(),
	}
	return r.JSON(http.StatusOK, res)
}

func (r *RContext) Fail(code int, str interface{}) error {
	res := &Res{
		Code:      code,
		Msg:       str,
		Timestamp: time.Now().UnixMilli(),
	}
	return r.JSON(http.StatusOK, res)
}

var cache = make(map[string]string)

var excludeUrl map[string]bool = map[string]bool{"/api/admin/login": true, "/api/admin/wechat": true}

func Start(db *sqlx.DB, port string) {
	go readInfo(db)
	e := echo.New()
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rc := &RContext{c, db}
			s := rc.Request().URL.Path
			_, ok := excludeUrl[s]
			if !ok && !strings.EqualFold(c.Request().Header.Get("Access-Token"), cache["access_token"]) {
				return rc.Fail(401, "Access Forbidden")
			}
			return h(rc)
		}
	})
	e.POST("/api/admin/login", login)
	e.GET("/api/admin/logout", logout)
	e.GET("/api/admin/router", menuList)
	e.GET("/api/admin/user", getUser)
	e.GET("/api/admin/system/info", sysInfo)
	e.GET("/api/admin/user/list", userList)
	e.POST("/api/admin/user/edit", userEdit)
	e.DELETE("/api/admin/user/del", userDel)
	e.GET("/api/admin/wechat", getWeChat)
	e.POST("/api/admin/wechat", postWeChat)
	e.Logger.Fatal(e.Start(port))
}
