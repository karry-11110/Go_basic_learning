package main

import "fmt"

//defer延迟调用，登记一下，压栈操作
//第一步：返回值赋值  defer  第二步：真正的返回   函数中如果真的存在defer，那执行时机一定是在第一步和第二步之间
// func f1() int { //无命名返回值
// 	x := 5
// 	defer func() {
// 		x++ //修改的是x不是返回值
// 	}()
// 	return x
// }

// func f2() (x int) { //返回值有命名
// 	defer func() {
// 		x++
// 	}()
// 	return 5 //返回值是x
// }

// func f3() (y int) {
// 	x := 5
// 	defer func() {
// 		x++ //defer修改的是x
// 	}()
// 	return x //第一步，返回值=y=x=5  defer修改的x  真正返回
// }
// func f4() (x int) {
// 	defer func(x int) {
// 		x++ //改变的是函数x中的副本
// 	}(x)
// 	return 5
// }

// func f5() (x int) {
// 	defer func(x int) int { //无人接手，还是5
// 		x++ //改变的是函数x中的副本
// 		return x
// 	}(x)
// 	return 5
// }
// func f6() (x int) {
// 	defer func(x *int) { //无人接手，还是5
// 		(*x)++

// 	}(&x)
// 	return 5 //1.返回值=x=5,defer：x=6 返回
// }

func F(n int) func() int {
	return func() int {
		n++
		fmt.Println("里面的：", n)
		return n
	}
}

func main() {
	// fmt.Println("start")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// // fmt.Println("end")
	// fmt.Println(f1())
	// fmt.Println(f2())
	// fmt.Println(f3())
	// fmt.Println(f4())
	// fmt.Println(f5())
	// fmt.Println(f6())

	f := F(5)
	defer func() {
		fmt.Println("第一个：", f())
	}()
	defer fmt.Println("第二个", f()) //defer后跟的函数，如果有参数的话会先计算参数的值
	i := f()
	fmt.Println(i)
}
