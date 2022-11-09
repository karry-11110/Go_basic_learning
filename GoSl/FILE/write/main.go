package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("../test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()

	//write和writestring
	// str := "hello王坤"
	// file.write([]byte(str))         //写入字节切片
	// file.writestring("hello world") //直接写入字符数据

	// bufio.newwriter
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello 王坤\n") //将数据线写入writer缓存
	}
	writer.Flush() //将缓存中的内容写入文件

	//ioutil.writefile
	// str := "hello 库里"
	// err := ioutil.WriteFile("../test.txt", []byte(str), 0666)
	// if err != nil {
	// 	fmt.Println("write file failed,err：", err)
	// 	return
	// }

}
