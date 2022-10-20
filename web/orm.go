package web

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type NgrokUser struct {
	Id         int64     `json:"id" db:"id"`
	UnionId    string    `json:"union_id" db:"union_id"`
	Domain     string    `json:"domain" db:"domain"`
	Sk         string    `json:"sk" db:"sk"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type User struct {
	Id          int64     `json:"id" db:"id"`
	Avatar      string    `json:"avatar" db:"avatar"`
	Description string    `json:"description" db:"description"`
	Email       string    `json:"email" db:"email"`
	MfaKey      string    `json:"-" db:"mfa_key"`
	MfaType     string    `json:"-" db:"mfa_type"`
	Nickname    string    `json:"nickname" db:"nickname"`
	Password    string    `json:"-" db:"password"`
	Username    string    `json:"username" db:"username"`
	ExpireTime  time.Time `json:"expire_time" db:"expire_time"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	UpdateTime  time.Time `json:"update_time" db:"update_time"`
}

type SystemInfo struct {
	Id         int64     `json:"id" db:"id"`
	Cpu        float64   `json:"cpu" db:"cpu"`
	Mem        float64   `json:"mem" db:"mem"`
	DiskR      float64   `json:"disk_r" db:"disk_r"`
	DiskW      float64   `json:"disk_w" db:"disk_w"`
	NetI       float64   `json:"net_i" db:"net_i"`
	NetO       float64   `json:"net_o" db:"net_o"`
	Load       float64   `json:"load" db:"load"`
	Pid        float64   `json:"pid" db:"pid"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

type ProxyUser struct {
	Id         int64     `json:"id" db:"id"`
	Domain     string    `json:"domain" db:"domain"`
	UnionId    string    `json:"union_id" db:"union_id"`
	Sk         string    `json:"sk" db:"sk"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type MenuInfo struct {
	Id         int64     `json:"id" db:"id"`
	ParentId   string    `json:"parentId" db:"parent_id"`
	Name       string    `json:"name" db:"name"`
	Path       string    `json:"path" db:"path"`
	Component  string    `json:"component" db:"component"`
	Redirect   string    `json:"redirect" db:"redirect"`
	Meta       Meta      `json:"meta" db:"meta" type:"json"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
}

type Meta struct {
	Title string `json:"title" db:"title"`
	Icon  string `json:"icon" db:"icon"`
}

func (t *Meta) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), &t)
}

func (t *Meta) Value() (driver.Value, error) {
	return json.Marshal(t)
}
