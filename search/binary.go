package search

import "golang.org/x/exp/constraints"

// Search for a target using Binary Search algorithm
// The slice has to be Ordered
// ascending param tells the algorithm if the slice is ordered ascending (true) or descending (false)
func Binary[T constraints.Ordered](slice []T, target T, ascending bool) int {
	low := 0
	high := len(slice)

	for low < high {
		middle := low + (high - low) / 2
		value := slice[middle]

		if value == target {
			return middle
		} else if value > target {
			if ascending {
				high = middle
			} else {
				low = middle
			}
		} else if value < target {
			if ascending {
				low = middle + 1
			} else {
				high = middle + 1
			}
		}
	}

	return -1;
}
