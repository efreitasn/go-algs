// Package mergesort provides a way of sorting a slice using the merge sort algorithm
package mergesort

func divide(s []int) (left, right []int) {
	half := len(s) / 2

	return s[0:half], s[half:len(s)]
}

func merge(left, right []int) []int {
	var r []int
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
		r = append(r, left[lIndex:len(left)]...)
	} else if rIndex != len(right) {
		r = append(r, right[rIndex:len(right)]...)
	}

	return r
}

// Exec sorts a slice of integers in increasing order using the merge sort algorithm
func Exec(s []int) []int {
	if len(s) == 1 {
		return s
	}

	left, right := divide(s)
	left, right = Exec(left), Exec(right)

	return merge(left, right)
}
