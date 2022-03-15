package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ngrok/util"
	"os"
	"runtime"

	"github.com/labstack/echo"
)

func login(c echo.Context) error {
	cc := c.(*RContext)
	i := cc.redis.HMGet("admin-user", "username", "password").Val()
	m := cc.BodyForm()
	if i[0] == m["username"] && i[1] == m["password"] {
		access_token := util.RandId(16)
		refresh_token := util.RandId(16)
		res := make(map[string]interface{})
		res["access_token"] = access_token
		res["refresh_token"] = refresh_token
		res["expired_in"] = util.RedisAccessExpired.Seconds()
		cc.redis.HMSet("admin-user", res)
		return cc.Ok(res)
	}
	return cc.Fail(http.StatusUnauthorized, "error")
}

func logout(c echo.Context) error {
	cc := c.(*RContext)
	cc.redis.HDel("admin-user", "access_token", "refresh_token", "expired_in")
	return cc.Ok("")
}

func userList(c echo.Context) error {
	cc := c.(*RContext)
	return cc.Ok(cc.redis.HGetAll(util.RedisUserList).Val())
}
func userAdd(c echo.Context) error {
	cc := c.(*RContext)
	m := cc.BodyForm()
	key := m["key"]
	value := m["value"]
	cc.redis.HSet(util.RedisUserList, key.(string), value)
	return cc.Ok("")
}
func userDel(c echo.Context) error {
	cc := c.(*RContext)
	m := cc.BodyForm()
	key := m["key"]
	cc.redis.HDel(util.RedisUserList, key.(string))
	return cc.Ok("")
}

func statistics(c echo.Context) error {
	cc := c.(*RContext)
	return cc.Ok("")
}

func sysInfo(c echo.Context) error {
	cc := c.(*RContext)
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
	return cc.Ok(m)
}

func getUser(c echo.Context) error {
	cc := c.(*RContext)
	maps := make(map[string]interface{})
	maps["name"] = "admin"
	maps["role"] = "admin"
	return cc.Ok(maps)
}

func menuList(c echo.Context) error {
	cc := c.(*RContext)
	var maps [10]map[string]interface{}
	str := `[{"id":2,"parentId":0,"name":"Project","path":"/Project","component":"Layout","redirect":"/Project/ProjectList","meta":{"title":"项目管理","icon":"el-icon-phone"}},{"id":20,"parentId":2,"name":"ProjectList","path":"/Project/ProjectList","component":"ProjectList","meta":{"title":"项目列表","icon":"el-icon-goods"}},{"id":21,"parentId":2,"name":"ProjectDetail","path":"/Project/ProjectDetail/: projName","component":"ProjectDetail","meta":{"title":"项目详情","icon":"el-icon-question","activeMenu":"/Project/ProjectList","hidden":true}},{"id":22,"parentId":2,"name":"ProjectImport","path":"/Project/ProjectImport","component":"ProjectImport","meta":{"title":"项目导入","icon":"el-icon-help"}},{"id":3,"parentId":0,"name":"Nav","path":"/Nav","component":"Layout","redirect":"/Nav/SecondNav/ThirdNav","meta":{"title":"多级导航","icon":"el-icon-picture"}},{"id":30,"parentId":3,"name":"SecondNav","path":"/Nav/SecondNav","redirect":"/Nav/SecondNav/ThirdNav","component":"SecondNav","meta":{"title":"二级导航","icon":"el-icon-camera","alwaysShow":true}},{"id":300,"parentId":30,"name":"ThirdNav","path":"/Nav/SecondNav/ThirdNav","component":"ThirdNav","meta":{"title":"三级导航","icon":"el-icon-platform"}},{"id":31,"parentId":3,"name":"SecondText","path":"/Nav/SecondText","redirect":"/Nav/SecondText/ThirdText","component":"SecondText","meta":{"title":"二级文本","icon":"el-icon-opportunity","alwaysShow":true}},{"id":310,"parentId":31,"name":"ThirdText","path":"/Nav/SecondText/ThirdText","component":"ThirdText","meta":{"title":"三级文本","icon":"el-icon-menu"}},{"id":4,"parentId":0,"name":"Components","path":"/Components","component":"Layout","redirect":"/Components/OpenWindowTest","meta":{"title":"组件测试","icon":"el-icon-phone"}},{"id":40,"parentId":4,"name":"OpenWindowTest","path":"/Components/OpenWindowTest","component":"OpenWindowTest","meta":{"title":"选择页","icon":"el-icon-goods"}},{"id":41,"parentId":4,"name":"CardListTest","path":"/Components/CardListTest","component":"CardListTest","meta":{"title":"卡片列表","icon":"el-icon-question-filled"}},{"id":42,"parentId":4,"name":"TableSearchTest","path":"/Components/TableSearchTest","component":"TableSearchTest","meta":{"title":"表格搜索","icon":"el-icon-question-filled"}},{"id":43,"parentId":4,"name":"ListTest","path":"/Components/ListTest","component":"ListTest","meta":{"title":"标签页列表","icon":"el-icon-question-filled"}},{"id":5,"parentId":0,"name":"Permission","path":"/Permission","component":"Layout","redirect":"/Permission/Directive","meta":{"title":"权限管理","icon":"el-icon-phone","alwaysShow":true}},{"id":50,"parentId":5,"name":"Directive","path":"/Permission/Directive","component":"Directive","meta":{"title":"指令管理","icon":"el-icon-goods"}}]`
	json.Unmarshal([]byte(str), &maps)
	return cc.Ok(maps)
}
