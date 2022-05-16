// Given a sequence of integers as an array, determine whether it is possible to obtain a strictly increasing sequence by removing no more than one element from the array.

// Note: sequence a0, a1, ..., an is considered to be a strictly increasing if a0 < a1 < ... < an. Sequence containing only one element is also considered to be strictly increasing.

// Example

// For sequence = [1, 3, 2, 1], the output should be
// solution(sequence) = false.

// There is no one element in this array that can be removed in order to get a strictly increasing sequence.

// For sequence = [1, 3, 2], the output should be
// solution(sequence) = true.

// You can remove 3 from the array to get the strictly increasing sequence [1, 2]. Alternately, you can remove 2 to get the strictly increasing sequence [1, 3].

func solution(a []int) bool {
	c := 0

	l := len(a)
	if l == 2 {
		return true
	}

	for i := 0; i < l-1; i++ {
		if a[i+1] <= a[i] {
			c++
			skipN := i+2 < l && a[i+2] <= a[i]
			skipB := i-1 >= 0 && a[i+1] <= a[i-1]
			if skipB && skipN || c >= 2 {
				return false
			}
		}
	}

	r := c <= 1

	return r
}