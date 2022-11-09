package main

import "fmt"

func main() {
	const pi = 3.14
	const (
		pie = 3.142
		wk  = 3
	)
	const (
		a = 1
		b
		c
	)
	const n0 = iota
	const (
		n1 = iota
		n2
		_
		n3 = 34
		n4
	)
	const n5 = iota
	fmt.Println(pi, pie, wk)
	fmt.Println(a, b, c)
	fmt.Println(n0, n1, n2, n3, n4, n5)

}
