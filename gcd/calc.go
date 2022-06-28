// Package gcd implements algorithms related to the greatest common divisor.
package gcd

import "golang.org/x/exp/constraints"

// Calc calculates the greatest common divisor of a and b using the Euclidean algorithm.
func Calc[T constraints.Integer](a, b T) T {
	var greaterInt T
	var smallerInt T

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
