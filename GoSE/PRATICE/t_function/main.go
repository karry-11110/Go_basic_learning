package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() int {
	for i := 0; i < len(users); i++ { //这里要注意我们要用什么方式来循环，参考STRUCT/struct2的例子，是否可以拷贝副本用for range(如果涉及到引用型变量要注意)
		distribution[users[i]] = 0
	}

	for k, v := range distribution {
		// fmt.Printf("%T", k)//string
		for i := 0; i < len(k); i++ {
			// for i := range k {//i对字符串的话就是下标

			fmt.Printf("%T", k[i]) //uint8，为什么是uint类型，
			//组成字符串的每个元素叫”字符“，go语言的字符有uint8类型和rune(int32)类型，要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string
			switch strings.ToLower(fmt.Sprintf("%c", k[i])) { //注意k[i]要转换为输出单个字符的格式
			case "e":
				v += 1
			case "i":
				v += 2
			case "o":
				v += 3
			case "u":
				v += 4
			default:
				v += 0
			}
		}
		fmt.Println(k, "分到：", v)
		coins -= v
	}
	if coins >= 0 {
		return coins
	} else {
		println("钱不够分啊！")
		return 0
	}
}
