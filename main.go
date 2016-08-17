package main

import (
	"github.com/astaxie/beego"
	"github.com/sunjun/piServer/mywebsocket"
	"github.com/sunjun/piServer/models"

	_ "github.com/sunjun/piServer/routers"
)

func main() {
	go mywebsocket.RunWebsocket()
	models.InitDB()

	beego.Run()
}
