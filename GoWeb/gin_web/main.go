package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //设置路由引擎

	//	模板渲染

	// r.LoadHTMLGlob("./templates/**/*")           //多个模板模板解析,注意用loadhtmlglob
	// r.GET("/posts/index", func(c *gin.Context) { //模板渲染之html渲染
	// 	//http请求
	// 	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
	// 		"title": "posts/index",
	// 	})
	// })
	// r.GET("/users/index", func(c *gin.Context) { //模板渲染之html渲染！！！！！！
	// 	//http请求
	// 	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
	// 		"title": "users/index",
	// 	})
	// })

	// r.SetFuncMap(template.FuncMap{ //设置自定义函数，让模板知道有这个函数,一定要放在解析模板前面，顺序不能错！！！！！
	// 	"curry": func(str string) template.HTML {
	// 		return template.HTML(str)
	// 	},
	// })
	// r.Static("/xxx", "./static")                 //一定要在解析模板前声明加载静态文件：html页面上用到的样式文件，css，js 图片
	// r.LoadHTMLFiles("./templates/diyindex.tmpl") //单个模板解析
	// r.GET("/diyindex", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "diyindex.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>") //模板渲染之自定义函数渲染加上静态文件！！！！！！
	// })

	// r.Static("/xxx", "./static/downstatic")  //一定要在解析模板前声明加载静态文件：html页面上用到的样式文件，css，js 图片
	// r.LoadHTMLGlob("./templates/downhtml/*") //多个模板解析
	// r.GET("/down", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "blog.html", nil) //模板渲染之网上下载的一个前端模板！！！！！！！
	// })

	// r.GET("/json01", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"message": "hello world! map"}) //方式一，自己拼接json类型map,gin.H实质就是一个map,json渲染方式一
	// })
	// r.GET("/json02", func(c *gin.Context) {
	// 	var msg struct { //字段首字母必须大写
	// 		Name    string `json:"name"` //要想返回结构体字段的小写，就要给结构体字段写tag,通过指定tag实现json序列化该字段时的key,做定制化操作
	// 		Message string
	// 		Age     int
	// 	}
	// 	msg.Name = "王坤"
	// 	msg.Message = "hello stephen"
	// 	msg.Age = 18
	// 	c.JSON(http.StatusOK, msg) //方式二，使用结构体渲染json
	// })

	// r.GET("/someXML", func(c *gin.Context) {
	// 	// 方式一：自己拼接JSON
	// 	c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
	// })
	// r.GET("/moreXML", func(c *gin.Context) {
	// 	// 方法二：使用结构体
	// 	type MessageRecord struct {
	// 		Name    string
	// 		Message string
	// 		Age     int
	// 	}
	// 	var msg MessageRecord //注意需要使用具名的结构体类型
	// 	msg.Name = "小王子"
	// 	msg.Message = "Hello world!"
	// 	msg.Age = 18
	// 	c.XML(http.StatusOK, msg)
	// })

	// r.GET("/someYAML", func(c *gin.Context) {
	// 	c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	// })

	// //获取参数

	// r.GET("/getquery", func(c *gin.Context) { //获取querystring参数！！！！！！！
	// 	// name:=c.Query("query")//第一种方式获取浏览器那边发请求携带的querystring参数,通过query来获取携带的参数
	// 	// name := c.DefaultQuery("query", "wangkun") //第二种方式通过query来获取携带的参数，取不到就去后面的默认值
	// 	age := c.DefaultQuery("age", "18")
	// 	name, ok := c.GetQuery("query") //第三种方式，根据取不到，第二个就返回布尔值
	// 	if !ok {
	// 		name = "wangkun"
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"name": name,
	// 		"age":  age,
	// 	})
	// })

	// // //一次请求对应一次响应 获取form参数!!!!!!!
	// r.LoadHTMLGlob("./templates/form/*")
	// r.GET("/login", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "login.html", nil)
	// })
	// r.POST("/login", func(c *gin.Context) { //获取form表单用post提交请求方式！！！！！！
	// 	//username := c.DefaultPostForm("username", "库里")// DefaultPostForm取不到值时会返回指定的默认值
	// 	username := c.PostForm("username")
	// 	password := c.PostForm("password")
	// 	c.HTML(http.StatusOK, "index.html", gin.H{ //输出json结果给调用方
	// 		"Name":     username,
	// 		"Password": password,
	// 	})
	// })

	// r.GET("/path/:name/:age", func(c *gin.Context) { //获取path参数!!!!!!
	// 	name := c.Param("name")
	// 	age := c.Param("age")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"name": name,
	// 		"age":  age,
	// 	})
	// })

	//参数绑定
	//type Login struct {
	//	User     string `form:"user" json:"user" binding:"required"`
	//	Password string `form:"password" json:"password" binding:"required"`
	//}
	//r.GET("/bindquerystring", func(c *gin.Context) { //querystring参数绑定
	//	var l Login
	//	err := c.ShouldBind(&l)
	//	if err == nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"user":     l.User,
	//			"password": l.Password,
	//		})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{
	//			"error": err.Error(),
	//		})
	//	}
	//})
	//
	//r.LoadHTMLGlob("./templates/form/*")
	//r.GET("/bindlogin", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "login.html", nil)
	//})
	//r.POST("/login", func(c *gin.Context) {
	//	var l Login
	//	err := c.ShouldBind(&l)
	//	if err == nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"user":     l.User,
	//			"password": l.Password,
	//		})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{
	//			"error": err.Error(),
	//		})
	//	}
	//})

	//上传文件 也是请求和响应一一对应
	// r.LoadHTMLFiles("./templates/file/index.html")
	// r.GET("/file", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })
	// r.POST("/upload", func(c *gin.Context) { //上传一个文件
	// 	file, _ := c.FormFile("f1") //从请求中读文件，参数和对应html中定义的一样
	// 	//将读取的文件保存到本地（服务端本地），首先要制定保存地址，有两种方式
	// 	// dst:=fmt.Sprintf("./%s",file.Filename)
	// 	dst := path.Join("./", file.Filename)
	// 	_ = c.SaveUploadedFile(file, dst)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": "ok",
	// 	})
	// })

	// r.MaxMultipartMemory = 8 << 20
	// r.LoadHTMLFiles("./templates/file/multiindex.html")
	// r.GET("/file", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "multiindex.html", nil)
	// })
	// r.POST("/upload", func(c *gin.Context) { //上传多个文件
	// 	form, _ := c.MultipartForm() //从请求中读多个文件，无参数
	// 	files := form.File["file"]
	// 	//将读取的文件循环保存到本地（服务端本地），再循环体内制定保存地址
	// 	for index, file := range files {

	// 		dst := fmt.Sprintf("./%s_%d", file.Filename, index)
	// 		_ = c.SaveUploadedFile(file, dst)
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": "ok",
	// 	})
	// })

	//重定向

	// r.GET("/cross", func(c *gin.Context) { //内部、外部重定向均支持
	// 	c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com")
	// })

	// r.GET("/a", func(c *gin.Context) {
	// 	c.Request.URL.Path = "/b" //指定重定向的地址
	// 	r.HandleContext(c)        //，转向指定的路由
	// })
	// r.GET("/b", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": "ok",
	// 	})
	// })

	//路由，给个地址就给你条路
	//不同情况的路由：Get(),Post(),Put(),Delete(),Head(),Any()，NoRoute()
	//路由组：拥有共同URL前缀的路由划分为一个路由组。习惯性一对{}包裹同组的路由，这只是为了看着清晰，你用不用{}包裹功能上没什么区别。

	// userGroup := r.Group("/user")
	// {
	// 	userGroup.GET("/index", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"status": "user get index",
	// 		})
	// 	})
	// 	userGroup.GET("/login", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"status": "user get login",
	// 		})
	// 	})
	// 	userGroup.POST("/login", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"status": "user post login",
	// 		})
	// 	})
	// }

	//中间件
	// 	gin.Default()默认使用了Logger和Recovery中间件，其中：
	// Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
	// Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	// 如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由
	//当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()

	//1.定义中间件，Gin中的中间件必须是一个gin.HandlerFunc类型，一般用闭包形式，因为在执行中间件之前还可以执行一些其他操作
	//2.注册中间件：为全局路由注册中间件：r.use(中间件名称) 为单个路由注册中间件：r.get("XXX",中间件1，中间件2.{有中间件执行先后顺序})
	//    为路由组注册中间件：shopgroup:=r.Group("xxxx",中间件1......)  shopgroup.use(中间件名称)

	// r.Use(m1(true))
	// r.GET("/index", m2())
	// r.GET("/index2", m2(), m1(true), m3())
	//r.Use(myTime)
	//// {}为了代码规范
	//shoppingGroup := r.Group("/shopping")
	//{
	//	shoppingGroup.GET("/index", shopIndexHandler)
	//	shoppingGroup.GET("/home", shopHomeHandler)
	//}
	//r.Run(":8080") //启动server

}

//1.定义中间件，Gin中的中间件必须是一个gin.HandlerFunc类型，一般用闭包形式，因为在执行中间件之前还可以执行一些其他操作
// func m1(doCheck bool) gin.HandlerFunc {
// 	//执行一些其他准备工作
// 	return func(c *gin.Context) {
// 		if doCheck {
// 			c.JSON(http.StatusOK, gin.H{
// 				"msa": "ok",
// 			})
// 		}
// 	}
// }
// func m2() gin.HandlerFunc { //统计请求处理函数的耗时的中间件
// 	//执行一些其他准备工作
// 	return func(c *gin.Context) {
// 		fmt.Println("m1 in")
// 		start := time.Now() //开始计时
// 		c.Next()            //调用后续处理函数，这里直接去别的函数了，千万注意！！！！
// 		//c.Abort()//阻止调用后续的处理函数
// 		cost := time.Since(start)
// 		fmt.Printf("cost:%v\n", cost)
// 		fmt.Println("m1 out")
// 	}
// }
// func m3() gin.HandlerFunc { //统计请求处理函数的耗时的中间件
// 	//执行一些其他准备工作
// 	return func(c *gin.Context) {
// 		fmt.Println("m2 in")
// 		c.Next() //调用后续处理函数，这里直接去别的函数了，千万注意！！！！
// 		fmt.Println("m2 out")
// 	}
// }
func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时：", since)

}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
