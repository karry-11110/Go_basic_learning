package main

import "fmt"

type Student struct {
	Name string
	Id   int
}

func main() {
	s := make([]Student, 1)
	s[0] = Student{
		Name: "chenchao",
		Id:   111,
	}
	s[0].Id = 222
	fmt.Println(s)
}
