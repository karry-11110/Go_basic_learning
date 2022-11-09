package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	lock   sync.Mutex
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

//需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来
// 读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
// 当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
func write() {
	// lock.Lock()
	rwlock.Lock() //加写锁，其他foroutine只能等待
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	// lock.Unlock()
	rwlock.Unlock() //解写锁
	wg.Done()
}
func read() {
	// lock.Lock()
	rwlock.RLock() //加读锁,只有加了读锁的gorotinue可以继续获得锁
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwlock.RUnlock()
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))

}
