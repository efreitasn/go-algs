// Package binarysearch provides a way of finding a value in an ordered slice using the binary search algorithm.
package binarysearch

import "golang.org/x/exp/constraints"

// Exec finds a value in a slice of a comparable type T sorted in increasing order using the binary search algorithm.
func Exec[T constraints.Ordered](s []T, value T) int {
	var (
		left  int
		right = len(s)
	)

	for left < right {
		i := left + (right-left)/2

		switch {
		case s[i] < value:
			left = i + 1
		case s[i] > value:
			right = i
		default:
			return i
		}
	}

	return -1
}
