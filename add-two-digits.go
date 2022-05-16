// You are given a two-digit integer n. Return the sum of its digits.

// Example

// For n = 29, the output should be
// solution(n) = 11.
package main

func solution(n int) int {
	return n/10 + n%10
}
