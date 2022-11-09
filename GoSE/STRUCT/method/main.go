package main

import "fmt"

//方法
//func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//     函数体
// }

type Person struct {
	name string
	age  int
}

//结构体构造函数
func NewPerson(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}

//不同接收者的接受函数：方法 有值类型接收者和指针类型接收者
//指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。

// 需要修改接收者中的值
// 接收者是拷贝代价比较大的大对象
// // 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
// func (p Person) Dream() {
// 	fmt.Printf("%s的梦想是什么？\n", p.name)
// }

func (p *Person) Dream() {
	p.age = 3

}
func main() {
	p1 := NewPerson("wk", 18)
	fmt.Println(p1.age)
	p1.Dream()
	fmt.Println(p1.age)

}
