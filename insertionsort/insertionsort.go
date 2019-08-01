// Package insertionsort provides a way of sorting a slice using the selection sort algorithm.
package insertionsort

// Exec sorts a slice of ints in increasing order using the insertion sort algorithm.
func Exec(s []int) {
	for currentIndex := 1; currentIndex < len(s); currentIndex++ {
		currentValue := s[currentIndex]
		iToPut := 0

		for i := currentIndex - 1; i >= 0; i-- {
			if currentValue > s[i] {
				iToPut = i + 1

				break
			}
		}

		// Swap
		for m := currentIndex - 1; m >= iToPut; m-- {
			s[m+1] = s[m]
		}

		s[iToPut] = currentValue
	}
}
