package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:]) //数组是值类型必须转换为切片类型才能比较大小
	fmt.Println(a)
}
