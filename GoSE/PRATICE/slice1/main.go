package main

import (
	"fmt"
)

func main() {
	var a = make([]string, 5, 10)
	fmt.Printf("%p,%#v", a, a) //0xc0000660a0,[]string{"", "", "", "", ""}
	fmt.Println(len(a), cap(a))
	fmt.Println()
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Printf("%p,%#v", a, a) //0xc000068140,[]string{"", "", "", "", "", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println(len(a), cap(a))
}

//可以注意长度，容量的改变，地址之所以改变是因为扩容重新开劈了空间
