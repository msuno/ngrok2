package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"ngrok/server/assets"
	"ngrok/util"
	"os"
	"runtime"
	"strings"
	"time"
)

func start() {
	http.HandleFunc("/", static)
	http.HandleFunc("/admin/", static)
	http.HandleFunc("/admin/login", login)
	http.HandleFunc("/admin/user/index", index)
	http.HandleFunc("/admin/user/list", list)
	http.HandleFunc("/admin/user/setting", setting)
	http.HandleFunc("/admin/user/del", del)
	http.HandleFunc("/admin/msg/add", msgAdd)
	http.HandleFunc("/admin/online", online)
	http.HandleFunc("/admin/msg/list", msgList)
	_ = http.ListenAndServe("0.0.0.0:8000", nil)
}

func static(w http.ResponseWriter, r *http.Request)  {
	uri := r.RequestURI
	if uri == "/" {
		uri = "/admin/index.html"
	}
	var contentType string
	if strings.HasSuffix(uri, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(uri, ".css") {
		contentType = "text/css"
	} else if strings.Contains(uri, "/fonts") {
		contentType = "application/font-woff"
	} else {
		contentType = "text/html;charset=utf-8"
	}
	w.Header().Add("Content-Type", contentType)
	f := strings.ReplaceAll(uri, "/admin", "assets/server/admin")
	by, err := assets.Asset(f)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, _ = w.Write(by)
}

func setJsonContentType(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
}

func filter(w http.ResponseWriter, r *http.Request) bool {
	if "OPTIONS" == r.Method {
		w.WriteHeader(http.StatusOK)
		return false
	}
	auth := r.Header.Get(redisAuthKey)
	if auth == "" {
		w.WriteHeader(http.StatusProxyAuthRequired)
		return false
	}
	username := client.Get(redisAuthKey + auth).Val()
	if username == "" {
		w.WriteHeader(http.StatusExpectationFailed)
		return false
	}
	client.Expire(redisAuthKey + auth, time.Hour)
	setJsonContentType(w)
	return true
}

func login(w http.ResponseWriter, r *http.Request)  {
	if "OPTIONS" == r.Method {
		w.WriteHeader(http.StatusOK)
		return
	}
	by, _ := ioutil.ReadAll(r.Body)
	maps := make(map[string]string)
	_ = json.Unmarshal(by, &maps)
	username, _ := maps["username"]
	password, _ := maps["password"]
	if username == "" || password == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if strings.EqualFold("msuno", username) && strings.EqualFold("123456", password) {
		auth := util.RandId(32)
		client.Set(redisAuthKey + auth, username, time.Hour)
		w.Header().Add("Access-Control-Allow-Headers", redisAuthKey)
		w.Header().Add(redisAuthKey , auth)
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
}


func list(w http.ResponseWriter, r *http.Request)  {
	if !filter(w,r) {
		return
	}
	keys := client.HGetAll(redisUserList).Val()
	by, _ := json.Marshal(keys)
	_, _ = w.Write(by)
}


func index(w http.ResponseWriter, r *http.Request) {
	if !filter(w,r) {
		return
	}
	var men runtime.MemStats
	runtime.ReadMemStats(&men)
	m := make(map[string]interface{})
	m["go root"] = runtime.GOROOT()
	m["version"] = runtime.Version()
	m["cpu"] = runtime.GOMAXPROCS(0)
	m["goarch"] = runtime.GOARCH
	m["goos"] = runtime.GOOS
	m["total alloc"] = fmt.Sprintf("%d Kb",men.Alloc/1024)
	m["frees"] = fmt.Sprintf("%d Kb",men.Frees/1024)
	m["sys"] = fmt.Sprintf("%d Kb",men.Sys/1024)
	m["gc sys"] = fmt.Sprintf("%d Kb",men.GCSys/1024)
	m["next gc"] = fmt.Sprintf("%d Kb",men.NextGC/1024)
	m["last gc"] = fmt.Sprintf("%d Kb",men.LastGC/1024)
	m["num gc"] = fmt.Sprintf("%d Kb",men.NumGC/1024)
	m["hostname"], _ = os.Hostname()
	buf, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(buf)
}

func online(w http.ResponseWriter, r *http.Request) {
	if !filter(w,r) {
		return
	}
	buf, err := metrics.Msg()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(buf)
}

func setting(w http.ResponseWriter, r *http.Request) {
	if !filter(w,r) {
		return
	}
	by, _ := ioutil.ReadAll(r.Body)
	maps := make(map[string]string)
	_ = json.Unmarshal(by, &maps)
	key, _ := maps["key"]
	value, _ := maps["value"]
	if key == "" || value == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	client.HSet(redisUserList, key, value)
	w.WriteHeader(http.StatusOK)
}

func del(w http.ResponseWriter, r *http.Request) {
	if !filter(w,r) {
		return
	}
	by, _ := ioutil.ReadAll(r.Body)
	maps := make(map[string]string)
	_ = json.Unmarshal(by, &maps)
	key, _ := maps["key"]
	if key == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	client.HDel(redisUserList, key).Val()
	w.WriteHeader(http.StatusOK)
}

func msgList(w http.ResponseWriter, r *http.Request)  {
	if !filter(w,r) {
		return
	}
	keys := client.HGetAll(redisMsgKey).Val()
	by, _ := json.Marshal(keys)
	_, _ = w.Write(by)
}

func msgAdd(w http.ResponseWriter, r *http.Request) {
	if !filter(w,r) {
		return
	}
	by, _ := ioutil.ReadAll(r.Body)
	maps := make(map[string]string)
	_ = json.Unmarshal(by, &maps)
	username, _ := maps["username"]
	msg, _ := maps["msg"]
	if username == "" || msg == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	client.HSet(redisMsgKey, username, msg)
	w.WriteHeader(http.StatusOK)
}
