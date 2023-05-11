package sort

import "golang.org/x/exp/constraints"

// Sort an array using bubble sort algorithm
// ascending param tells the algorithm to sort ascending (true) or descending (false)
// WARNING:	This method mutates the slice
func Bubble[T constraints.Ordered](slice []T, ascending bool) {
	for i := range slice {
		for j := 0; j < len(slice)-i-1; j++ {
			var comparation bool

			if ascending {
				comparation = slice[j] > slice[j+1]
			} else {
				comparation = slice[j] < slice[j+1]
			}
			if comparation {
				swap(slice, j, j+1)
			}
		}
	}
}
