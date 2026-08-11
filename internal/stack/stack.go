package stack

import (
	"errors"
	"fmt"
	"iter"
	"slices"
)

// STACK CLASS/STRUCT
type Stack[T any] struct {
	value []T
	len   int
}

// CREATE A POINTER TO AN EMPTY STACK STRUCTURE
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		value: make([]T, 0),
		len:   0,
	}
}

// PUBLIC FUNCTION LEN THAT RETURNS LENGTH OF A STACK
func (s *Stack[T]) Len() int {
	return s.len
}

// ADD ELEMENT (T) TO STACK
func (s *Stack[T]) Push(value T) {
	s.value = append(s.value, value)
	s.len++
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.isEmpty() {
		return zero, errors.New("stack is empty")
	}
	s.len--
	result := s.value[len(s.value)-1]
	s.value = s.value[:len(s.value)-1]
	return result, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.isEmpty() {
		return zero, errors.New("stack is empty")
	}
	return s.value[s.len-1], nil
}

func (s *Stack[T]) DisplayStack() {
	if s.isEmpty() {
		fmt.Printf("cannot display: stack is empty")
		return
	}
	for i := len(s.value) - 1; i >= 0; i-- {
		fmt.Printf("%v\n", s.value[i])
	}
}

func (s *Stack[T]) isEmpty() bool {
	return s.len == 0
}

func (s *Stack[T]) Clear() {
	s.value = s.value[:0]
	s.len = 0
}

func (s *Stack[T]) Iter() iter.Seq2[int, T] {
	return slices.Backward(s.value)
}
