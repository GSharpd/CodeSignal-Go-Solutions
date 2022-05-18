// Some people are standing in a row in a park. There are trees between them which cannot be moved. Your task is to rearrange the people by their heights in a non-descending order without moving the trees. People can be very tall!

// Example

// For a = [-1, 150, 190, 170, -1, -1, 160, 180], the output should be
// solution(a) = [-1, 150, 160, 170, -1, -1, 180, 190].

package main

import (
	"sort"
)

func solution(a []int) []int {
	r := make([]int, len(a))
	temp := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		if a[i] == -1 {
			r[i] = a[i]
			temp[i] = 2000
		} else {
			temp[i] = a[i]
			r[i] = 0
		}
	}

	sort.Ints(temp)

	pointer := 0
	for i := 0; i < len(r); i++ {
		if r[i] == -1 {
			continue
		} else {
			r[i] = temp[pointer]
			pointer++
		}
	}

	return r
}
