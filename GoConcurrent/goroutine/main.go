package main

import (
	"fmt"
)

func hello() {
	fmt.Println("hello wnagkun")
}
func main() {
	go hello()
	fmt.Println("hello main")
	// time.Sleep(time.Second) //很生硬
}
