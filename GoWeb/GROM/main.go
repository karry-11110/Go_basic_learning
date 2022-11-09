package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Id   string
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True") //连接mysql数据库
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&UserInfo{}) //创建表 自动迁移 把结构体和数据表进行对应
	//创建
	u1 := UserInfo{"6", "wangkun", 21} //创建数据行
	db.Create(&u1)
	// //查询
	// var u UserInfo
	// db.First(&u) //把查询的第一行数据给到u
	// fmt.Println(u)
	// var uu UserInfo
	// db.Find(&uu, "id=?", "8") //把查询的第一行数据给到u
	// fmt.Println(uu)
	// //更新
	// db.Model(&u).Update("name", "wk")
	// //删除
	// db.Delete(&u)

}
