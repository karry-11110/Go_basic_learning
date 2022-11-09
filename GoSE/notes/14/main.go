package main

import (
	"fmt"
)

// 使用各自独立的 6 个 slice 来创建 [2][3] 的动态多维数组
func main() {
	x := 2
	y := 4

	table := make([][]int, x)
	table = append(table, []int{2, 4})
	fmt.Println(table)

	for i := range table {
		fmt.Println(i)
		table[i] = make([]int, y)
	}
	fmt.Println(table)
}
