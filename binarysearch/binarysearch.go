// Package binarysearch provides a way of finding a value in an ordered slice using the binary search algorithm.
package binarysearch

import "golang.org/x/exp/constraints"

// Exec finds a value in a slice of a comparable type T sorted in increasing order using the binary search algorithm.
func Exec[T constraints.Ordered](s []T, value T) (index int, found bool) {
	var start, end int
	halfIndex := len(s) / 2

	for halfIndex < len(s) {
		half := s[halfIndex]

		if value > half {
			start = halfIndex + 1
			end = len(s)
		} else if value < half {
			start = 0
			end = halfIndex
		} else {
			return halfIndex, true
		}

		halfIndex = ((end - start) / 2) + start
	}

	return 0, false
}
