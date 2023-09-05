package binary_tree

import (
	"errors"
	"sync"
)

type Order uint8

const (
	InOrder Order = iota
	PreOrder
	PostOrder
)

func walkPreOrder[T any](n *BinaryNode[T], f func(T)) {
	if n == nil {
		return
	}
	f(n.Value)
	walkPreOrder(n.Left, f)
	walkPreOrder(n.Right, f)
}

func walkPostOrder[T any](n *BinaryNode[T], f func(T)) {
	if n == nil {
		return
	}
	walkPostOrder(n.Left, f)
	walkPostOrder(n.Right, f)
	f(n.Value)
}

func walkInOrder[T any](n *BinaryNode[T], f func(T)) {
	if n == nil {
		return
	}

	walkInOrder(n.Left, f)
	f(n.Value)
	walkInOrder(n.Right, f)
}

func (n *BinaryNode[T]) DFExec(order Order, f func(T)) {
	switch order {
	case InOrder:
		walkInOrder(n, f)
	case PreOrder:
		walkPreOrder(n, f)
	case PostOrder:
		walkPostOrder(n, f)
	}
}

func walkAsync[T any](n *BinaryNode[T], f func(T), wg *sync.WaitGroup) {
	defer wg.Done()

	if n == nil {
		return
	}

	wg.Add(3)
	go func() {
		f(n.Value)
		wg.Done()
	}()
	go walkAsync(n.Left, f, wg)
	go walkAsync(n.Right, f, wg)
}

func (n *BinaryNode[T]) DFExecAsync(f func(T), wg *sync.WaitGroup) {
	walkAsync(n, f, wg)
}

func searchInOrder[T any](n *BinaryNode[T], f func(T) bool) (*BinaryNode[T], error) {
	if n == nil {
		return nil, errors.New("not found")
	}

	ln, err := searchPostOrder(n.Left, f)
	if err == nil {
		return ln, nil
	}

	if f(n.Value) {
		return n, nil
	}

	rn, err := searchPostOrder(n.Right, f)
	if err == nil {
		return rn, nil
	}

	return nil, errors.New("not found")
}

func searchPreOrder[T any](n *BinaryNode[T], f func(T) bool) (*BinaryNode[T], error) {
	if n == nil {
		return nil, errors.New("not found")
	}

	if f(n.Value) {
		return n, nil
	}

	ln, err := searchPostOrder(n.Left, f)
	if err == nil {
		return ln, nil
	}

	rn, err := searchPostOrder(n.Right, f)
	if err == nil {
		return rn, nil
	}

	return nil, errors.New("not found")
}

func searchPostOrder[T any](n *BinaryNode[T], f func(T) bool) (*BinaryNode[T], error) {
	if n == nil {
		return nil, errors.New("not found")
	}

	ln, err := searchPostOrder(n.Left, f)
	if err == nil {
		return ln, nil
	}

	rn, err := searchPostOrder(n.Right, f)
	if err == nil {
		return rn, nil
	}

	if f(n.Value) {
		return n, nil
	}

	return nil, errors.New("not found")
}

func (n *BinaryNode[T]) DFS(order Order, f func(T) bool) (*BinaryNode[T], error) {
	switch order {
	case InOrder:
		return searchInOrder(n, f)
	case PreOrder:
		return searchPreOrder(n, f)
	case PostOrder:
		return searchPostOrder(n, f)
	}

	return nil, errors.New("not found")
}
