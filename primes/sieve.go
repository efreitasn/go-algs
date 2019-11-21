// Package primes implements algorithms for generating prime numbers.
package primes

import (
	"math"
)

// Sieve generates a list of primes less than or equal to n using the Sieve of Eratosthenes algorithm.
func Sieve(n int) []int {
	nums := make([]bool, n-1)

	for p := 2; p*p <= n; p++ {
		// (p-2) is used because it's a zero-based index and p >= 2.
		// If this condition is true, then p is a composite and the
		// current iteration must be skipped. This way, p is incremented
		// until a prime is found while p^2 <= n.
		if nums[p-2] {
			continue
		}

		// Mark all multiples of p as composite (starting with p^2).
		for m := p; (m * p) <= n; m++ {
			nums[(m*p)-2] = true
		}
	}

	// Prime number theorem
	resultCap := int(math.Ceil(float64(n) / math.Log(float64(n))))
	result := make([]int, 0, resultCap)

	for i, num := range nums {
		if !num {
			result = append(result, i+2)
		}
	}

	return result
}
