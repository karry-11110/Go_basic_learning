```go

	si := []int{1, 2, 3, 4}
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
	change(si) //值传递，在引用函数内是一个副本，函数内指向的是一个新数组
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
}

func change(si []int) {
	si = append(si[:1], si[2:]...)
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
}
[1 2 3 4]  address 0xc00000e1e0 len 4 cap 4
[1 3 4]  address 0xc00000e1e0 len 3 cap 4//副本 但还是改变了底层数组
[1 3 4 4]  address 0xc00000e1e0 len 4 cap 4


func main() {
	si := []int{1, 2, 3, 4}
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
	change(si) //值传递，在引用函数内是一个副本，函数内指向的是一个新数组
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
}
func change(si []int) {
	si = append(si, 5)//副本，扩容问题，开辟了新的底层数组
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
}
[1 2 3 4]  address 0xc00000e1e0 len 4 cap 4
[1 2 3 4 5]  address 0xc00000c380 len 5 cap 8
[1 2 3 4]  address 0xc00000e1e0 len 4 cap 4



func main() {
	si := []int{1, 2, 3, 4}
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
	change(&si) //引用传递，在引用函数内是自己
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
}
func change(si *[]int) {
	(*si) = append((*si), 5)
	fmt.Printf("%v  address %p len %d cap %d\n", (*si), (*si), len(*si), cap(*si))
}
[1 2 3 4]  address 0xc0000a8060 len 4 cap 4
[1 2 3 4 5]  address 0xc0000c40c0 len 5 cap 8
[1 2 3 4 5]  address 0xc0000c40c0 len 5 cap 8


func main() {
	si := []int{1, 2, 3, 4}
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
	change(&si) //值传递，在引用函数内是一个副本，函数内指向的是一个新数组
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
}
func change(si *[]int) {
	(*si) = append((*si)[:2], (*si)[3:]...)
	fmt.Printf("%v  address %p len %d cap %d\n", (*si), (*si), len(*si), cap(*si))
}
[1 2 3 4]  address 0xc00000e1e0 len 4 cap 4
[1 2 4]  address 0xc00000e1e0 len 3 cap 4
[1 2 4]  address 0xc00000e1e0 len 3 cap 4
