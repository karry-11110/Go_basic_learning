package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	ctx, clear := context.WithCancel(context.Background())
// 	message := make(chan int)
// 	go son(ctx, message)
// 	for i := 0; i < 10; i++ {
// 		message <- i
// 	}
// 	fmt.Println("1")
// 	clear() //// 当我们弄完之后调用cancel
// 	fmt.Println("2")
// 	time.Sleep(time.Second)
// 	fmt.Println("主进程结束")
// }
// func son(ctx context.Context, msg chan int) {
// 	t := time.Tick(time.Second)
// 	for _ = range t {
// 		select {
// 		case m := <-msg:
// 			fmt.Printf("收到值%d\n", m)
// 		case <-ctx.Done(): //Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭
// 			fmt.Println("结束")
// 			return // return结束该goroutine，防止泄露
// 		}
// 	}
// }

//withdeadline
// func main() {
// 	d := time.Now().Add(50 * time.Microsecond) //定义了一个50毫秒之后过期的deadline
// 	ctx, cancel := context.WithDeadline(context.Background(), d)
// 	defer cancel()
// 	select { //让主程序陷入等待：等待1秒后打印overslept退出或者等待ctx过期后退出。
// 	case <-time.After(time.Second):
// 		fmt.Println("overslepy")
// 	case <-ctx.Done(): //ctx 50毫秒后就会过期，所以ctx.Done()会先接收到context到期通知，并且会打印ctx.Err()的内容。
// 		fmt.Println(ctx.Err())
// 	}
// }

//withtimeout
var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}
func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
