package main

import "fmt"

func main() {
	//channel基础操作
	var chan1 chan int
	var chan2 chan string
	fmt.Println(chan1)
	fmt.Println(chan2)
	chan3 := make(chan int, 5) //有缓冲的通道
	fmt.Println(chan3)

	//发送
	chan3 <- 10
	//接受
	x := <-chan3
	// <-chan3
	close(chan3)
	// chan3 <- 4
	// 	关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。
	// 	通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

	// 关闭后的通道有以下特点：

	// 对一个关闭的通道再发送值就会导致panic。
	// 对一个关闭的通道进行接收会一直获取值直到通道为空。
	// 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
	// 关闭一个已经关闭的通道会导致panic。
	fmt.Println(x)
	fmt.Println(chan3) //关闭后通道并没有回收还

}
