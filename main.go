package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"

	"gin-vue/pkg/setting"
	"gin-vue/routers"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket请求ping 返回pong
//noinspection ALL
func Ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
func main() {
	router := routers.InitRouter()
	//演示的ws 路由
	router.GET("ping", Ping)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//延迟关闭数据库
	//models.CloseDB()
	_ = s.ListenAndServe()
}
