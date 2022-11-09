package main

import "fmt"

//panic在哪都可能引发
//recover只有在defer调用的函数中有效
func a() {
	fmt.Println("a")
}

func b() {
	defer func() {
		err := recover() //此时引发panic错误即将程序return的时候调用defer后面的函数，而recover将程序错误收集起来，err接受
		if err != nil {  //如果真有err，说明程序真有错误，可以记录一下日志等操作
			fmt.Println("func b error")

		}
	}()
	panic("panic in b") //编译文件没问题，go build
}

func c() {
	fmt.Println("c")
}

func main() {
	a()
	b()
	c()

}
