package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {

	chan3 := make(chan int) //无缓冲的通道
	fmt.Println(chan3)
	go recv(chan3)
	chan3 <- 10
	fmt.Println("发送成功")

}
