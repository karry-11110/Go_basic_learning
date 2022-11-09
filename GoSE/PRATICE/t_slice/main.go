package main

import "fmt"

// func change(a []int) {
// 	a[0] = 2
// 	fmt.Printf("%p\n", a)
// 	fmt.Println(a)
// }
// func main() {
// 	s := make([]int, 5)
// 	fmt.Println(s)
// 	change(s)
// 	fmt.Println(s)
// 	fmt.Printf("%p\n", s)
// }

// 结果
// [0 0 0 0 0]
// 0xc0000700c0
// [2 0 0 0 0]
// [2 0 0 0 0]
// 0xc0000700c0

// ，不难想到，将切片传入函数中，形参是实参值的拷贝，此时虽然形参和实参本身是不同的内存空间，
// 但他们的值使得各自都指向同一底层数组，因此当形参改变时(添加元素，改变某元素值）
// 	，函数调用后实参同样会被改变(其实是指向的底层数组的改变)，这和其他一些类型传入函数时有很大不同

// func change(a []int) {
// 	a = append(a, 1)
// 	fmt.Printf("%p\n", a)
// 	fmt.Println(a)

// }
// func main() {
// 	s := make([]int, 5)
// 	fmt.Println(s)
// 	change(s)
// 	fmt.Println(s)
// 	fmt.Printf("%p\n", s)
// }
// [0 0 0 0 0]
// 0xc000010280
// [0 0 0 0 0 1]
// [0 0 0 0 0]
// 0xc00000a390
// 当添加元素时，若数组容量不够的话，则将原数组进行扩容
// ，扩容后由于要求数组的地址空间连续，因此原地址不满足条件，
// 会在开辟一片新的空内存区域存放扩容后的数组，因此形参存放的地址改变了，而实参对应的原数组并没有改变。

// package main

// import (
// 	"fmt"
// )

// func change(a []int) {
// 	a = append(a, 1)
// 	fmt.Printf("%p\n", a)
// 	fmt.Println(a)
// }
// func main() {
// 	s := make([]int, 5, 8)//在s初始化时增加其容量：
// 	fmt.Println(s)
// 	change(s)
// 	fmt.Println(s)
// 	fmt.Printf("%p\n", s)
// }

// package main

// import (
// 	"fmt"
// )

func change(a []int) {
	a = append(a, 1)
	fmt.Printf("%p\n", a)
	fmt.Println(a)
}
func main() {
	s := make([]int, 5, 8)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
	fmt.Println(s[:6])
	change(s)
	s = s[:6]
	fmt.Println(s)
	fmt.Printf("%p\n", s)
}

// 只不过s初始化时预定义的长度为5，
// 当数组长度变为6之后s并不能查看到数组后面的元素，需要重新进行切片(实际上是把s的右指针向后移动，增加s的长度)，这样就能看到变化了。
