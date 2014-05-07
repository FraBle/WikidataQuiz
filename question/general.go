package question

import (
	// standard library
	"math/rand"
)

func intInArray(elem int, array []int) bool {
	for _, b := range array {
		if b == elem {
			return true
		}
	}
	return false
}

func fourRandomNumbersIn(area int) (result []int) {
	for len(result) != 4 {
		number := rand.Intn(area)
		if !intInArray(number, result) {
			result = append(result, number)
		}
	}
	return
}
