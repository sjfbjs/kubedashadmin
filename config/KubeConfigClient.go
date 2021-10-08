package config

import (
	"fmt"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

//全局kubeclient
var KubeClient *kubernetes.Clientset

//var (
var (
	kubeconfig []byte
	restConf   *rest.Config
	err        error
)

//noinspection ALL
func init() {
	/**
	   type Config struct {
	    Host  string
	    Token string
	    Port  int
	}

	func main()  {

	    clientSet,_:=NewKubernetesClient(&Config{
	       Host:  "192.168.139.134",
	       Port:  6443,
	       Token: "eyJhbGciOiJSUzI1NiIsImtpZCI6Ikg3THcwRzVHTkdQNDREUi1fRjlQOE45NjBUd29pZDQydEs4Tmd1SFJpNnMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkYXNoYm9hcmQtYWRtaW4tdG9rZW4tbWpudmsiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGFzaGJvYXJkLWFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNjdiNGZjM2ItZGEyZS00OTdhLWJmOTEtZGFmYWU3NmU5ZTU3Iiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmUtc3lzdGVtOmRhc2hib2FyZC1hZG1pbiJ9.kW_8iz1GpZPWL2hqDd2Jhkc-rLEX5QPKYrDCmEATyhl_834rmxRg9PJBmRhPY6T7IL58JP9ffUXlF-m65A3H8nOi47dVoAOy9jAPul8C0jS2uZrXYB4zrz_XwXfoonK4lEJtiT86ULd3M3lrUXvEI5kR8ywn3fRBTz5hVRbs0lrgfmFRY_87zELZuBFjSi-pAZTNr_lrAUtBT3Q3h3JyDXHdUJzqoWM-WcszNAZD2wJDV06PpSkNxMOCl6l0BNvUmaY3uLODb5-2yiywasfI9Ue6vKIYEmisNTk48mvbaoIEO34Gg7N1DnvFsO7raoiV_NZ_1KCJDYnxw0jC88Cr0w",
	    })
	 **/

	configPath := "conf/kube.conf"
	// 读kubeconfig文件
	if kubeconfig, err = ioutil.ReadFile(configPath); err != nil {
		log.Fatal(err)
	}
	// 生成rest client配置
	if restConf, err = clientcmd.RESTConfigFromKubeConfig(kubeconfig); err != nil {
		log.Fatal(err)
	}

	//k8s地址
	//config, err := clientcmd.BuildConfigFromFlags("", configPath)
	//config.APIPath = "api"
	//config.GroupVersion = &coreV1.SchemeGroupVersion
	//config.NegotiatedSerializer = scheme.Codecs
	KubeClient, err = kubernetes.NewForConfig(restConf)
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
