package main // 声明 main 包，表明当前是一个可执行程序

import (
	"fmt" // 导入内置 fmt 包
	mypackage "test_gomod/package"
)

func main() { // main函数，是程序执行的入口
	mypackage.Ne()              //方法名首字母大写
	fmt.Println("Hello World!") // 在终端打印 Hello World!
}
