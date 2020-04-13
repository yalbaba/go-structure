package main

import (
	"DataStructures/Stack"
	"fmt"
)

func main() {
	s := Stack.New()

	s.Push("hello")
	s.Push("world")

	//v, ok := s.Pop()
	//fmt.Println(v, ok)
	//
	//v, ok = s.Pop()
	//fmt.Println(v, ok)

	v, ok := s.Pop()
	fmt.Println(v, ok)
}
