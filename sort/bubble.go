package sort

// Sort an array using bubble sort algorithm
// WARNING:	This method mutates the slice
func BubbleSort[T comparable](slice []T, isAscending bool) {
	for i := range slice {
		for j := 0; j < len(slice) - i - 1; j++ {
			var comparation bool

			if (isAscending) {
				comparation = slice[j] < slice[j + 1];
			} else {
				comparation = slice[j] < slice[j + 1];
			}
			if (comparation) {
				swap(slice, j, j + 1);
			}
		}
	}
}
