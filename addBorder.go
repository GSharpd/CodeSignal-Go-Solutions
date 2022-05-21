// Given a rectangular matrix of characters, add a border of asterisks(*) to it.

// Example

// For

// picture = ["abc",
//            "ded"]
// the output should be

// solution(picture) = ["*****",
//                       "*abc*",
//                       "*ded*",
//                       "*****"]

package main

import (
	"fmt"
)

func main() {
	solution([]string{"abc",
		"ded"})
}

func solution(p []string) []string {
	r := make([]string, len(p)+2)
	for i := 0; i < len(p)+2; i++ {
		if i == 0 || i == len(p)+1 {
			for c := 0; c < len(p[0])+2; c++ {
				r[i] += "*"
			}
		} else {
			r[i] = fmt.Sprintf("*"+"%s"+"*", p[i-1])
		}
	}

	return r
}
