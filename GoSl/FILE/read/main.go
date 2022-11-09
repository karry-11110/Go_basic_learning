package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close() //为了防止文件忘记关闭，通常使用defer注册文件关闭语句

	//使用read（）方法读取数据
	//它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF
	// var tmp = make([]byte, 128)
	// num, err := file.Read(tmp)
	// if err == io.EOF {
	// 	fmt.Println("文件读完了")
	// 	return
	// }
	// if err != nil {
	// 	fmt.Println("read file failed ,err:", err)
	// 	return
	// }

	// fmt.Printf("读取了%d的字节数据\n", num)
	// fmt.Println(string(tmp[:num]))

	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))

}
