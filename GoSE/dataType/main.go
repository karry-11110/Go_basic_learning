package main

import "fmt"

func main() {
	//var a uint8 = 255 //256梯度溢出了
	//b := a >> 2       //无符号数右移补0
	//fmt.Printf("%d, %b,", a, a)
	//fmt.Printf("%d, %b,", b, b)

	//var a int8 = -1 //128梯度溢出了
	//b := a >> 2     //有符号数右移补符号位
	//fmt.Printf("%d, %b,", a, a)
	//fmt.Printf("%d, %b,", b, b)
	//fmt.Printf("%b," 10001010)

	var a int8 = -1
	var b int8 = (-128 / a) + 1
	fmt.Println(b)

	//var d int8 = 127
	//fmt.Printf("%d \n", a) //十进制
	//fmt.Printf("%b \n", a) //十进制转二进制
	//fmt.Printf("%o \n", a) //十进制转八进制
	//fmt.Printf("%x \n", a) //十进制转十六进制

	//var b int = 077
	//fmt.Printf("%o \n", b) // 77
	//
	//// 十六进制  以0x开头
	//var c int = 0xff
	//fmt.Printf("%x \n", c) // ff
	//
	////浮点型
	//fmt.Printf("%f \n", math.Pi)
	//fmt.Printf("%.3f \n", math.Pi)
	//
	////复数
	//var d complex64
	//d = 1 + 2i
	//fmt.Println(d)
	//
	////布尔值
	////字符串
	//
	////类型转换
	//var m, n = 3, 4
	//x := int(math.Sqrt(float64(m*m + n*n))) //math.sqrt()只接受float64
	//fmt.Println(x)

}
