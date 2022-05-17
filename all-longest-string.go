// Given an array of strings, return another array containing all of its longest strings.

// Example

// For inputArray = ["aba", "aa", "ad", "vcd", "aba"], the output should be
// solution(inputArray) = ["aba", "vcd", "aba"].
package main

func solution(a []string) []string {
	l := len(a[0])
	var b []string

	for _, v := range a {
		if len(v) > l {
			l = len(v)
		}
	}

	for x := 0; x < len(a); x++ {
		if len(a[x]) == l {
			b = append(b, a[x])
		}
	}

	return b
}
