package data_structures_test

import (
	"testing"

	"github.com/kalogs-c/all-go-rythm/data_structures"
)

func TestQueueEnqueue(t *testing.T) {
	q := data_structures.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	slc := q.ToSlice()
	expected := []int{1, 2, 3}

	for i := 0; i < len(expected); i++ {
		if expected[i] != slc[i] {
			t.Errorf("expected %v, got %v", expected, slc)
		}
	}
}

func TestQueueDequeue(t *testing.T) {
	q := data_structures.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, err := q.Dequeue()
	if err != nil {
		t.Error(err)
	}

	if item != 1 {
		t.Errorf("expected 1, got %v", item)
	}

	slc := q.ToSlice()
	expected := []int{2, 3}

	for i := 0; i < len(expected); i++ {
		if expected[i] != slc[i] {
			t.Errorf("expected %v, got %v", expected, slc)
		}
	}
}
