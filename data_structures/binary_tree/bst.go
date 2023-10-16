package binary_tree

import "golang.org/x/exp/constraints"

func BSTInsert[T constraints.Ordered](bst *BinaryNode[T], value T) *BinaryNode[T] {
	if bst == nil {
		return NewBinaryNode(value)
	}

	if value > bst.Value {
		bst.Right = BSTInsert(bst.Right, value)
	} else {
		bst.Left = BSTInsert(bst.Left, value)
	}

	return bst
}

func BSTSearch[T constraints.Ordered](bst *BinaryNode[T], value T) *BinaryNode[T] {
	if bst == nil {
		return nil
	}

	if bst.Value == value {
		return bst
	}

	if value > bst.Value {
		return BSTSearch(bst.Right, value)
	} else {
		return BSTSearch(bst.Left, value)
	}
}

func BSTDelete[T constraints.Ordered](bst *BinaryNode[T], value T) *BinaryNode[T] {
	if bst == nil {
		return nil
	}

	if value > bst.Value {
		bst.Right = BSTDelete(bst.Right, value)
	} else if value < bst.Value {
		bst.Left = BSTDelete(bst.Left, value)
	}

	if bst.Value == value {
		if bst.Left == nil && bst.Right == nil {
			return nil
		} else if bst.Left == nil {
			return bst.Right
		} else if bst.Right == nil {
			return bst.Left
		} else {
			max := BSTMax(bst.Left)
			bst.Value = max.Value
			bst.Left = BSTDelete(bst.Left, max.Value)
			return bst
		}
	}

	return bst
}

func BSTMax[T constraints.Ordered](bst *BinaryNode[T]) *BinaryNode[T] {
	if bst == nil {
		return nil
	}

	if bst.Right == nil {
		return bst
	} else {
		return BSTMax(bst.Right)
	}
}
