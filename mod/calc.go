// Package mod implements operations related to modular arithmetic.
package mod

import "golang.org/x/exp/constraints"

// Calc calculates a mod p.
// Note that this is not the same as the remainder of a/p.
// The result of a mod p is a number r such that a = qp + r
// for a number q and 0 <= r < p.
// if p < 0, the result is 0.
func Calc[T constraints.Integer](a, p T) T {
	if p < 0 {
		return 0
	}

	res := a % p

	if res < 0 {
		return res + p
	}

	return res
}
