package main

import "fmt"

func main() {
	// a := 10
	// b := &a
	// fmt.Printf("type of b:%T \n", b) //b是指针类型变量
	// c := *b                          //指针取值
	// fmt.Printf("type of c：%T,value of c:%v\n", c, c)

	//new和make
	// 	二者都是用来做内存分配的。
	// make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	// 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
	a := new(int)
	fmt.Printf("%T", a)

	b := make(map[int]string, 2)
	fmt.Printf("%T", b)
}
