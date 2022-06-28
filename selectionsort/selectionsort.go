// Package selectionsort provides a way of sorting a slice using the selection sort algorithm.
package selectionsort

import "golang.org/x/exp/constraints"

// Exec sorts a slice of an ordered type T in increasing order using the selection sort algorithm.
func Exec[T constraints.Ordered](s []T) {
	for currentIndex := 0; currentIndex < len(s); currentIndex++ {
		minIndex := currentIndex

		for i := currentIndex + 1; i < len(s); i++ {
			min := s[minIndex]
			item := s[i]

			if item < min {
				minIndex = i
			}
		}

		s[currentIndex], s[minIndex] = s[minIndex], s[currentIndex]
	}
}
