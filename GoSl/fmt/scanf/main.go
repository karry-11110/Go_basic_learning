package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//scan
	//fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。
	// var (
	// 	name    string
	// 	age     int
	// 	married bool
	// )
	// fmt.Scan(&name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	// //scanfScanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
	// //本函数返回成功扫描的数据个数和遇到的任何错误。
	// fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	// //Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
	// //本函数返回成功扫描的数据个数和遇到的任何错误。
	// fmt.Scanln(&name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	//fscan fscanf fscanfln
	// fmt.Fscanln(os.Stdin, "curry")

	//sscan系列

	//bufio.newreader
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)

}
