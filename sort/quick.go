package sort

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func qs[T constraints.Ordered](slice []T, low, high int) {
	if low >= high {
		return
	}
	pivot := slice[high]
	fmt.Println(slice)

	pivotIndex := partition(slice, low, high, pivot)

	qs(slice, low, pivotIndex-1)
	qs(slice, pivotIndex+1, high)
}

func partition[T constraints.Ordered](slice []T, low, high int, pivot T) int {
	index := low - 1
	for i := low; i < high; i++ {
		if slice[i] < pivot {
			index++
			slice[i], slice[index] = slice[index], slice[i]
		}
	}

	index++
	slice[index], slice[high] = slice[high], slice[index]
	return index
}

func Quick[T constraints.Ordered](slice []T) {
	qs(slice, 0, len(slice)-1)
}
