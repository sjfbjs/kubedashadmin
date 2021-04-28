package config

import (
	"fmt"
	"gin-vue/pkg/setting"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

//全局Oldkubeclient   RestKubeClient
var KubeClient *kubernetes.Clientset

func init() {
	sec, err := setting.Cfg.GetSection("kube")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	kubeHost := sec.Key("HOST").String()
	//k8s地址
	config := &rest.Config{
		//使用代理
		// nohup  kubectl proxy --port=9999 --address='0.0.0.0' --accept-hosts='^*$'  &
		Host: kubeHost,
		//APIPath: "api",
	}
	KubeClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(kubeClient)
	if KubeClient == nil {
		fmt.Println("未获取到k8s链接")
	}
	//list,err:=c.AppsV1().Deployments("default").List(context.Background(),v1.ListOptions{})
	//fmt.Println(list)
	//return  kubeClient
}
