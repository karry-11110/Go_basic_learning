package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p, &p, p.name)
	// &{one} 0xc000086000 one
	// &{two} 0xc000180000 two
	// &{three} 0xc000206000 three
	// &{six} 0xc000006030 six
	// &{six} 0xc000180008 six
	// &{six} 0xc000086010 six
}

// func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		// fmt.Printf("Value: %v add: %p \n", v, &v)
		// Value: &{one} add: 0xc000006028
		// Value: &{two} add: 0xc000006028
		// Value: &{three} add: 0xc000006028
		go v.print()
	}
	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		// fmt.Printf("Value: %v add: %p \n", v, &v)
		// Value: {four} add: 0xc0000422c0
		// Value: {five} add: 0xc0000422c0
		// Value: {six} add: 0xc0000422c0
		go v.print()
	}
	time.Sleep(3 * time.Second)
}
