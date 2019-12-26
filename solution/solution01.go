package solutions

type A []int

func Solution(A []int) int {
	min := A[0]
	max := A[0]
	var sum, count int

	for _, a := range A {
		count += 1
		if min > a {
			sum += min
			min = a
			continue
		}
		if max < a {
			sum += max
			max = a
			continue
		}
		sum += a
	}
	return (sum - 2*A[0]) / (count - 2)
}
