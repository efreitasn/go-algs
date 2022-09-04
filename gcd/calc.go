// Package gcd implements algorithms related to the greatest common divisor.
package gcd

import "golang.org/x/exp/constraints"

// Calc calculates the greatest common divisor of a and b using the Euclidean algorithm.
func Calc[T constraints.Integer](a, b T) T {
	if b == 0 {
		return a
	}

	return Calc(b, a%b)
}
