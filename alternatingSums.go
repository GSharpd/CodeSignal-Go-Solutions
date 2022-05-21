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
