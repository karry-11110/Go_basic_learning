package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	//开启一个协程将1-100发送到chan1中
	go func() {
		for i := 0; i < 100; i++ {
			chan1 <- i
		}
		close(chan1)
	}()

	//开启一个协程将从chan1中接受值,并将该值的平方发送到chan2中
	go func() {
		for {
			i, ok := <-chan1 //通道关闭后再取值，ok=false
			if !ok {
				break
			}
			chan2 <- i * i
		}
		close(chan2)
	}()

	//主协程中从chan2中接受之打印
	for i := range chan2 { //通道关闭后会退出 for range循环
		fmt.Println(i)
	}
}
