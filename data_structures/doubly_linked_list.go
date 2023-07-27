package data_structures

import (
	"fmt"
	"sync"
)

type lk_node[T any] struct {
	value T
	next  *lk_node[T]
	prev  *lk_node[T]
}

type LinkedList[T any] struct {
	head *lk_node[T]
	tail *lk_node[T]
	len  uint
}

func newNode[T any](value T) *lk_node[T] {
	return &lk_node[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (lk *LinkedList[T]) addLen() {
	lk.len++
}

func (lk *LinkedList[T]) subLen() {
	lk.len--
}

func (lk *LinkedList[T]) ToSlice() []T {
	slc := []T{}
	curr := lk.head
	for curr != nil {
		slc = append(slc, curr.value)
		curr = curr.next
	}

	return slc
}

func (lk *LinkedList[T]) Prepend(value T) {
	lk.addLen()

	node := newNode(value)

	if lk.head == nil {
		lk.head, lk.tail = node, node
		return
	}

	currFirst := lk.head

	currFirst.prev, node.next = node, currFirst
	lk.head = node
}

func (lk *LinkedList[T]) Append(value T) {
	lk.addLen()

	node := newNode(value)

	if lk.head == nil {
		lk.head, lk.tail = node, node
		return
	}

	node.prev, lk.tail.next = lk.tail, node
	lk.tail = node
}

func (lk *LinkedList[T]) InsertAt(index uint, value T) error {
	if index > lk.len {
		return fmt.Errorf(
			"index out of bounds, the linked list has lenth of %d but the index is %d\n",
			lk.len,
			index,
		)
	}

	if index == lk.len {
		lk.Append(value)
		return nil
	} else if index == 0 {
		lk.Prepend(value)
		return nil
	}

	lk.addLen()
	curr := *lk.head
	for curr.next != nil {
		curr = *curr.next
	}

	node := newNode(value)
	curr.prev.next = node
	node.prev, node.next = curr.prev, &curr
	curr.prev = node

	return nil
}

func (lk *LinkedList[T]) Remove(validation_func func(curr T) bool) error {
	var target *lk_node[T]

	var curr *lk_node[T]
	for curr = lk.head; curr != nil; curr = curr.next {
		if validation_func(curr.value) {
			target = curr
			break
		}
	}

	if target == nil {
		return fmt.Errorf("item not found")
	}

	lk.subLen()
	if lk.len == 0 {
		lk.head, lk.tail = nil, nil
		return nil
	}

	target_prev := target.prev
	target_next := target.next

	if lk.head == target {
		lk.head = target_next
		target_next.prev = nil
		return nil
	}

	if lk.tail == target {
		lk.tail = target_prev
		target_prev.next = nil
		return nil
	}

	target_prev.next, target_next.prev = target_next, target_prev

	return nil
}

func (lk *LinkedList[T]) ForEach(f func(curr T) T) {
	wg := new(sync.WaitGroup)
	wg.Add(int(lk.len))

	var curr *lk_node[T]
	for curr = lk.head; curr != nil; curr = curr.next {
		go func(curr *lk_node[T]) {
			curr.value = f(curr.value)
			wg.Done()
		}(curr)
	}

	wg.Wait()
}
