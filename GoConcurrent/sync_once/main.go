package main

import (
	"fmt"
	"sync"
)

func main() { // sync.Once 也是 Go 官方的一并发辅助对象，它能够让函数方法只执行一次
	var once sync.Once
	onceFunc := func() {
		fmt.Println("Only once")
	}

	for i := 0; i < 10; i++ {
		once.Do(onceFunc)
	}
}

// 源码//sync.Once 是通过对一个标识值，原子性的修改和加载，来减少锁竞争的
// type Once struct {
// 	done uint32
// 	m    Mutex
//    }

//    func (o *Once) Do(f func()) {
// 	   // 原子加载标识值，判断是否已被执行过
// 	if atomic.LoadUint32(&o.done) == 0 {
// 	 o.doSlow(f)
// 	}
//    }

//    func (o *Once) doSlow(f func()) { // 还没执行过函数
// 	o.m.Lock()
// 	defer o.m.Unlock()
// 	if o.done == 0 { // 再次判断下是否已被执行过函数
// 	 defer atomic.StoreUint32(&o.done, 1) // 原子操作：修改标识值
// 	 f() // 执行函数
// 	}
//    }
