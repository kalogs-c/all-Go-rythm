package sort

func swap[T comparable](slice []T, positionOne int, positionTwo int) {
	slice[positionOne], slice[positionTwo] = slice[positionTwo], slice[positionOne]
}
