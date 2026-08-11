package main

import (
	"ads_course_go/internal/stack"
	"fmt"
)

// MAIN ENTRY POINT
func main() {
	s := stack.NewStack[string]()

	s.Push("Hello")
	s.Push("World")
	s.Push("TS")

	s.DisplayStack()

	v, err := s.Pop()
	fmt.Printf("POP 1: %v\nERROR: %v\n", v, err)

	peekValue, _ := s.Peek()
	fmt.Printf("PEEK: %v\n", peekValue)

	v, err = s.Pop()
	fmt.Printf("POP 2: %v\nERROR: %v\n", v, err)
	v, err = s.Pop()
	fmt.Printf("POP 3: %v\nERROR: %v\n", v, err)

	v, err = s.Pop()
	fmt.Printf("POP 4: %v\nERROR: %v\n", v, err)

	s.Push("Hello")
	s.Push("World")
	s.Push("TS")

	s.DisplayStack()
	s.Clear()
	s.DisplayStack()

	s.Push("Hello")
	s.Push("World")
	s.Push("TS")

	for i, val := range s.Iter() {
		fmt.Println(i, val)
	}

}
