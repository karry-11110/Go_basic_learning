package main

import (
	"fmt"
	"strconv"
	"sync"
)

//适用于读多写少的场景
//创建 var 变量名 sync.Map{}
//写入 变量名.Store("key",value)
//读取 value,_=变量名.Load("key")
//删除 变量名.Delete("key")
//读取或写入 变量名.LoadOrStore("key",value)
//  遍历
// 变量名.Range(func(key, value interface{}) bool {
//     name := key.(string)
//     age := value.(int)
//     fmt.Println(name, age)
//     return true
// })
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
