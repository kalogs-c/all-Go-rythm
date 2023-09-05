package binary_tree

import (
	"errors"
	"sync"

	"github.com/kalogs-c/all-go-rythm/data_structures"
)

// Execute a function on each node
func (n *BinaryNode[T]) BFExec(f func(*BinaryNode[T])) {
	q := data_structures.NewQueue[*BinaryNode[T]]()
	q.Enqueue(n)

	for !q.Empty() {
		v, err := q.Dequeue()
		if err != nil {
			panic(err)
		}
		f(v)
		if n.Left != nil {
			q.Enqueue(n.Left)
		}
		if n.Right != nil {
			q.Enqueue(n.Right)
		}
	}
}

// Execute a function on each node asynchronously
func (n *BinaryNode[T]) BFExecAsync(f func(*BinaryNode[T]), wg *sync.WaitGroup) {
	q := data_structures.NewQueue[*BinaryNode[T]]()
	q.Enqueue(n)

	for !q.Empty() {
		wg.Add(1)
		v, err := q.Dequeue()
		if err != nil {
			wg.Done()
			panic(err)
		}
		go func() {
			f(v)
			wg.Done()
		}()
		if n.Left != nil {
			q.Enqueue(n.Left)
		}
		if n.Right != nil {
			q.Enqueue(n.Right)
		}
	}
}

// Search for a node based on the function
// Returns nil if not found
func (n *BinaryNode[T]) BFS(f func(T) bool) (*BinaryNode[T], error) {
	q := data_structures.NewQueue[*BinaryNode[T]]()
	q.Enqueue(n)

	for !q.Empty() {
		v, err := q.Dequeue()
		if err != nil {
			return nil, err
		}

		if f(v.Value) {
			return v, nil
		}

		if v.Left != nil {
			q.Enqueue(v.Left)
		}

		if v.Right != nil {
			q.Enqueue(v.Right)
		}
	}

	return nil, errors.New("value not found")
}
