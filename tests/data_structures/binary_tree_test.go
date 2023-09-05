package data_structures_test

import (
	"sync"
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

func TestBTBFExec(t *testing.T) {
	binary_node := binary_tree.NewBinaryNode(1)
	left := binary_node.AddLeft(2)
	right := binary_node.AddRight(3)
	left.AddLeft(4)
	left.AddRight(5)
	right.AddLeft(6)
	right.AddRight(7)

	binary_node.BFExec(func(v *binary_tree.BinaryNode[int]) {
		v.Value *= 2
	})

	if binary_node.Value != 2 {
		t.Errorf("expected %v, got %v", 2, binary_node.Value)
	}

	if left.Value != 4 {
		t.Errorf("expected %v, got %v", 4, left.Value)
	}

	if right.Value != 6 {
		t.Errorf("expected %v, got %v", 6, right.Value)
	}

	binary_tree.PrettyPrint(binary_node, "", true)
}

func TestBTBFExecAsync(t *testing.T) {
	binary_node := binary_tree.NewBinaryNode(1)
	left := binary_node.AddLeft(2)
	right := binary_node.AddRight(3)
	left.AddLeft(4)
	left.AddRight(5)
	right.AddLeft(6)
	right.AddRight(7)

	wg := &sync.WaitGroup{}
	binary_node.BFExecAsync(func(v *binary_tree.BinaryNode[int]) {
		v.Value *= 2
	}, wg)
	wg.Wait()

	if binary_node.Value != 2 {
		t.Errorf("expected %v, got %v", 2, binary_node.Value)
	}

	if left.Value != 4 {
		t.Errorf("expected %v, got %v", 4, left.Value)
	}

	if right.Value != 6 {
		t.Errorf("expected %v, got %v", 6, right.Value)
	}

	binary_tree.PrettyPrint(binary_node, "", true)
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

func TestBTDFExec(t *testing.T) {
	binary_node := binary_tree.NewBinaryNode(1)
	left := binary_node.AddLeft(2)
	right := binary_node.AddRight(3)
	left.AddLeft(4)
	left.AddRight(5)
	right.AddLeft(6)
	right.AddRight(7)

	binary_node.DFExec(binary_tree.InOrder, func(v *binary_tree.BinaryNode[int]) {
		v.Value *= 2
	})

	if binary_node.Value != 2 {
		t.Errorf("expected %v, got %v", 2, binary_node.Value)
	}

	if left.Value != 4 {
		t.Errorf("expected %v, got %v", 4, left.Value)
	}

	if right.Value != 6 {
		t.Errorf("expected %v, got %v", 6, right.Value)
	}

	binary_tree.PrettyPrint(binary_node, "", true)
}

func TestBTDFExecAsync(t *testing.T) {
	binary_node := binary_tree.NewBinaryNode(1)
	left := binary_node.AddLeft(2)
	right := binary_node.AddRight(3)
	left.AddLeft(4)
	left.AddRight(5)
	right.AddLeft(6)
	right.AddRight(7)

	wg := &sync.WaitGroup{}
	binary_node.DFExecAsync(func(v *binary_tree.BinaryNode[int]) {
		v.Value *= 2
	}, wg)
	wg.Wait()

	if binary_node.Value != 2 {
		t.Errorf("expected %v, got %v", 2, binary_node.Value)
	}

	if left.Value != 4 {
		t.Errorf("expected %v, got %v", 4, left.Value)
	}

	if right.Value != 6 {
		t.Errorf("expected %v, got %v", 6, right.Value)
	}

	binary_tree.PrettyPrint(binary_node, "", true)
}
