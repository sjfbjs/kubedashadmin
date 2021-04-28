package config

import (
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

//全局kubeclient
var KubeClient *kubernetes.Clientset

//noinspection ALL
func init() {
	configPath := "conf/kube.conf"
	//k8s地址
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	config.APIPath = "api"
	config.GroupVersion = &coreV1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
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
