// Package gcd implements algorithms related to the greatest common divisor.
package gcd

// Calc calculates the greatest common divisor of a and b using the Euclidean algorithm.
func Calc(a, b int) int {
	var greaterInt int
	var smallerInt int

	if a > b {
		greaterInt = a
		smallerInt = b
	} else {
		greaterInt = b
		smallerInt = a
	}

	for smallerInt != 0 {
		smallerInt, greaterInt = greaterInt%smallerInt, smallerInt
	}

	return greaterInt
}
