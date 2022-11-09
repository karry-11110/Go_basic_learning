package main

import "fmt"

//因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意
//面试题
type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student) //map
	stus := []student{             //slice
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	fmt.Printf("%p%p%p", stus, stus[0], stus[1])

	for _, stu := range stus { //需要强调的是，range 返回的是每个元素的副本，而不是直接返回对该元素的引用

		m[stu.name] = &stu //因为迭代返回的变量是一个在迭代过程中根据切片依次赋值的新变量，所以 value 的地址总是相同的，
		//要想获取每个元素的地址，需要使用切片变量和索引值（例如上面代码中的 &slice[index]）。
		// for i := 0; i < len(stus); i++{
		// 	m[stus[i].name] = &stus[i]
		// }
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

// func main() {
// 	m := make(map[string]student) //map
// 	stus := []student{            //slice
// 		{name: "小王子", age: 18},
// 		{name: "娜扎", age: 23},
// 		{name: "大王八", age: 9000},
// 	}

// 	for _, stu := range stus {

// 		m[stu.name] = stu ///////????????????????
// 	}

// 	for k, v := range m {
// 		fmt.Println(k, "=>", v.name)
// 	}
// }
