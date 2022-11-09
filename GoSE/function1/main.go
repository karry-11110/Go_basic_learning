package main

import (
	"errors"
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func calc(x, y int, op func(int, int) int) int { //函数作为参数
	return op(x, y)
}

func do(s string) (func(int, int) int, error) { //函数作为返回值
	switch s {
	case "+":
		return add, nil
	default:
		err := errors.New(("无法识别操作符"))
		return nil, err
	}
}

func main() {
	//fmt.Println(calc(12, 3, add))
	fmt.Println(do("+"))
	fmt.Println(do("-"))
}
