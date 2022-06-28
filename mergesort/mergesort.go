// Package mergesort provides a way of sorting a slice using the merge sort algorithm.
package mergesort

import "golang.org/x/exp/constraints"

func divide[T constraints.Ordered](s []T) (left, right []T) {
	half := len(s) / 2

	return s[0:half], s[half:]
}

func merge[T constraints.Ordered](left, right []T) []T {
	var r []T
	var lIndex int
	var rIndex int

	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] < right[rIndex] {
			r = append(r, left[lIndex])

			lIndex++
		} else {
			r = append(r, right[rIndex])

			rIndex++
		}
	}

	if lIndex != len(left) {
		r = append(r, left[lIndex:]...)
	} else if rIndex != len(right) {
		r = append(r, right[rIndex:]...)
	}

	return r
}

// Exec sorts a slice of an ordered type T in increasing order using the merge sort algorithm.
func Exec[T constraints.Ordered](s []T) []T {
	if len(s) == 1 {
		return s
	}

	left, right := divide(s)
	left, right = Exec(left), Exec(right)

	return merge(left, right)
}
