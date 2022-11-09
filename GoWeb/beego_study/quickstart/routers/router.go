package routers

func init() {
	//固定路由
	//web.Router("/", &controllers.MainController{})
	//正则路由
	//web.Router("/?:id", &controllers.MainController{})
	//web.Router("/:id", &controllers.MainController{})
	//web.Router("/:id([0-9]+)", &controllers.MainController{})
	//自定义方法及restful规则
	//web.Router("/:id", &controllers.MainController{}, "*:Allfunc;get,post:Getfunc")

}
