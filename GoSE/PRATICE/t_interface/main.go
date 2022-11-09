package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

// 值接收者实现接口
// ① 值类型结构体实例去调用实现的方法。无需多言，同类型传参当然可以。
// ② 值类型结构体实例去调用实现的方法。也可以。因为Go语言中有对指针类型变量求值的语法糖，Student类型的指针peo内部会自动求值*peo。
// 指针接收者实现接口
// ① 指针类型结构体实例去调用实现的方法。可以，同类型传参当然可以。
// ② 值类型结构体实例去调用实现的方法。不可以，因为类型不同。

func main() {
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
	// var peo People
	// var stu *Student
	// peo = stu
	// fmt.Println(peo.Speak("think"))

}
