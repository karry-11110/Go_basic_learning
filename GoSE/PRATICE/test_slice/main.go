package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
	a := si[3:6]
	fmt.Printf("%v  address %p len %d cap %d\n", a, a, len(a), cap(a))
	si = si[3:6]
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
	si = append(si, 2, 4, 6, 8, 9, 0, 2, 4)
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
	si = append(si[:3], si[4:]...)
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))
	si = si[:(len(si) + 2)]
	fmt.Printf("%v  address %p len %d cap %d\n", si, si, len(si), cap(si))

}
