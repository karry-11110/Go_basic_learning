package main

import "fmt"

// 关于接口需要注意的是，
// 只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。
// 不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。
//定义接口
type Mover interface {
	move()
	// eat()
}

//定义结构体
type dog struct {
	name string
}

//定义值接受者方法实现接口
// 使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
// 因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui
func (d dog) move() { // 使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
	fmt.Println("狗会动")
}

// func (d *dog) move() { //此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值。
// 	fmt.Println("吃东西")
// }

func main() {
	var x Mover
	var y1 = dog{}
	x = y1 //x接受y类型
	x.move()
	var y2 = &dog{}
	x = y2 //x也可以接受y2类型
	x.move()

}
