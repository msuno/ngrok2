package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"ngrok/log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
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

var cache = make(map[string]interface{})

func Start(db *sqlx.DB, port string) {
	go readInfo(db)
	e := echo.New()
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rc := &RContext{
				c,
				db,
			}
			s := rc.Request().RequestURI
			if "/api/admin/login" != s && c.Request().Header.Get("Access-Token") != cache["access_token"] {
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

func readInfo(db *sqlx.DB) {
	t_net_i := float64(0)
	t_net_o := float64(0)
	t_disk_r := float64(0)
	t_disk_w := float64(0)
	for {
		percent, _ := cpu.Percent(time.Second, false)
		memInfo, _ := mem.VirtualMemory()
		d, _ := disk.IOCounters()
		nv, _ := net.IOCounters(false)
		ld, _ := load.Avg()
		pd, _ := process.Pids()

		net_i := float64(0)
		net_o := float64(0)
		for _, n := range nv {
			if n.BytesRecv == 0 && n.BytesSent == 0 {
				continue
			}
			net_i += float64(n.BytesRecv)
			net_o += float64(n.BytesSent)
		}
		tmp_i := net_i
		tmp_o := net_o
		if t_net_i != 0 || t_net_o != 0 {
			net_i -= t_net_i
			net_o -= t_net_o
		} else {
			net_i = float64(0)
			net_o = float64(0)
		}
		t_net_i = tmp_i
		t_net_o = tmp_o

		disk_r := float64(0)
		disk_w := float64(0)
		for _, v := range d {
			if v.ReadBytes == 0 && v.WriteBytes == 0 {
				continue
			}
			disk_r += float64(v.ReadBytes)
			disk_w += float64(v.WriteBytes)
		}

		tmp_r := disk_r
		tmp_w := disk_w
		if t_disk_r != 0 || t_disk_w != 0 {
			disk_r -= t_disk_r
			disk_w -= t_disk_w
		} else {
			disk_r = float64(0)
			disk_w = float64(0)
		}
		t_disk_r = tmp_r
		t_disk_w = tmp_w

		sys := &SystemInfo{
			Cpu:        percent[0],
			Mem:        memInfo.UsedPercent,
			DiskR:      disk_r,
			DiskW:      disk_w,
			NetI:       net_i,
			NetO:       net_o,
			Load:       ld.Load1,
			Pid:        float64(len(pd)),
			CreateTime: time.Now(),
		}

		_, err := db.NamedExec("insert into `system_info`(`cpu`,`mem`,`disk_r`, `disk_w`,`net_i`,`net_o`,`load`,`pid`,`create_time`) values(:cpu,:mem,:disk_r,:disk_w,:net_i,:net_o,:load,:pid,:create_time)", sys)
		if err != nil {
			log.Info("%v", err.Error())
		}
	}
}
