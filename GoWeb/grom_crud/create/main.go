package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//1.定义模型
type NbaInfo struct {
	Id   string //由于指定了ID，默认为主键
	Name string `gorm:"default:'wangkun'"` //通过tag定义字段的默认值，在创建记录时候生成的 SQL 语句会排除没有值或值为 零值 的字段。 在将记录插入到数据库后，Gorm会从数据库加载那些字段的默认值。
	//所有字段的零值, 比如0, "",false或者其它零值，都不会保存到数据库内，但会使用他们的默认值。
	Age int
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True") //连接mysql数据库
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//2.创建表 自动迁移 把结构体和数据表进行对应
	// db.SingularTable(true) //表名禁用复数
	db.AutoMigrate(&NbaInfo{})

	//3.创建
	//在创建记录前查询模型的主键是否存在就是数据库中有没有这一条数据，主键为空就创建记录
	p := NbaInfo{"30", "curry", 33} //代码层面创建一个结构体模型对象，此时数据表是没有的
	// fmt.Println(db.NewRecord(&p))        //返回true就可以create
	db.Create(&p) //防止结构体过大，传指针好点

}
