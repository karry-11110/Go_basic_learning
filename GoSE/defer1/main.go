package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y)) //defer后跟的函数，如果有参数的话会先计算参数的值
	x = 10
	defer calc("BB", x, calc("B", x, y)) //defer后跟的函数，如果有参数的话会先计算参数的值
	y = 20
}
