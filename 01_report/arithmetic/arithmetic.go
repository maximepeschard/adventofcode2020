package arithmetic

// Sum returns the sum of the items of l.
func Sum(l []int) int {
	sum := 0

	for _, v := range l {
		sum += v
	}

	return sum
}

// SumTo checks if the sum of the items of l is n.
func SumTo(n int, l []int) bool {
	return Sum(l) == n
}
