// Package quicksort provides a way of sorting a slice using the quick sort algorithm
package quicksort

// Exec sorts a slice of integers in increasing order using the quick sort algorithm
func Exec(s []int) []int {
	// Base case
	if len(s) <= 1 {
		return s
	}

	var left []int
	var right []int
	// In case there's one or more items equal to the pivot
	var pivotList []int

	pivot := s[len(s)-1]

	for _, item := range s {
		if item > pivot {
			right = append(right, item)
		} else if item < pivot {
			left = append(left, item)
		} else {
			pivotList = append(pivotList, pivot)
		}
	}

	left = Exec(left)
	right = Exec(right)

	// Result
	left = append(left, pivotList...)
	left = append(left, right...)

	return left
}
