package main

import "fmt"

func main() {
	s := "wangkun王坤"
	for i := 0; i < len(s); i++ { //byte类型
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune类型
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()

	//修改字符串先转换为byte或者rune类型
	bytes := []byte(s)
	bytes[0] = 'W'
	fmt.Println(string(bytes))

	runes := []rune(s)
	runes[8] = '昆'
	fmt.Println(string(runes))

}
