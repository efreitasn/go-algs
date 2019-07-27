// Package selectionsort provides a way of sorting a slice using the selection sort algorithm
package selectionsort

// Exec sorts a slice of integers in increasing order using the selection sort algorithm
func Exec(s []int) {
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
