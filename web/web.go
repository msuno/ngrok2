package web

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"ngrok/log"
	"ngrok/util"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func login(c echo.Context) error {
	cc := c.(*RContext)
	access_token := util.RandId(16)
	refresh_token := util.RandId(16)
	res := make(map[string]interface{})
	res["access_token"] = access_token
	res["refresh_token"] = refresh_token
	res["expired_in"] = util.RedisAccessExpired.Seconds()
	return cc.Ok(res)
}

func logout(c echo.Context) error {
	cc := c.(*RContext)
	return cc.Ok("")
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
	var maps [2]map[string]interface{}
	str := `[{"id":1,"parentId":0,"name":"ProxyUser","path":"/ProxyUser","component":"Layout","redirect":"/ProxyUser/Directive","meta":{"title":"代理用户","icon":"el-icon-phone"}},{"id":10,"parentId":1,"name":"ProxyUser","path":"/ProxyUser/ProxyUser","component":"ProxyUser","meta":{"title":"用户管理","icon":"el-icon-goods"}}]`
	json.Unmarshal([]byte(str), &maps)
	return cc.Ok(maps)
}

func getWeChat(c echo.Context) error {
	cc := c.(*RContext)
	log.Info(cc.QueryString())
	return cc.String(http.StatusOK, cc.QueryParam("echostr"))
}

type Msg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        string   `xml:"MsgId,omitempty"`
}

func postWeChat(c echo.Context) error {
	cc := c.(*RContext)
	var a Msg
	cc.Bind(&a)
	formUser := a.FromUserName
	a.FromUserName = a.ToUserName
	a.ToUserName = formUser

	var u NgrokUser
	u.UnionId = formUser
	u.Sk = formUser
	u.Domain = formUser
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()
	cc.db.NamedExec("insert into `ngrok_user`(`union_id`,`domain`,`sk`, `create_time`, `update_time`) values(:union_id,:domain,:sk,:create_time,:update_time)", u)

	by, _ := xml.MarshalIndent(a, " ", " ")
	return cc.String(http.StatusOK, string(by))
}
