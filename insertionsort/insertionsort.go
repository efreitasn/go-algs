// Package insertionsort provides a way of sorting a slice using the insertion sort algorithm.
package insertionsort

import "golang.org/x/exp/constraints"

// Exec sorts a slice of an ordered type T in increasing order using the insertion sort algorithm.
func Exec[T constraints.Ordered](s []T) {
	for currentIndex := 1; currentIndex < len(s); currentIndex++ {
		currentValue := s[currentIndex]
		i := currentIndex - 1

		for i >= 0 && s[i] > currentValue {
			s[i+1] = s[i]

			i--
		}

		s[i+1] = currentValue
	}
}
