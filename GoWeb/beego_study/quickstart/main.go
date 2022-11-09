package main

import (
	"github.com/beego/beego/v2/server/web"
	"quickstart/controllers"
	_ "quickstart/routers"
)

func main() {
	web.AutoRouter(&controllers.MainController{})
	web.Run()
}
