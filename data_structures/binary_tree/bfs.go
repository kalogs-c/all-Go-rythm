package binary_tree

func walkPreOrder[T any](n *BinaryNode[T], f func(T)) {
	if n == nil {
		return
	}
	f(n.Value)
	walkPreOrder(n.Left, f)
	walkPreOrder(n.Right, f)
}

func (n *BinaryNode[T]) GetPreOrderPath() []T {
	if n == nil {
		return nil
	}

	path := []T{}
	walkPreOrder(n, func(value T) {
		path = append(path, value)
	})

	return path
}

func walkPostOrder[T any](n *BinaryNode[T], f func(T)) {
	if n == nil {
		return
	}
	walkPostOrder(n.Left, f)
	walkPostOrder(n.Right, f)
	f(n.Value)
}

func (n *BinaryNode[T]) GetPostOrderPath() []T {
	if n == nil {
		return nil
	}

	path := []T{}
	walkPostOrder(n, func(value T) {
		path = append(path, value)
	})

	return path
}

func walkInOrder[T any](n *BinaryNode[T], f func(T)) {
	if n == nil {
		return
	}

	walkInOrder(n.Left, f)
	f(n.Value)
	walkInOrder(n.Right, f)
}

func (n *BinaryNode[T]) GetInOrderPath() []T {
	if n == nil {
		return nil
	}

	path := []T{}
	walkInOrder(n, func(value T) {
		path = append(path, value)
	})

	return path
}
