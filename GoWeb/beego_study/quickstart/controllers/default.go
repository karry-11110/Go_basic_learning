package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = " index.tpl"
	//c.Ctx.WriteString(c.Ctx.Input.Param(":id"))
}
func (c *MainController) Allfunc() {
	//fmt.Println(c.Ctx.Input.Params())
}
func (c *MainController) Getfunc() {
	c.Data["Website"] = "wangkun"
	c.Data["Email"] = "qq.com"
	c.TplName = "index.tpl"
}

//func (c *MainController) Get() {
//	c.Ctx.WriteString("hello stephen") //不使用模板，直接输出字符串
//}
