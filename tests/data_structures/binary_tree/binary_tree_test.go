package binary_tree_test

import (
	"testing"

	"github.com/kalogs-c/all-go-rythm/data_structures/binary_tree"
)

func TestBinaryTreePreOrderPath(t *testing.T) {
	binary_node := binary_tree.NewBinaryNode(1)
	left := binary_node.AddLeft(2)
	right := binary_node.AddRight(3)
	left.AddLeft(4)
	left.AddRight(5)
	right.AddLeft(6)
	right.AddRight(7)

	expected := []int{1, 2, 4, 5, 3, 6, 7}
	actual := binary_node.GetPreOrderPath()
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	}
}

func TestBinaryTreePostOrderPath(t *testing.T) {
	binary_node := binary_tree.NewBinaryNode(1)
	left := binary_node.AddLeft(2)
	right := binary_node.AddRight(3)
	left.AddLeft(4)
	left.AddRight(5)
	right.AddLeft(6)
	right.AddRight(7)

	expected := []int{4, 5, 2, 6, 7, 3, 1}
	actual := binary_node.GetPostOrderPath()
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	}
}

func TestBinaryTreeInOrderPath(t *testing.T) {
	binary_node := binary_tree.NewBinaryNode(1)
	left := binary_node.AddLeft(2)
	right := binary_node.AddRight(3)
	left.AddLeft(4)
	left.AddRight(5)
	right.AddLeft(6)
	right.AddRight(7)

	expected := []int{4, 2, 5, 1, 6, 3, 7}
	actual := binary_node.GetInOrderPath()
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	}
}
