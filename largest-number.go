// Given an integer n, return the largest number that contains exactly n digits.

// Example

// For n = 2, the output should be
// solution(n) = 99.

package main

import "math"

func solution(n int) int {
	p := math.Pow10(n)
	r := p - 1

	return int(r)
}
