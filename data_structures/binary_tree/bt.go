package binary_tree

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
