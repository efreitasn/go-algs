// Package insertionsort provides a way of sorting a slice using the selection sort algorithm.
package insertionsort

// Exec sorts a slice of ints in increasing order using the insertion sort algorithm.
func Exec(s []int) {
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
