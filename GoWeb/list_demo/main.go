package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type Todo struct { //9.定义todo模型，与数据库表建立联系
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn) //这里一定注意，DB已经是全局对象了，不能用简短声明
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
func main() {

	//5.创建数据库sql:CREATE DATABASE bubble
	err := initMySQL() //6.连接数据库
	if err != nil {
		panic(err)
	}
	defer DB.Close()        //7.关闭数据库
	DB.AutoMigrate(&Todo{}) //10.模型与数据库的绑定

	r := gin.Default()                    //1. 加载路由引擎
	r.Static("/static", "static")         //4.告诉gin框架模板文件引用的静态文件去哪里找
	r.LoadHTMLGlob("./templates/*")       //3.指定加载html文件的路径,加載所有文件用什麽函數記清楚
	r.GET("/demo", func(c *gin.Context) { //2.把index.html模板渲染
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1") //8.使用组路由v1，完成待办事项
	{                        //11.完成事项的操作

		//添加事项
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办事项，点击提交，会发请求到这里
			//步骤：从请求中把数据拿出来 存入数据库 发出反响
			var todo Todo
			c.BindJSON(&todo)                             //从请求中把数据拿出来
			if err = DB.Create(&todo).Error; err != nil { //存入数据库
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo) //发出反响
			}

		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询todo这个表里的所有数据
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}

	r.Run(":8080")
}
