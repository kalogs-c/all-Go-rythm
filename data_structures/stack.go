package data_structures

import (
	"errors"
)

type snode[T any] struct {
	value *T
	prev  *snode[T]
}

type Stack[T any] struct {
	length uint
	tail   *snode[T]
}

func NewStack[T any]() *Stack[T] {
	s := Stack[T]{}

	s.length = 0
	s.tail = nil

	return &s
}

func (s *Stack[T]) Push(item T) {
	node := snode[T]{
		value: &item,
		prev:  nil,
	}

	s.length++
	if s.tail != nil {
		prev := s.tail
		s.tail = &node
		node.prev = prev
	} else {
		s.tail = &node
	}
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.tail == nil {
		return nil, errors.New("Stack is empty")
	}

	s.length--
	item := s.tail.value
	if s.length == 0 {
		s.tail = nil

		return item, nil
	} else {
		s.tail = s.tail.prev
		return item, nil
	}
}

func (s *Stack[T]) ToSlice() []T {
	if s.length == 0 {
		return []T{}
	}

	slice := make([]T, s.length)

	for i := len(slice) - 1; i >= 0; i-- {
		item, err := s.Pop()
		if err == nil {
			slice[i] = *item
		}
	}

	return slice
}

func (s *Stack[T]) Peek() T {
	return *s.tail.value
}

func (s *Stack[T]) Length() uint {
	return s.length
}
