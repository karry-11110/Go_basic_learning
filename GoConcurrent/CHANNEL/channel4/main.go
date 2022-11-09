package main

import "fmt"

// func counter(out chan<- int) {
// 	for i := 0; i < 100; i++ {
// 		out <- i
// 	}
// 	close(out)
// }

// func squarer(out chan<- int, in <-chan int) {
// 	for i := range in {
// 		out <- i * i
// 	}
// 	close(out)
// }
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	//定义一个支队其写入int型的单向channel
	go func(out chan<- int) {
		for i := 0; i < 100; i++ {
			out <- i
		}
		close(out)
	}(chan1)
	//定义一个一边读取一边写入的单向channel
	go func(out chan<- int, in <-chan int) {
		for i := range in {
			out <- i * i

		}
		close(out)
	}(chan2, chan1)
	//定义一个主channel
	// go func(in <-chan int) {
	// 	for i := range in {
	// 		fmt.Println(i)
	// 	}

	// }(chan2)
	printer(chan2)
}
