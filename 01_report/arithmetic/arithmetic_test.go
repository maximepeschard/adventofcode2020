package arithmetic

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		l        []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 6},
	}

	for _, c := range cases {
		result := Sum(c.l)
		if result != c.expected {
			t.Errorf("Sum(%v) = %d, expected %d", c.l, result, c.expected)
		}
	}
}

func TestSumTo(t *testing.T) {
	cases := []struct {
		n        int
		l        []int
		expected bool
	}{
		{0, []int{}, true},
		{6, []int{}, false},
		{1, []int{1}, true},
		{6, []int{1}, false},
		{6, []int{1, 2, 3}, true},
	}

	for _, c := range cases {
		result := SumTo(c.n, c.l)
		if result != c.expected {
			t.Errorf("Sum(%d, %v) = %t, expected %t", c.n, c.l, result, c.expected)
		}
	}
}
