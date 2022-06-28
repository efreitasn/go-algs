// Package quicksort provides a way of sorting a slice using the quick sort algorithm.
package quicksort

import "golang.org/x/exp/constraints"

// Exec sorts a slice of an ordered type T in increasing order using the quick sort algorithm.
func Exec[T constraints.Ordered](s []T) {
	pivotIndex := len(s) - 1
	currentIndex := 0

	for currentIndex < len(s) && pivotIndex != currentIndex {
		item := s[currentIndex]
		pivot := s[pivotIndex]

		if item > pivot {
			s[currentIndex], s[pivotIndex-1], s[pivotIndex] = s[pivotIndex-1], s[pivotIndex], s[currentIndex]

			pivotIndex--
		} else {
			currentIndex++
		}
	}

	if pivotIndex > 1 {
		Exec(s[0:pivotIndex])
	}

	if pivotIndex < (len(s) - 2) {
		Exec(s[(pivotIndex + 1):])
	}
}
