package main

import (
	"fmt"
	"hw3/stack"
)

func test() {
	s := stack.CreateStack(1, 2, 3, 4, 5)
	fmt.Println(s.Slice)
	s.Push(99)
	fmt.Println(s.Slice)
	val, ok := s.Pop()
	fmt.Println(val, ok, s.Slice)

	if s.IsEmpty() == true {
		fmt.Println("Empty")
	} else {
		fmt.Println("Not Empty")
	}
	s.Clear()
	fmt.Println(s.Slice)
}

func main() {
	test()
}
