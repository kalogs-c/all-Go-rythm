package tests

import (
	"math/rand"
	"time"
)

func GenerateRandomIntSlice(size int, nRange int, hasNegatives bool) []int {
	numbers := make([]int, size)

	for i := range numbers {
		randomNumber := rand.Intn(nRange)
		if hasNegatives {
			if rand.Intn(2) == 1 {
				numbers[i] = randomNumber * -1
				continue
			}
		}
		numbers[i] = randomNumber
	}

	return numbers
}

func GenerateRandomFloatSlice(size int, nRange int, hasNegatives bool) []float64 {
	rand.Seed(time.Now().UnixNano())

	numbers := make([]float64, size)

	for i := range numbers {
		randomNumber := (float64(rand.Intn(nRange)) + rand.Float64()) 
		if hasNegatives {
			if rand.Intn(2) == 1 {
				numbers[i] = randomNumber * -1
				continue
			}
		}
		numbers[i] = randomNumber
	}

	return numbers
}

func GenerateRandomRuneSlice(size int) []rune {
	rand.Seed(time.Now().UnixNano())

	runes := make([]rune, size)

	for i := range runes {
		runes[i] = rune(rand.Intn(0x110000))
	}

	return runes
}

func GenerateOrderedIntSlice(size int, ascending bool) []int {
	ints := make([]int, size)

	for i := range ints {
		if ascending {
			ints[i] = i + 1
		} else {
			ints[i] = len(ints) - i
		}
	}

	return ints
}
