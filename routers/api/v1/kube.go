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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type MyPod struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

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

	//var podList []*v1.Pod
	var podList []v1.Pod
	podList = list.Items
	var myPodList []MyPod

	for _, podInfo := range podList {
		myPod := MyPod{}
		myPod.Name = podInfo.Name
		myPod.Status = string(podInfo.Status.Phase)
		myPodList = append(myPodList, myPod)
	}
	//fmt.Println(list.Items,"\n")

	data["lists"] = myPodList
	if list == nil {
		fmt.Println("list 为空")
		data["total"] = 0
	} else {
		data["total"] = len(list.Items)
	}
	//格式需要修改     deploymentName:xxx count:xxx status:xxxx
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
