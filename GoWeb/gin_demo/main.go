package main

import (
	"github.com/gin-gonic/gin"
)

//为什么使用http/net包就可以实现web服务端，框架的作用就是搭好这个舞台供你来实现
func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang",
	})
}
func main() {
	r := gin.Default()        //返回默认的路由引擎
	r.GET("/hello", sayHello) //指定用户使用GET请求访问helllo时，执行sayHello函数
	//启动服务
	r.Run(":9090") //默认8080端口
}

// func sayHello(w http.ResponseWriter, r *http.Request) { //一个响应，一个请求
// 	file, _ := ioutil.ReadFile("./js.txt")
// 	fmt.Fprintln(w, string(file))
// }
// func main() {
// 	http.HandleFunc("/30", sayHello)         //Handle和HandleFunc函数可以向DefaultServeMux添加处理器。定义如果向浏览器访问地址30，就执行sayhello函数
// 	err := http.ListenAndServe(":9090", nil) //ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。
// 	if err != nil {
// 		fmt.Printf("http serve failed,err:%v\n", err)
// 		return
// 	}
// }
