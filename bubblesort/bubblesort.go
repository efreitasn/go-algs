package bubblesort

// Exec sorts a slice of integers using the bubblesort algorithm
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
