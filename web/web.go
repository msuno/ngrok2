package web

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"net/http"
	"ngrok/log"
	"ngrok/util"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func login(c echo.Context) error {
	cc := c.(*RContext)
	m := cc.BodyForm()
	i := m["username"]
	p := m["password"]
	pp, ok := p.(string)
	if !ok {
		return cc.Fail(401, "password is empty")
	}
	md5Password := fmt.Sprintf("%x", md5.Sum([]byte(pp)))
	var u User
	err := cc.db.Get(&u, "select * from users where email = ? and password = ?", i, md5Password)
	if err != nil {
		return cc.Fail(401, "email error")
	}
	access_token := util.RandId(16)
	cache["access_token"] = access_token
	return cc.Ok(cache)
}

func logout(c echo.Context) error {
	cc := c.(*RContext)
	delete(cache, "access_token")
	return cc.Ok(cache)
}

func userList(c echo.Context) error {
	cc := c.(*RContext)
	page, err := strconv.ParseInt(cc.QueryParam("page"), 10, 64)
	if err != nil {
		page = 1
	}
	size, err := strconv.ParseInt(cc.QueryParam("size"), 10, 64)
	if err != nil {
		size = 10
	}
	start, err := strconv.ParseInt(cc.QueryParam("start"), 10, 64)
	if err != nil {
		start = time.Now().AddDate(-1, 0, 0).UnixMilli()
	}
	end, err := strconv.ParseInt(cc.QueryParam("end"), 10, 64)
	if err != nil {
		end = time.Now().UnixMilli()
	}
	search := "%" + cc.QueryParam("search") + "%"
	offset := (page - 1) * size
	var p []ProxyUser
	err = cc.db.Select(&p, "select * from ngrok_user where create_time >= ? and create_time <= ? and domain like ? limit ?,?", time.UnixMilli(start), time.UnixMilli(end), search, offset, size)
	if err != nil {
		log.Info("%v", err.Error())
	}
	var t int
	err = cc.db.Get(&t, "select count(*) from ngrok_user where create_time >= ? and create_time <= ? and domain like ? ", time.UnixMilli(start), time.UnixMilli(end), search)
	if err != nil {
		log.Info("%v", err.Error())
	}
	resPage := &Page{
		Page:  int(page),
		Size:  int(size),
		Total: t,
		Data:  p,
	}
	return cc.Ok(resPage)
}

func userEdit(c echo.Context) error {
	cc := c.(*RContext)
	var user ProxyUser
	cc.Bind(&user)
	if user.Id > 0 {
		cc.db.NamedExec("update ngrok_user set domain=:domain, union_id=:union_id, sk=:sk where id = :id", &user)
	} else {
		cc.db.NamedExec("insert into ngrok_user(domain, union_id, sk) values(:domain, :union_id, :sk)", &user)
	}
	return cc.Ok("")
}

func userDel(c echo.Context) error {
	cc := c.(*RContext)
	id := cc.QueryParam("id")
	_, err := cc.db.Exec("delete from ngrok_user where id = ?", id)
	if err != nil {
		return cc.Fail(501, err.Error())
	}
	return cc.Ok("")
}

func sysInfo(c echo.Context) error {
	cc := c.(*RContext)
	id := cc.QueryParam("id")
	var list []SystemInfo
	if id != "" {
		cc.db.Select(&list, "select * from system_info where id > ? order by create_time desc limit 100", id)
		return cc.Ok(list)
	}

	cc.db.Select(&list, "select * from system_info order by create_time desc limit 60")
	return cc.Ok(list)
}

func getUser(c echo.Context) error {
	cc := c.(*RContext)
	token := "msuno@msuno.cn"
	var u User
	cc.db.Get(&u, "select * from users where email = ? limit 1", token)
	return cc.Ok(u)
}

func menuList(c echo.Context) error {
	cc := c.(*RContext)
	var list []MenuInfo
	err := cc.db.Select(&list, "select * from menu_info order by id desc")
	if err != nil {
		panic(err)
	}
	return cc.Ok(list)
}

func getWeChat(c echo.Context) error {
	cc := c.(*RContext)
	log.Info(cc.QueryString())
	return cc.String(http.StatusOK, cc.QueryParam("echostr"))
}

type ReqMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        string   `xml:"MsgId,omitempty"`
}

type RespMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}

func postWeChat(c echo.Context) error {
	cc := c.(*RContext)
	var a ReqMsg
	cc.Bind(&a)

	var u NgrokUser
	u.UnionId = a.FromUserName
	u.Sk = a.FromUserName
	u.Domain = a.FromUserName
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()
	_, err := cc.db.NamedExec("insert into `ngrok_user`(`union_id`,`domain`,`sk`, `create_time`, `update_time`) values(:union_id,:domain,:sk,:create_time,:update_time)", u)
	if err != nil {
		panic(err)
	}

	var rs RespMsg
	rs.Content = a.Content
	rs.ToUserName = a.FromUserName
	rs.FromUserName = a.ToUserName
	rs.CreateTime = time.Now().Unix()
	rs.MsgType = a.MsgType

	by, _ := xml.MarshalIndent(rs, "", "")
	return cc.String(http.StatusOK, string(by))
}
