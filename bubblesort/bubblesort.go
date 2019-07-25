// Package bubblesort provides a way of sorting a slice using the bubble sort algorithm
package bubblesort

// Exec sorts a slice of integers in increasing order using the bubble sort algorithm
func Exec(s []int) {
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
