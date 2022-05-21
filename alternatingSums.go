// Several people are standing in a row and need to be divided into two teams. The first person goes into team 1, the second goes into team 2, the third goes into team 1 again, the fourth into team 2, and so on.

// You are given an array of positive integers - the weights of the people. Return an array of two integers, where the first element is the total weight of team 1, and the second element is the total weight of team 2 after the division is complete.

// Example

// For a = [50, 60, 60, 45, 70], the output should be
// solution(a) = [180, 105].

package main

func main() {

}

func solution(a []int) []int {
	res := make([]int, 0)
	t1 := make([]int, 0)
	t2 := make([]int, 0)
	var w1 int
	var w2 int

	for i := 0; i < len(a); i++ {
		if i == len(a)-1 {
			t1 = append(t1, a[i])
		} else {
			t1 = append(t1, a[i])
			t2 = append(t2, a[i+1])
		}
		i++
	}

	for i := 0; i < len(t1); i++ {
		w1 += t1[i]
		if i < len(t2) {
			w2 += t2[i]
		}
	}

	res = append(res, w1)
	res = append(res, w2)

	return res
}

// Easier solution would be assigning to an array by the %2 of i since even numbers are always 0 and odds are always 1
func solution2(a []int) (r [2]int) {
	for i, w := range a {
		r[i%2] += w
	}
	return
}
