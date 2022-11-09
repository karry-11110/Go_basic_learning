//结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
package main

import "fmt"

//结构体嵌套
type Address struct {
	Province string
	City     string
}
type User struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	user := &User{
		Name: "wangkun",
		Age:  12,
		Address: Address{
			Province: "湖北",
			City:     "武汉",
		},
	}

	fmt.Printf("user=%#v\n", user)

}

//嵌套匿名字段
// //Address 地址结构体
// type Address struct {
// 	Province string
// 	City     string
// }

// //User 用户结构体
// type User struct {
// 	Name    string
// 	Gender  string
// 	Address //匿名字段
// }

// func main() {
// 	var user2 User
// 	user2.Name = "小王子"
// 	user2.Gender = "男"
// 	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
// 	user2.City = "威海"                // 匿名字段可以省略
// 	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
// }
