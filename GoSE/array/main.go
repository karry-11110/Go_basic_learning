package main

import (
	"fmt"
)

func main() {
	//数组定义初始化
	// var testArray [5]int
	// fmt.Println(testArray)

	// var numArray = [3]int{2, 3}
	// var cityArray = [3]string{"wuhan", "csc"}
	// fmt.Println(numArray)
	// fmt.Println(cityArray)

	// flowerArray := [5]float32{3.2, 6.5}
	// fmt.Println(flowerArray)

	// a := [...]int{1, 2, 3}
	// b := [...]string{1: "wh", 3: "bj"}
	// fmt.Println(a)
	// fmt.Println(b)

	//数组遍历
	// a := [...]string{"bj", "sh", "wh"}
	// //for遍历
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	// //for range遍历
	// for index, value := range a {
	// 	fmt.Println(index, value)
	// }

	//多维数组
	a := [3][3]string{
		{"beijing", "tiananmen", "yangrouhuoguo"},
		{"wuhan", "huanghelou", "reganmian"},
		{"shanghai", "chenghuangmiao", "wu"},
	}
	for _, value1 := range a {
		for _, value2 := range value1 {
			fmt.Printf("%s\t", value2)
		}
		fmt.Println()
	}

}
