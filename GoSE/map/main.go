package main

import "fmt"

func main() {
	// var a []int
	// fmt.Println(a)
	a := make([]int, 1)
	a[0] = 1
	fmt.Println(a)
	// var b map[int]string
	// fmt.Println(b)
	var b map[int]string
	b[3] = "wk"
	fmt.Println(b)

	var c *int
	fmt.Println(c)

	//声明初始化方式
	// a := make(map[int]string, 3)
	// a[1] = "wuhan"
	// a[2] = "beijing"
	// a[3] = "shanghai"
	// fmt.Println(a)
	// fmt.Println(a[1])
	// fmt.Printf("type of a :%T \n", a)

	// b := map[string]string{"one": "wuhan", "two": "beijing"}
	// fmt.Println(b)
	// value, ok := b["one"]
	// if ok {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("查无此人")
	// }

	//遍历map
	// a := make(map[int]string, 5)
	// a[1] = "wh"
	// a[2] = "bj"
	// a[3] = "sz"
	// fmt.Println(a)
	// delete(a, 2)
	// fmt.Println(a)
	// for _, value := range a {
	// 	fmt.Println(value)
	// }

	//元素为map类型的切片
	// a := make([]map[int]string, 3)
	// fmt.Println(a)
	// for index, value := range a {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }

	// a[0] = make(map[int]string, 3)
	// a[0][1] = "wh"
	// a[0][2] = "bj"
	// for index, value := range a {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }

	//值为切片类型的map
	// a := make(map[int][]string, 3)
	// fmt.Println(a)
	// value := []string{"wuhan"}
	// a[1] = value
	// fmt.Println(a)

}
