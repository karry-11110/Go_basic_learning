package main

import "fmt"

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v %p\n", s, s)
	s1 := make([]int, 3, 4)
	copy(s1, s)
	m["curry"] = s1
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v  %p\n", s, s)
	fmt.Printf("%+v %p\n", m["curry"], m["curry"])
	// 删除切片中的元素其实是把删除元素后面的元素向前一位复制一份，切片的长度len再减一，
	// 1234567变为1256767
	// map[“q1mi”]与slice s指针指向的位置相同，但s长度的长度少1，所以会比map[“q1mi”]少一个元素3。
	s = s[:3]
	fmt.Printf("%+v  %p\n", s, s)

	// a := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Printf("a:%s,长度：%d,地址：%p\n", a, len(a), a)
	// a = append(a[:2], a[3:]...)
	// fmt.Printf("a:%s,长度：%d,地址：%p\n", a, len(a), a)
	// a = a[:(len(a) + 1)]
	// fmt.Printf("a:%s,长度：%d,地址：%p\n", a, len(a), a)

}
