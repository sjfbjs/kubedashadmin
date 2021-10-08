package v1

import (
	"fmt"
	"gin-vue/config"
	_ "gin-vue/models"
	"gin-vue/pkg/e"
	_ "gin-vue/pkg/setting"
	_ "gin-vue/pkg/util"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
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
	Name      string           `json:"name"`
	Image     string           `json:"image"`
	PodDetail map[string]int32 `json:"poddetail"`
	Status    string           `json:"status"`
	//PodList []MyPod `json:"podlist"`
}

//获取指定命名空间下的pod
//
func GetPodsByNS(c *gin.Context) {
	namespace := c.Param("namespace")
	//maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS
	list, err := config.KubeClient.CoreV1().Pods(namespace).List(c, metav1.ListOptions{})
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

// todo 新增命名空间获取方法
// todp pageNation
func GetDeploymentsByNS(c *gin.Context) {
	//namespace := c.DefaultQuery("namespace", "default")
	namespace := c.Param("namespace")
	//maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS
	list, err := config.KubeClient.AppsV1().Deployments(namespace).List(c, metav1.ListOptions{})
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
		var myDeployList []MyDeploy
		for _, deploy := range deployList {
			myDeploy := MyDeploy{}
			myDeploy.Name = deploy.Name
			myDeploy.Image = fmt.Sprintf("%s", deploy.Spec.Template.Spec.Containers[0].Image)
			if deploy.Status.UnavailableReplicas > 0 {
				myDeploy.Status = "unhealth"
			} else {
				myDeploy.Status = "health"
			}
			podDetailMap := make(map[string]int32)
			podDetailMap["disiredrs"] = deploy.Status.Replicas
			//   deploy.Status.AvailableReplicas   deploy.Status.ReadyReplicas
			podDetailMap["currentrs"] = deploy.Status.ReadyReplicas
			myDeploy.PodDetail = podDetailMap
			myDeployList = append(myDeployList, myDeploy)
			data["lists"] = myDeployList
			data["total"] = len(list.Items)
		}
	}
	//格式需要修改     deploymentName:xxx count:xxx status:xxxx
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

//GetPodByDeployment
//noinspection ALL
func GetNameSpace(c *gin.Context) {
	var code int
	data := make(map[string]interface{})
	list, err := config.KubeClient.CoreV1().Namespaces().List(c, metav1.ListOptions{})
	if err != nil {
		code = e.ERROR
	} else {
		code = e.SUCCESS
	}
	var nameSpaceList []string
	for _, namespace := range list.Items {
		nameSpaceList = append(nameSpaceList, namespace.Name)
	}
	data["lists"] = nameSpaceList
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetPodDetail(c *gin.Context) {
	var code int
	namespace := c.Param("namespace")
	deployName := c.Param("deployment")
	fmt.Println(namespace, deployName)
	//list,err := config.KubeClient.CoreV1().Pods(namespace).Get(c,podName,metav1.GetOptions{})
	deployInfo, _ := config.KubeClient.AppsV1().Deployments(namespace).Get(c, deployName, metav1.GetOptions{})
	//noinspection GoNilness
	selector, err := metav1.LabelSelectorAsSelector(deployInfo.Spec.Selector)

	if err != nil {
		fmt.Println(err.Error())
		code = e.ERROR
		log.Fatal(err)
	} else {
		code = e.SUCCESS
	}

	//fmt.Println(selector.String())
	listOpts := metav1.ListOptions{
		LabelSelector: selector.String(),
	}
	rs, _ := config.KubeClient.AppsV1().ReplicaSets(namespace).List(c, listOpts)

	data := make(map[string]interface{})
	//podList := make(map[string]interface{})
	//这个是知道只会找到一个
	podSelector, _ := metav1.LabelSelectorAsSelector(rs.Items[0].Spec.Selector)
	//for  _, item := range rs.Items {
	//	//rsName = item.Name
	//	fmt.Println(item.Name)
	//}
	//根据rs获取pod详情
	listPodsOpts := metav1.ListOptions{LabelSelector: podSelector.String()}
	pods, _ := config.KubeClient.CoreV1().Pods(namespace).List(c, listPodsOpts)

	//podList := make(map[string]interface{})
	data["deploydetails"] = pods
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
