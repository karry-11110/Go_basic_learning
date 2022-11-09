package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// var a person //声明一个结构体变量
	//实例化
	// a.age = 18
	// a.name = "wangkun"
	// fmt.Printf("a=%#v\n", a)

	// //匿名结构体
	// var user struct {
	// 	Name string
	// 	Age  int
	// }
	// user.Name = "小王子"
	// user.Age = 18
	// fmt.Printf("%#v\n", user)

	//创建指针类型结构体实例化，go中支持对结构体指针直接使用.来访问成员,
	// var b = new(person)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("b=%#v\n", b)
	// b.age = 23
	// b.name = "cny"
	// fmt.Printf("%T\n", b)
	// fmt.Printf("b=%#v\n", b)

	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	// c := &person{}
	// fmt.Printf("%T\n", c)
	// fmt.Printf("c=%#v\n", c)
	// c.age = 23
	// c.name = "cny"
	// fmt.Printf("%T\n", c)
	// fmt.Printf("c=%#v\n", c)

	//初始化方式

	c := person{name: "wk", age: 12}
	d := &person{name: "cny", age: 13}
	e := &person{name: "curry"}
	f := &person{"stephen", 45} //声明顺序一定要一样
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", d)
	fmt.Printf("%#v\n", e)
	fmt.Printf("%#v\n", f)

}
