package main

import "fmt"

func main() {
	// //定义空接口
	// var x interface{}
	// s := "Hello 沙河"
	// x = s
	// fmt.Printf("type:%T value:%v\n", x, x)
	// i := 100
	// x = i
	// fmt.Printf("type:%T value:%v\n", x, x)
	// b := true
	// x = b
	// fmt.Printf("type:%T value:%v\n", x, x)

	// 	// 空接口作为函数参数
	// func show(a interface{}) {
	// 	fmt.Printf("type:%T value:%v\n", a, a)
	// }
	var x = make(map[string]interface{}) //使用空接口实现可以保存任意值的字典
	x["name"] = "库里"
	x["age"] = 33
	x["number"] = 30
	fmt.Println(x)

}
