package v1

import (
	"context"
	"fmt"
	"gin-vue/config"
	_ "gin-vue/models"
	"gin-vue/pkg/e"
	_ "gin-vue/pkg/setting"
	_ "gin-vue/pkg/util"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

//获取指定命名空间下的pod
func GetPods(c *gin.Context) {
	namespace := c.Query("namespace")
	if namespace == "" {
		namespace = "default"
	}
	//maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS
	ctx := context.Background()
	list, err := config.KubeClient.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
		code = e.ERROR
	}
	data["lists"] = list
	if list == nil {
		fmt.Println("list 为空")
		data["total"] = 0
	} else {
		data["total"] = len(list.Items)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
