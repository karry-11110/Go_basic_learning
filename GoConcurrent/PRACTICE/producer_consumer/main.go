package main

import (
	"fmt"
	"time"
)

/*
知识点：
1、开启多个goroutine
2、利用生产数据输入chan，goroutine消费的概念实现
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func workerMain() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启3个goroutine---消费者
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	//开启5个任务--生产者
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	//输出结果---消费者
	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}
func main() {
	workerMain()
	// worker:1 start job:2
	// worker:2 start job:3
	// worker:3 start job:1
	// worker:2 end job:3
	// worker:3 end job:1
	// 6
	// 2
	// worker:2 start job:4
	// worker:3 start job:5
	// worker:1 end job:2
	// 4
	// worker:2 end job:4
	// 8
	// worker:3 end job:5
	// 10
}
