//结构体的继承
package main

import "fmt"

type Animal struct {
	name string
}

type Dog struct {
	feet    int
	*Animal //通过嵌套匿名结构体实现继承
}

func (a *Animal) move() {

	fmt.Printf("%s会动", a.name)
}

func (d *Dog) wang() {
	fmt.Printf("%s有%d只脚\n", d.name, d.feet)
}

func main() {
	d1 := &Dog{
		feet: 4,
		Animal: &Animal{ //嵌套的是结构体指针
			name: "curry",
		},
	}

	d1.wang()
	d1.move()
}
