package sort

type comparable interface {
	float32 | float64 | int | rune
}

func swap[T comparable](slice []T, positionOne int, positionTwo int) {
	slice[positionOne], slice[positionTwo] = slice[positionTwo], slice[positionOne]
}
