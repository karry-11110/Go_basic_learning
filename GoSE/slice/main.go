package main

import "fmt"

func main() {
	//声明初始化，这里slice引用类型不用初始化声明了
	// var a []string
	// fmt.Println(a)
	// var b = []int{}
	// c := []bool{false}
	// var d = []bool{false}
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(c == nil)
	//fmt.Println(c==d)//切片为引用类型不支持直接比较，只能和nil比较

	// a := [5]int{1, 2, 3, 4, 5}
	// b := a[1:4]//通过切片表达式获得切片，len和cap求长度和容量函数，容量有扩容策略
	// fmt.Printf("s:%v len(s):%v cap(s):%v \n", b, len(b), cap(b))

	// //通过make函数获得切片
	// a := make([]int64, 3, 10)
	// fmt.Println(a)
	// if len(a) != 0 { //判断切片是否为空
	// 	fmt.Println("切片不为空")
	// }

	// //切片不能通过是否为nil判断为空
	// b1 := []int{}
	// var b2 []int
	// fmt.Println(b1)
	// fmt.Println(b2)

	// //切片的赋值拷贝
	// s1 := []int{1, 2, 3}
	// s2 := make([]int, 4, 5)
	// s2[3] = s1[1]
	// fmt.Println(s1)
	// fmt.Println(s2)
	// s3 := make([]int, 3, 4)
	// s4 := s3 //共用一个底层数组
	// s4[2] = 30
	// fmt.Println(s3)
	// fmt.Println(s4)

	//切片遍历方式
	// s := []int{2, 4, 5, 6}
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(i, s[i])
	// }
	// for index, value := range s {
	// 	fmt.Println(index, 2*value)
	// }

	//append为切片添加元素,扩容策略要了解
	// var a []int //通过var声明的零值切片可以在append函数中直接使用，无需初始化
	// a = append(a, 2, 4, 5, 6)
	// fmt.Println(a)
	// b := []int{10}
	// a = append(a, b...)
	// fmt.Println(a)

	// //使用copy函数复制切片来不改变原来的切片，因为是引用类型
	// a := []int{1, 2, 3}
	// b := make([]int, 3)
	// copy(b, a)
	// fmt.Println(a)
	// fmt.Println(b)
	// b[2] = 30
	// fmt.Println(a)
	// fmt.Println(b)

	//从切片中删除元素
	// a := []int{12, 3, 45, 67, 7}
	// b := append(a[:2], a[3:]...)
	// fmt.Println(b)
	// b[2] = 4
	// fmt.Println(a)
	// fmt.Println(b)

	//参数传递，传slice还是slice指针
	// s := []int{1, 1, 1}
	// newS := myAppend(s)
	// fmt.Println(cap(newS))          //6
	// fmt.Printf("%p %v", s, s)       //0xc000014180 [1 1 1]
	// fmt.Printf("%p %v", newS, newS) //0xc00000a390 [1 1 1 100]
	// myAppendPtr(&newS)
	// fmt.Printf("%p %v", newS, newS) //0xc00000a390 [1 1 1 100 100]
	// myAppendPtr(&s)
	// fmt.Printf("%p %v", s, s) //0xc00000a420 [1 1 1 100]
	// fmt.Println(cap(s))       //6

	s := make([]int, 3, 9) //注意此时创建的slice容量默认和len一样
	s = []int{1, 1}
	fmt.Println(cap(s)) //2
	newS := myAppend(s)
	fmt.Println(cap(newS))          //4
	fmt.Printf("%p %v", s, s)       //0xc0000160a0 [1 1]
	fmt.Printf("%p %v", newS, newS) //0xc00000e1e0 [1 1 100]
	myAppendPtr(&newS)
	fmt.Println(cap(newS))          //4
	fmt.Printf("%p %v", newS, newS) //0xc00000e1e0 [1 1 100 100]
	myAppendPtr(&s)
	fmt.Printf("%p %v", s, s) //0xc00000e220 [1 1 100]
	fmt.Println(cap(s))       //6

}
func myAppend(s []int) []int {
	//这里是slice副本，不会改变外层slice
	s = append(s, 100)
	return s
}
func myAppendPtr(s *[]int) {
	*s = append(*s, 100)
}
