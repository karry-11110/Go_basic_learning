package main

import (
	"fmt"
	"os"
)

func main() {
	//print
	fmt.Print("stephencurry")
	name := "stephecurry"
	fmt.Printf("%s\n", name)
	fmt.Println("stephencurry")

	//fprint
	//Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("../test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name1 := "沙河小王子"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name1)

	//Sprint系列函数会把传入的数据生成并返回一个字符串。
	s1 := fmt.Sprint("沙河小王子")
	name2 := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name2, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)

}
