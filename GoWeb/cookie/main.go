package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("a") //获取客户端cookie并校验
		if err == nil {
			if cookie == "13" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "err",
		})
		c.Abort() // 若验证不通过，不再调用后续的函数处理
		return
	}
}
func main() {
	r := gin.Default() //1.创建路由
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("a", "13", 60, "/", "localhost", false, true) //有服务端给客户端设置cookie,这里注意域名为localhost，不要写成127.0.0.1
		c.String(200, "Login success!!!")
	})
	r.GET("/home", MiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "home",
		})
	})
	r.Run(":8080")
}
