package main

import (
	"fmt"
	"gin-vue/models"
	"net/http"

	"gin-vue/pkg/setting"
	"gin-vue/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//延迟关闭数据库
	models.CloseDB()
	_ = s.ListenAndServe()
}
