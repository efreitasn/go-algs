// Package bubblesort provides a way of sorting a slice using the bubble sort algorithm.
package bubblesort

import "golang.org/x/exp/constraints"

// Exec sorts a slice of a comparable type T in increasing order using the bubble sort algorithm.
func Exec[T constraints.Ordered](s []T) {
	swap := true

	for swap {
		swap = false

		for i := 0; i < (len(s) - 1); i++ {
			if s[i] > s[i+1] {
				swap = true

				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
}
