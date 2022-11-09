package main

import "fmt"

//类型断言
func main() {
	var x interface {
	}
	x = "hello 世界"
	value, ok := x.(string)
	if ok {

		fmt.Println(value)
	} else {
		fmt.Println("类型断言失败")
	}
}

// func justifyType(x interface{}) {
// 	switch v := x.(type) {
// 	case string:
// 		fmt.Printf("x is a string，value is %v\n", v)
// 	case int:
// 		fmt.Printf("x is a int is %v\n", v)
// 	case bool:
// 		fmt.Printf("x is a bool is %v\n", v)
// 	default:
// 		fmt.Println("unsupport type！")
// 	}
// }
