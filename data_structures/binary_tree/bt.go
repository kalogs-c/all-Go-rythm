package binary_tree

import "fmt"

type BinaryNode[T any] struct {
	Value       T
	Left, Right *BinaryNode[T]
}

func NewBinaryNode[T any](value T) *BinaryNode[T] {
	return &BinaryNode[T]{
		Value: value,
	}
}

func (n *BinaryNode[T]) AddLeft(value T) *BinaryNode[T] {
	newNode := NewBinaryNode(value)
	n.Left = newNode

	return newNode
}

func (n *BinaryNode[T]) AddRight(value T) *BinaryNode[T] {
	newNode := NewBinaryNode(value)
	n.Right = newNode

	return newNode
}

func PrettyPrint[T any](root *BinaryNode[T], prefix string, isLeft bool) {
	if root == nil {
		return
	}

	if isLeft {
		fmt.Printf("%s├── %v\n", prefix, root.Value)
	} else {
		fmt.Printf("%s└── %v\n", prefix, root.Value)
	}

	childPrefix := prefix
	if isLeft {
		childPrefix += "│   "
	} else {
		childPrefix += "    "
	}

	if root.Left != nil || root.Right != nil {
		PrettyPrint(root.Left, childPrefix, true)
		PrettyPrint(root.Right, childPrefix, false)
	}
}
