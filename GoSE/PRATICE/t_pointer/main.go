package main

import "fmt"

func main() {
	//运行下面肯定报错，因为指针，切片map都是引用型，肯定要初始化
	// var a *int
	// *a = 100
	// fmt.Println(*a)
	// var v map[string]int
	// v["库里"] = 30
	// fmt.Println(v)
	// var b []int
	// b[0] = 1
	// fmt.Println(b)

	// a := new(int)
	// fmt.Printf("%p\n", a)
	// *a = 100
	// fmt.Println(*a)
	// b := 1
	// a = &b
	// fmt.Printf("最后a的地址为：%p\n", a)
	// fmt.Printf("最后a指向的内存值为：%d\n", *a)

	var a *int
	var b = 10
	a = &b
	*a = 10
	fmt.Printf("%p,%d\n", a, a)

}
