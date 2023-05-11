package sort_test

import (
	"testing"

	"github.com/kalogs-c/all-go-rythm/sort"
	"github.com/kalogs-c/all-go-rythm/tests"
)

func TestQuickSort(t *testing.T) {
	unsortedSlice := tests.GenerateRandomIntSlice(200, 100, false)

	sort.Quick(unsortedSlice)

	for i := 0; i < len(unsortedSlice)-1; i++ {
		if unsortedSlice[i] > unsortedSlice[i+1] {
			t.Errorf("Failed to sort integer slice ascending, output: %v\n", unsortedSlice)
		}
	}
}
