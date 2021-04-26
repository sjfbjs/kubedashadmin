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
	Name       string   `json:"name"`
	NameSpace  string   `json:"namespace"`
	Images     string   `json:"images"`
	NodeName   string   `json:"nodename"`
	IP         []string `jsob:"ip"` // [ podip , nodeip ]
	Status     string   `json:"status"`
	IsReady    bool     `json:"isready"`
	Message    string   `json:"message"`
	CreateTime string   `json:"createtime"`
}
type MyDeploy struct {
	Name    string  `json:"name"`
	Image   string  `json:"image"`
	PodList []MyPod `json:"podlist"`
}

//获取指定命名空间下的pod
func GetPodsByNS(c *gin.Context) {
	namespace := c.DefaultQuery("namespace", "default")
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
			ipList := make([]string, 2)
			////podIp  nodeIp
			ipList[0] = podInfo.Status.PodIP
			ipList[1] = podInfo.Status.HostIP
			myPod.IsReady = podInfo.Status.ContainerStatuses[0].Ready
			myPod.Images = podInfo.Status.ContainerStatuses[0].Image
			myPod.NodeName = podInfo.Spec.NodeName
			myPod.IP = ipList
			myPod.Name = podInfo.Name
			myPod.Status = string(podInfo.Status.Phase)
			myPod.NameSpace = namespace
			myPod.CreateTime = podInfo.CreationTimestamp.String()
			myPodList = append(myPodList, myPod)
		}
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

func GetDeploymentsByNS(c *gin.Context) {
	namespace := c.DefaultQuery("namespace", "default")
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
			for key, v := range deploy.ObjectMeta.Annotations {
				if key == "kubectl.kubernetes.io/last-applied-configuration" {
					specinfo := make(map[string]map[string]map[string]interface{})
					_ = json.Unmarshal([]byte(v), &specinfo)
					tplspec := specinfo["spec"]
					for speckey, sv := range tplspec {
						if speckey == "template" {
							for tsk, tsv := range sv {
								if tsk == "spec" {
									m := tsv.(map[string]interface{})
									containers := m["containers"]
									//类型判断和转换，后面是是否这个类型的布尔
									containerfir, ok := containers.([]interface{})
									if ok {
										contmap := containerfir[0].(map[string]interface{})
										image := contmap["image"]
										fmt.Println(image)
										myDeply.Image = fmt.Sprintf("%v", image)
									}

								}
							}
						}
					}

				}

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

//GetPodByDeployment
