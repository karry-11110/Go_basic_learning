package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func main() {
	list := make(map[string]Student)
	student := Student{"Aceld"}
	list["student"] = student
	fmt.Println(list["student"].Name)
	list["student"].Name = "LDB"
}
