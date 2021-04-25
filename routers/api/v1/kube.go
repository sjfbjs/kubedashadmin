package v1

import (
	"context"
	"encoding/json"
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
type MyDeploy struct {
	Name string `json:"name"`
}

//获取指定命名空间下的pod
func GetPodsByNS(c *gin.Context) {
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
	if list == nil {
		fmt.Println("list 为空")
		data["total"] = 0
		data["list"] = ""

	} else {
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

		data["total"] = len(list.Items)

	}

	//格式需要修改     deploymentName:xxx count:xxx status:xxxx
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

type MyAnnotations struct {
	ApiVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	Metadata   string      `json:"metadata"`
	Spec       interface{} `json:"spec"`
}

func GetDeploymentsByNS(c *gin.Context) {
	namespace := c.Query("namespace")
	if namespace == "" {
		namespace = "default"
	}
	//maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS
	ctx := context.Background()
	list, err := config.KubeClient.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
		code = e.ERROR
	}
	if list == nil {
		data["list"] = ""
		fmt.Println("deploylist 为空")
		data["total"] = 0
	} else {
		deployList := list.Items
		//var myPodList []MyPod
		var myDeployList []MyDeploy
		for _, deploy := range deployList {
			myDeply := MyDeploy{}
			myDeply.Name = deploy.Name
			//fmt.Printf("%+v\n", deploy.ObjectMeta.Annotations)
			//fmt.Println("---------------------------\n")
			for _, v := range deploy.ObjectMeta.Annotations {
				specinfo := make(map[string]interface{})
				_ = json.Unmarshal([]byte(v), &specinfo)

				//for _,sv := range specinfo {
				//
				//	tpl := make(map[string]interface{})
				//	_ = json.Unmarshal(sv.([]byte),&tpl)
				//	fmt.Println(sv)
				//
				//}

				fmt.Println(specinfo)
				//fmt.Println(v)
				//fmt.Println(reflect.TypeOf(v))
			}
			myDeployList = append(myDeployList, myDeply)
		}
		//fmt.Println(list.Items,"\n")

		data["lists"] = myDeployList

		data["total"] = len(list.Items)

	}

	//格式需要修改     deploymentName:xxx count:xxx status:xxxx
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
