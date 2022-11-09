package main

import (
	"fmt"
	"time"
)

func main() {
	//时间类型
	//time.Time类型标识时间
	now := time.Now() //通过time.Now()获得当前时间对象，然后获取时间对象的信息
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %2d:%2d:%2d\n", year, month, day, hour, minute, second)

	//时间操作
	//add
	later := now.Add(time.Hour)
	fmt.Println(later)
	//sub
	now1 := time.Now()
	dura := now.Sub(now1)
	fmt.Println(dura)
	//equl before after
	if now.Equal(now1) {
		fmt.Println("两时间相等")
	} else if now.Before(now1) {
		fmt.Println("now在now1前面")
	} else if now.After(now1) {
		fmt.Println("now在now1后面")
	}

	//定时器
	// ticker := time.Tick(time.Second) //定义了一个1秒的间隔定时器
	// for i := range ticker {
	// 	fmt.Println(i) //每秒都会实现的操作
	// }

	//时间格式化
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

}
