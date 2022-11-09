package main

import (
	"fmt"
	"os"
	"time"
)

//创建用户结构体
type user struct {
	username string
	password string
}

//创建对象构造函数
func newUser(username, password string) *user {
	return &user{
		username: username,
		password: password,
	}
}

type Logger interface {
	consoleLog()
	fileLog()
}

func (u *user) consoleLog() { //指针接收者实现接口
	t := time.Now()
	fmt.Printf("用户创建成功！用户名为：%s", (*u).username)
	fmt.Printf("用户创建成功！密码为：%s", (*u).password)
	fmt.Printf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func (u *user) fileLog() {
	t := time.Now()
	file, err := os.OpenFile("./"+(*u).username+".txt", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
	}
	data := "用户创建成功！ 用户名为：" + fmt.Sprintf("%s\n", (*u).username) + "密码是：" + (*u).password + fmt.Sprintf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	_, _ = file.WriteString(data)
	_ = file.Close()

}
func main() {
	var (
		username string
		password string
	)
	fmt.Println("请输入用户名：")
	fmt.Scan(&username)
	fmt.Println("请输入密码：")
	fmt.Scan(&password) //就没返回错误了

	var loggerer Logger
	u := newUser(username, password) //指针创建结构体对象，此时只能用指针接收者实现接口
	loggerer = u
	loggerer.consoleLog()
	loggerer.fileLog()

}
