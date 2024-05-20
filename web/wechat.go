package web

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"ngrok/log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type ReqMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        string   `xml:"MsgId,omitempty"`
	Event        string   `xml:"Event,omitempty"`
}

type RespMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}

func getWeChat(c echo.Context) error {
	cc := c.(*RContext)
	log.Info(cc.QueryString())
	return cc.String(http.StatusOK, cc.QueryParam("echostr"))
}

func postWeChat(c echo.Context) error {
	cc := c.(*RContext)
	var a ReqMsg
	cc.Bind(&a)

	var content string
	if a.MsgType == "event" {
		if a.Event == "subscribe" {
			content = "欢迎订阅"
			var u NgrokUser
			err := cc.db.Get(&u, "select * from ngrok_user where union_id = ?", a.FromUserName)
			if err != nil {
				u.UnionId = a.FromUserName
				u.Sk = strings.ToLower(a.FromUserName)
				u4 := uuid.New()
				uuid := []rune(strings.ReplaceAll(u4.String(), "-", ""))
				u.Domain = string(uuid[0:10])
				u.CreateTime = time.Now()
				u.UpdateTime = time.Now()
				_, err := cc.db.NamedExec("insert ignore into `ngrok_user`(`union_id`,`domain`,`sk`, `create_time`, `update_time`) values(:union_id,:domain,:sk,:create_time,:update_time)", u)
				if err != nil {
					panic(err)
				}
			}
		} else {
			content = "欢迎下次关注"
		}
	} else if a.MsgType == "text" {
		switch a.Content {
		case "我的密钥":
			var u NgrokUser
			err := cc.db.Get(&u, "select * from ngrok_user where union_id = ?", a.FromUserName)
			if err == nil {
				content = "你的密钥为：" + u.Sk + "，域名为：" + u.Domain
			} else {
				content = "请先关注"
			}
		case "内网穿透":
			content = "下载地址：https://github.com/msuno/ngrok2/releases"
		default:
			content = a.Content
		}

	}

	log.Info(fmt.Sprintf("%v", a))

	var rs RespMsg
	rs.Content = content
	rs.ToUserName = a.FromUserName
	rs.FromUserName = a.ToUserName
	rs.CreateTime = time.Now().Unix()
	rs.MsgType = a.MsgType

	by, _ := xml.MarshalIndent(rs, "", "")
	return cc.String(http.StatusOK, string(by))
}
