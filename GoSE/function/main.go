package main

import "fmt"

//函数参数相邻参数类型相同可以只写一个类型，可变参数类型只能在最后一个，可变参数实际为slice
//函数返回类型可以为slice,多个返回类型用括号，slice可以返回一个nil
func calc(x, y int, z float32, w ...int) (float32, []int) {
	a := (x * y) + int(z)
	sum := 0
	for _, v := range w {
		sum = sum + v //注意全局变量和局部变量
	}

	a = a + sum
	return float32(a), nil
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func main() {
	result, _ := calc(1, 2, 3.2, 5, 6, 7, 8)
	fmt.Println(result)

	type calculation func(int, int) int //定义一个函数类型
	var c calculation = add             //定义一个函数类型，并将add函数给它初始化
	fmt.Printf("type of c:%T\n", c)
	fmt.Println(c(1, 3))

	d := sub
	fmt.Println(d(31, 1))

}
