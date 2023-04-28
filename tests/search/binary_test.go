package search_test

import (
	"testing"

	"github.com/kalogs-c/all-go-rythm/search"
	"github.com/kalogs-c/all-go-rythm/tests"
)

func TestBinarySearchOrdered(t *testing.T) {
	slice := tests.GenerateOrderedIntSlice(20, true)

	index := search.Binary(slice, 16, true)

	if index != 15 {
		t.Errorf("Failed to search for 16, expect index 15 but was %v\n", index)
	}
}

func TestBinarySearchOrderedDesc(t *testing.T) {
	slice := tests.GenerateOrderedIntSlice(20, false)

	index := search.Binary(slice, 1, false)

	if index != 19 {
		t.Errorf("Failed to search for 16 in %v, expect index 15 but was %v\n", slice, index)
	}
}
