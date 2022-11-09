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
	// bufio按行读取示例
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}

}
