// Write a function that reverses characters in (possibly nested) parentheses in the input string.

// Input strings will always be well-formed with matching ()s.

// Example

// For inputString = "(bar)", the output should be
// solution(inputString) = "rab";
// For inputString = "foo(bar)baz", the output should be
// solution(inputString) = "foorabbaz";
// For inputString = "foo(bar)baz(blim)", the output should be
// solution(inputString) = "foorabbazmilb";
// For inputString = "foo(bar(baz))blim", the output should be
// solution(inputString) = "foobazrabblim".
// Because "foo(bar(baz))blim" becomes "foo(barzab)blim" and then "foobazrabblim".

package main

import "strings"

func solution(s string) string {
	// We keep running the function while
	// there's an opening bracket in the string
	for strings.Index(s, "(") != -1 {
		s = reverseP(s)
	}
	return s
}

func reverseP(s string) string {
	len1 := len(s)
	// We find the innermost/last opening bracket,
	// and then we're going to find the matching closing bracket,
	// which is the next closing bracket after the opening bracket
	openBracketInd := strings.LastIndex(s, "(")

	// The characters before the opening bracket
	beforeBracket := s[0 : openBracketInd+1]

	// The characters before the opening bracket
	afterBracket := s[openBracketInd+1 : len1]

	// To get the index of the closing bracket we add the
	// characters before the bracket to the index of the first closing
	// bracket in the string after the opening bracket
	closeBracketInd := len(beforeBracket) + strings.Index(afterBracket, ")")

	// Once we have the indexes, we're going to slice the string
	// to remove the brackets
	firstPart := s[0:openBracketInd]
	secondPart := s[closeBracketInd+1 : len1]
	middle := s[openBracketInd+1 : closeBracketInd]

	// We reverse the string between the brackets
	middle = reverse(middle)

	// And finally we join the parts and return the string
	return firstPart + middle + secondPart
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
