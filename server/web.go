package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"ngrok/util"
	"os"
	"runtime"
	"strings"
)

func start() {
	http.HandleFunc("/admin/login", login)
	http.HandleFunc("/admin/reflash", reflash)
	http.HandleFunc("/admin/logout", logout)
	http.HandleFunc("/admin/check", check)
	http.HandleFunc("/admin/user/list", list)
	http.HandleFunc("/admin/user/setting", setting)
	http.HandleFunc("/admin/user/del", del)
	http.HandleFunc("/admin/system/info", info)
	http.HandleFunc("/admin/system/statistics", statistics)
	_ = http.ListenAndServe("0.0.0.0:8000", nil)
}

func setJsonContentType(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
}

func filter(w http.ResponseWriter, r *http.Request) bool {
	if "OPTIONS" == r.Method {
		w.WriteHeader(http.StatusOK)
		return false
	}
	auth := r.Header.Get(redisAuthHeader)
	if "" == auth {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(retFail(10012, "admin authorization is empty"))
		return false
	}
	if client.Exists(buildAccessToken(auth)).Val() < 1 {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write(ret(10013, "error", "admin authorization is expried"))
		return false
	}
	return true
}

func reflash(w http.ResponseWriter, r *http.Request) {
	setJsonContentType(w)
	access_token := r.Header.Get(redisAuthHeader)
	by, _ := ioutil.ReadAll(r.Body)
	maps := make(map[string]string)
	_ = json.Unmarshal(by, &maps)
	refresh_token, _ := maps["refresh_token"]
	redis_access_token := client.Get(buildReflashToken(refresh_token)).Val()
	if redis_access_token != access_token {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write(retFail(10014, "refresh token not match admin authorization"))
		return
	}
	client.Del(buildAccessToken(access_token))
	client.Del(buildReflashToken(refresh_token))
	access_token = util.RandId(16)
	refresh_token = util.RandId(16)
	res := make(map[string]interface{})
	res["access_token"] = access_token
	res["refresh_token"] = refresh_token
	res["expired_in"] = redisAccessExpired.Seconds()
	client.Set(buildReflashToken(refresh_token), access_token, redisReflashExpired)
	client.Set(buildAccessToken(access_token), refresh_token, redisAccessExpired)
	_, _ = w.Write(retOk(res))

}

func logout(w http.ResponseWriter, r *http.Request) {
	setJsonContentType(w)
	if filter(w, r) {
		access_token := r.Header.Get(redisAuthHeader)
		refresh_token := client.Get(buildAccessToken(access_token))
		client.Del(buildAccessToken(access_token))
		client.Del(buildReflashToken(refresh_token.Val()))
		_, _ = w.Write(retOk("SUCCESS"))
	}

}

func check(w http.ResponseWriter, r *http.Request) {
	setJsonContentType(w)
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
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(retFail(10010, "username or password is empty"))
		return
	}
	if strings.EqualFold("msuno", username) && strings.EqualFold("123456", password) {
		_, _ = w.Write(retOk("SUCCESS"))
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(retFail(10011, "username or password is incorrect"))
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	setJsonContentType(w)
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
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(retFail(10010, "username or password is empty"))
		return
	}
	if strings.EqualFold("msuno", username) && strings.EqualFold("123456", password) {
		access_token := util.RandId(16)
		refresh_token := util.RandId(16)
		res := make(map[string]interface{})
		res["access_token"] = access_token
		res["refresh_token"] = refresh_token
		res["expired_in"] = redisAccessExpired.Seconds()
		client.Set(buildReflashToken(refresh_token), access_token, redisReflashExpired)
		client.Set(buildAccessToken(access_token), refresh_token, redisAccessExpired)
		_, _ = w.Write(retOk(res))
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(retFail(10011, "username or password is incorrect"))
		return
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	if !filter(w, r) {
		return
	}
	keys := client.HGetAll(redisUserList).Val()
	_, _ = w.Write(retOk(keys))
}

func info(w http.ResponseWriter, r *http.Request) {
	if !filter(w, r) {
		return
	}
	var men runtime.MemStats
	runtime.ReadMemStats(&men)
	m := make(map[string]interface{})
	m["root"] = runtime.GOROOT()
	m["version"] = runtime.Version()
	m["cpu"] = runtime.GOMAXPROCS(0)
	m["goarch"] = runtime.GOARCH
	m["goos"] = runtime.GOOS
	m["alloc"] = fmt.Sprintf("%d Kb", men.Alloc/1024)
	m["frees"] = men.Frees
	m["sys"] = fmt.Sprintf("%d Kb", men.Sys/1024)
	m["gssys"] = fmt.Sprintf("%d Kb", men.GCSys/1024)
	m["nextgc"] = fmt.Sprintf("%d Kb", men.NextGC/1024)
	m["lastgc"] = men.LastGC
	m["numgc"] = men.NumGC
	m["hostname"], _ = os.Hostname()
	_, _ = w.Write(retOk(m))
}

func statistics(w http.ResponseWriter, r *http.Request) {
	if !filter(w, r) {
		return
	}
	buf, err := metrics.Msg()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(retFail(10017, "get metrics msg error"+err.Error()))
		return
	}
	var data map[string]interface{}
	_ = json.Unmarshal(buf, &data)
	_, _ = w.Write(retOk(data))
}

func setting(w http.ResponseWriter, r *http.Request) {
	if !filter(w, r) {
		return
	}
	by, _ := ioutil.ReadAll(r.Body)
	maps := make(map[string]string)
	_ = json.Unmarshal(by, &maps)
	key, _ := maps["key"]
	value, _ := maps["value"]
	if key == "" || value == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(retFail(10015, "key or value is empty"))
		return
	}
	client.HSet(redisUserList, key, value)
	_, _ = w.Write(retOk("SUCCESS"))
}

func del(w http.ResponseWriter, r *http.Request) {
	if !filter(w, r) {
		return
	}
	by, _ := ioutil.ReadAll(r.Body)
	maps := make(map[string]string)
	_ = json.Unmarshal(by, &maps)
	key, _ := maps["key"]
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(retFail(10016, "key is empty"))
		return
	}
	client.HDel(redisUserList, key)
	_, _ = w.Write(retOk("SUCCESS"))
}

func buildAccessToken(token string) string {
	return redisAccessToken + token
}

func buildReflashToken(token string) string {
	return redisRefreshToken + token
}

func retOk(data interface{}) []byte {
	return ret(200, "OK", data)
}

func retFail(status int, data interface{}) []byte {
	return ret(status, "FAIL", data)
}

func ret(status int, message interface{}, data interface{}) []byte {
	maps := make(map[string]interface{})
	maps["status"] = status
	maps["message"] = message
	maps["data"] = data
	by, _ := json.Marshal(maps)
	return by
}
