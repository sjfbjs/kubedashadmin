package v1

import (
	"gin-vue/models"
	"gin-vue/pkg/e"
	"gin-vue/pkg/setting"
	"gin-vue/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取指定命名空间下的pod
func GetPods(c *gin.Context) {
	namespace := c.Query("namespace")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if namespace != "" {
		maps["name"] = namespace
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}




