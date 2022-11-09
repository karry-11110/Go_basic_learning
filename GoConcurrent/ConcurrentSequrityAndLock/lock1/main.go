package main

import (
	"fmt"
	"sync"
)

var x int64
var lock sync.Mutex
var wg sync.WaitGroup

func add() {
	for i := 0; i < 15000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //解锁
	}
	wg.Done()
}
func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
