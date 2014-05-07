package utility

import (
	// standard library
	"math/rand"
	"time"
)

// intInArray() determines if the given number is already in the given array.
func intInArray(elem int, array []int) bool {
	for _, b := range array {
		if b == elem {
			return true
		}
	}
	return false
}

// FourRandomNumbersIn() generates 4 random numbers in a range from 0 to given maximum and returns an array.
func FourRandomNumbersIn(area int) (result []int) {
	for len(result) != 4 {
		number := Random(0, area)
		if !intInArray(number, result) {
			result = append(result, number)
		}
	}
	return
}

// Random() generates a random number in a range from 0 to given maximum.
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
