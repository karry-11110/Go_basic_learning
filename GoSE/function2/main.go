package main

import (
	"fmt"
)

//闭包实现
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	//匿名函数两种实现方式
	// add := func(x, y int) {
	// 	fmt.Println(x + y)
	// }
	// add(1, 2)

	// func(a, b int) {
	// 	fmt.Println(a - b)
	// }(3, 2)

	f1, f2 := calc(30)
	fmt.Println(f1(1), f2(1))

}
