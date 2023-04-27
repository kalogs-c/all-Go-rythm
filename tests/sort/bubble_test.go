package sort_test

import (
	"testing"

	"github.com/kalogs-c/all-go-rythm/sort"
	"github.com/kalogs-c/all-go-rythm/tests"
)

func TestBubbleSortIntAscending(t *testing.T) {
	unsortedSlice := tests.GenerateRandomIntSlice(200, 100, true)

	sort.BubbleSort(unsortedSlice, true)

	for i := 0; i < len(unsortedSlice) - 1; i++ {
		if (unsortedSlice[i] > unsortedSlice[i + 1]) {
			t.Errorf("Failed to sort integer slice ascending, output: %v\n", unsortedSlice)
		}
	}
}

func TestBubbleSortIntDescending(t *testing.T) {
	unsortedSlice := tests.GenerateRandomIntSlice(200, 10000, true)

	sort.BubbleSort(unsortedSlice, false)

	for i := 0; i < len(unsortedSlice) - 1; i++ {
		if (unsortedSlice[i] < unsortedSlice[i + 1]) {
			t.Errorf("Failed to sort integer slice descending, output: %v\n", unsortedSlice)
		}
	}
}

func TestBubbleSortFloatAscending(t *testing.T) {
	unsortedSlice := tests.GenerateRandomFloatSlice(200, 10000, true)

	sort.BubbleSort(unsortedSlice, true)

	for i := 0; i < len(unsortedSlice) - 1; i++ {
		if (unsortedSlice[i] > unsortedSlice[i + 1]) {
			t.Errorf("Failed to sort integer slice ascending, output: %v\n", unsortedSlice)
		}
	}
}

func TestBubbleSortFloatDescending(t *testing.T) {
	unsortedSlice := tests.GenerateRandomFloatSlice(200, 10000, true)

	sort.BubbleSort(unsortedSlice, false)

	for i := 0; i < len(unsortedSlice) - 1; i++ {
		if (unsortedSlice[i] < unsortedSlice[i + 1]) {
			t.Errorf("Failed to sort integer slice ascending, output: %v\n", unsortedSlice)
		}
	}
}

func TestBubbleSortRuneAscending(t *testing.T) {
	unsortedSlice := tests.GenerateRandomRuneSlice(200)

	sort.BubbleSort(unsortedSlice, true)

	for i := 0; i < len(unsortedSlice) - 1; i++ {
		if (unsortedSlice[i] > unsortedSlice[i + 1]) {
			t.Errorf("Failed to sort integer slice descending, output: %v\n", unsortedSlice)
		}
	}
}

func TestBubbleSortRuneDescending(t *testing.T) {
	unsortedSlice := tests.GenerateRandomRuneSlice(200)

	sort.BubbleSort(unsortedSlice, false)

	for i := 0; i < len(unsortedSlice) - 1; i++ {
		if (unsortedSlice[i] < unsortedSlice[i + 1]) {
			t.Errorf("Failed to sort integer slice descending, output: %v\n", unsortedSlice)
		}
	}
}
