package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //结构体类型，好调用方法

func hello(i int) {
	defer wg.Done() //gorotinue结束就登记-1
	fmt.Println("hello,number:", i)

}
func main() {
	// wg.Add(5)
	wg.Add(10)
	// w3g.Add(15)
	for i := 0; i < 10; i++ {
		// wg.Add(1)
		go hello(i)

		// 	go func(){
		// 		fmt.Println("hello;",i)//不仅是匿名函数还是闭包，所以会i引用出错，到外面找
		// wg.Done()
		// }

		// 	go func(i int) {
		// 		fmt.Println("hello;", i) //不仅是匿名函数还是闭包，所以会i引用出错，到外面找
		// 		wg.Done()
		// 	}(i)
	}

	wg.Wait() // 等待所有登记的goroutine都结束

}
