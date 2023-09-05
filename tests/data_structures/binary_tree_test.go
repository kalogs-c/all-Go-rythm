package binary_tree_test

import (
	"testing"

	"github.com/kalogs-c/all-go-rythm/data_structures/binary_tree"
)

func TestBTBFS(t *testing.T) {
	binary_node := binary_tree.NewBinaryNode(1)
	left := binary_node.AddLeft(2)
	right := binary_node.AddRight(3)
	left.AddLeft(4)
	left.AddRight(5)
	right.AddLeft(6)
	expected := right.AddRight(7)

	actual, err := binary_node.BFS(func(v int) bool {
		return v == 7
	})
	if err != nil {
		t.Error(err)
	}

	if actual.Value != expected.Value {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestBTDFS(t *testing.T) {
	binary_node := binary_tree.NewBinaryNode(1)
	left := binary_node.AddLeft(2)
	right := binary_node.AddRight(3)
	left.AddLeft(4)
	left.AddRight(5)
	right.AddLeft(6)
	expected := right.AddRight(7)

	actual, err := binary_node.DFS(binary_tree.PreOrder, func(v int) bool {
		return v == 7
	})
	if err != nil {
		t.Error(err)
	}

	if actual.Value != expected.Value {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
