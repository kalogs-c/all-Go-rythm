package data_structures

import (
	"errors"
)

type qnode[T any] struct {
	value T
	next  *qnode[T]
}

type Queue[T any] struct {
	head   *qnode[T]
	tail   *qnode[T]
	length uint
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *Queue[T]) addLen() {
	q.length++
}

func (q *Queue[T]) subLen() {
	q.length++
}

func (q *Queue[T]) Enqueue(item T) {
	node := qnode[T]{
		value: item,
		next:  nil,
	}

	defer q.addLen()

	if q.length == 0 {
		q.head = &node
		q.tail = &node
		return
	}

	q.tail.next = &node
	q.tail = &node
}

func (q *Queue[T]) Dequeue() (*T, error) {
	if q.length == 0 {
		return nil, errors.New("Queue is empty")
	}

	defer q.subLen()

	node := q.head
	q.head = q.head.next

	return &node.value, nil
}

func (q *Queue[T]) ToSlice() []T {
	if q.length == 0 {
		return []T{}
	}

	slice := make([]T, q.length)
	i := 0
	for node := q.head; node != nil; node = node.next {
		slice[i] = node.value
		i++
	}

	return slice
}

func (q *Queue[T]) Len() uint {
	return q.length
}
