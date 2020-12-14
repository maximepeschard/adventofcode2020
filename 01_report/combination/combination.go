package combination

/*

Finding all k-combinations of n items without repetion :

The idea of the algorithm used in Combinations is to enumerate all
ths combinations of k indices in lexicographic order.

For example, for an array of n = 5 elements and k = 3 :
0 1 2 <- Initial indices. Third index can be increased.
0 1 3 <- Third index can be increased.
0 1 4 <- Third index can *not* be increased. Second index can be increased.
0 2 3 <- Second index increased and third "reset" to the following integer. Etc
0 2 4
0 3 4
1 2 3
1 2 4
1 3 4
2 3 4

*/

// Combinations returns all k-combinations of items of l (without repetition) verifying predicate p.
func Combinations(k int, l []int, p func([]int) bool) [][]int {
	combinations := [][]int{}
	n := len(l)

	if k > n {
		return combinations
	}

	// Initial indexes are [0 1 ... k-1]
	indexes := make([]int, k)
	for i := 0; i < k; i++ {
		indexes[i] = i
	}

	for {
		// The current combination is formed by taking
		// the values of l for the current indexes.
		combination := make([]int, k)
		for i, index := range indexes {
			combination[i] = l[index]
		}
		if p(combination) {
			combinations = append(combinations, combination)
		}

		// The next set of indexes if formed by :
		// - finding the rightmost index we can increase,
		// - increasing its value (k -> k'),
		// - resetting the indexes after it to the next integer values (k'+1, k'+2, ...)
		rightmostIncreasable := -1
		for i := k - 1; i >= 0; i-- {
			if indexes[i] < n-k+i {
				rightmostIncreasable = i
				break
			}
		}

		if rightmostIncreasable == -1 {
			// Nothing left to increase : we have all possible combinations.
			break
		}

		indexes[rightmostIncreasable]++
		for i := rightmostIncreasable + 1; i < k; i++ {
			indexes[i] = indexes[i-1] + 1
		}
	}

	return combinations
}
