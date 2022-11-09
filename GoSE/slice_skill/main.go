package main

import "fmt"

func main() {
	//先定义一个切片
	a := []int{1, 2, 3}
	for index, value := range a {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &a[index])
	}
	fmt.Printf("%v  address %p len %d cap %d\n", a, a, len(a), cap(a))
	//操作

	//复制,凑会产生一个新的切片地址
	// b := append(a[:0:0], a...)

	// b := make([]int, len(a)) // 一次将内存申请到位
	// copy(b, a)

	//剪切
	// b := append(a[:1], a[2:]...)

	//删除
	// b := append(a[:1], a[4:]...)

	//b := a[:5+copy(a[5:], a[6:])]//return copy(a[5:], a[6:]))  the number of elements copied

	//b := a[:len(a)-1] //删掉最后一个元素

	//a[3] = a[len(a)-1] // 将最后一个元素移到索引i处

	// 剪切或删除操作可能引起的内存泄露,如果切片a中的元素是一个指针类型或包含指针字段的结构体类型（需要被垃圾回收），
	// 上面剪切和删除的示例代码会存在一个潜在的内存泄漏问题：一些具有值的元素仍被切片a引用，因此无法被垃圾回收机制回收掉。

	//内部扩张
	//b := append(a[:3], append(make([]int, 3), a[3:]...)...)//扩张的这一部分元素为T类型的零值

	//尾部扩张
	//b := append(a, make([]int, 3)...)

	//过滤
	// n := 0
	// for _, x := range a {
	// 	if x > 5 {
	// 		a[n] = x
	// 		n++
	// 	}
	// }
	// b := a[:n]

	//过滤而不分配内存
	// b := a[:0]
	// for _, x := range a {
	// 	if x < 8 {
	// 		b = append(b, x)
	// 	}
	// }
	// for i := len(b); i < len(a); i++ { //对于必须被垃圾回收的元素，在完成上述操作后可以添加以下代码
	// 	a[i] = nil
	// }

	//插入
	//b := append(a[:3], append([]int{0}, a[3:]...)...)

	//追加
	//b := append(a, 11)//这里注意如果初始指定容量够大，追加后还是能放在底层数组，就不会另外创一个底层数组

	//弹出
	// x := a[len(a)-1]
	// print(x)

	//前插
	// b := append([]int{0}, a...)

	//操作后的切片

	// fmt.Printf("%v  address %p len %d cap %d\n", a, a, len(a), cap(a))
	// fmt.Printf("%v  address %p len %d cap %d\n", b, b, len(b), cap(b))

}
