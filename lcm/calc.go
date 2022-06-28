// Package lcm implements algorithms related to the least common multiple.
package lcm

import (
	"github.com/efreitasn/go-algs/gcd"
	"golang.org/x/exp/constraints"
)

// Calc calculates the least common multiple of a and b.
func Calc[T constraints.Integer](a, b T) T {
	return a * b / gcd.Calc(a, b)
}
