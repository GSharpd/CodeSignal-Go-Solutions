// After becoming famous, the CodeBots decided to move into a new building together. Each of the rooms has a different cost, and some of them are free, but there's a rumour that all the free rooms are haunted! Since the CodeBots are quite superstitious, they refuse to stay in any of the free rooms, or any of the rooms below any of the free rooms.

// Given matrix, a rectangular matrix of integers, where each value represents the cost of the room, your task is to return the total sum of all rooms that are suitable for the CodeBots (ie: add up all the values that don't appear below a 0).

// Example

// For

// matrix = [[0, 1, 1, 2],
//           [0, 5, 0, 0],
//           [2, 0, 3, 3]]
// the output should be
// solution(matrix) = 9

package main

import "fmt"

func main() {
	m := [][]int{{1, 2, 4, 0},
		{1, 0, 4, 2},
		{0, 3, 2, 0}} // Example matrix of int for testing purposes

	sum := solution(m)
	fmt.Println(sum)
}

func solution(m [][]int) int {

	sum := 0
	for x := 0; x < len(m[0]); x++ {
		for y := 0; y < len(m) && m[y][x] != 0; y++ {
			sum += m[y][x]
		}
	}

	return sum
}
