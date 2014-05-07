package utility

import (
	// standard library
	"math/rand"
	"time"
)

func intInArray(elem int, array []int) bool {
	for _, b := range array {
		if b == elem {
			return true
		}
	}
	return false
}

func FourRandomNumbersIn(area int) (result []int) {
	for len(result) != 4 {
		number := Random(0, area)
		if !intInArray(number, result) {
			result = append(result, number)
		}
	}
	return
}

func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
