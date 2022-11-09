package main

import "fmt"

var x = 100 //在函数体外的语句只能以关键字开头，var const import等等

//变量
func main() {
	//标准声明格式
	var name string
	var age int
	var isABoy bool
	fmt.Println(name, age, isABoy)
	//批量神明格式
	var (
		a string
		b int16
		c float32
		d bool
	)
	fmt.Println(a, b, c, d)
	//声明同时指定初始值
	var name1 string = "wangkun"
	var age1 int = 22
	var isABoy1 = true
	fmt.Println(name1, age1, isABoy1)
	var name2, age2 = "cny", 18
	fmt.Println(name2, age2)
	//类型自动推导
	var name3, age3 = "wk", 22
	fmt.Println(name3, age3)
	m := 10
	//短变量声明
	fmt.Println(m)
	//匿名变量
	_, name4 := "wk", "cny"
	fmt.Println(name4)

}
