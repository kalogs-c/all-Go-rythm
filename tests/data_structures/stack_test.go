package data_structures_test

import (
	"testing"

	ds "github.com/kalogs-c/all-go-rythm/data_structures"
)

func TestStackPushPeek(t *testing.T) {
	s := ds.NewStack[int]()

	length := s.Length()
	if length != 0 {
		t.Errorf("error getting length from stack, expected 0 but was %v", length)
	}

	items := []int{5, 15, 25, 30, 55}

	for _, v := range items {
		s.Push(v)
	}
	length = s.Length()
	if length != 5 {
		t.Errorf("error getting length from stack, expected 1 but was %v", length)
	}

	peek := s.Peek()
	if peek != 55 {
		t.Errorf("error testing peek function, expected %v to be equal %v\n", peek, 55)
	}
}

func TestStackPop(t *testing.T) {
	s := ds.NewStack[int]()

	length := s.Length()
	if length != 0 {
		t.Errorf("error getting length from stack, expected 0 but was %v", length)
	}

	items := []int{5, 15, 25, 30, 55}

	for _, v := range items {
		s.Push(v)
	}
	length = s.Length()
	if length != 5 {
		t.Errorf("error getting length from stack, expected 1 but was %v", length)
	}

	l_item, err := s.Pop()
	if err != nil {
		t.Error(err)
	}
	length = s.Length()
	if *l_item != 55 && length != 4 {
		t.Errorf("error pop stack, expected return 55 and length be 4 but was %v and length %v", l_item, length)
	}

	l_item, err = s.Pop()
	if err != nil {
		t.Error(err)
	}
	length = s.Length()
	if *l_item != 30 || length != 3 {
		t.Errorf("error pop stack, expected return 55 and length be 4 but was %v and length %v", *l_item, length)
	}
}

func TestStackToSlice(t *testing.T) {
	s := ds.NewStack[int]()

	length := s.Length()
	if length != 0 {
		t.Errorf("error getting length from stack, expected 0 but was %v", length)
	}

	items := []int{5, 15, 25, 30, 55}

	for _, v := range items {
		s.Push(v)
	}
	length = s.Length()
	if length != 5 {
		t.Errorf("error getting length from stack, expected 1 but was %v", length)
	}

	slice := s.ToSlice()
	for i, v := range slice {
		if v != items[i] {
			t.Errorf("error parsing stack to slice, expected %v but was %v", items, slice)
		}
	}
}
