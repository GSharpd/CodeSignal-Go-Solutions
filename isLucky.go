// Ticket numbers usually consist of an even number of digits. A ticket number is considered lucky if the sum of the first half of the digits is equal to the sum of the second half.

// Given a ticket number n, determine if it's lucky or not.

// Example

// For n = 1230, the output should be
// solution(n) = true;
// For n = 239017, the output should be
// solution(n) = false.

package main

import "strconv"

func solution(n int) bool {
	strN := strconv.Itoa(n)
	l := len(strN)
	a := make([]int, l)
	var sum1 int
	var sum2 int

	for i := 0; i < l; i++ {
		a[i] = int(strN[i] - '0')
	}

	for x := 0; x < l/2; x++ {
		sum1 += a[x]
	}

	for y := l / 2; y < l; y++ {
		sum2 += a[y]
	}

	return sum1 == sum2
}
