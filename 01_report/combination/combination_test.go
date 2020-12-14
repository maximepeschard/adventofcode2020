package combination

import (
	"reflect"
	"testing"

	"github.com/maximepeschard/adventofcode2020/01_report/arithmetic"
)

func TestCombinationsAll(t *testing.T) {
	cases := []struct {
		k        int
		l        []int
		expected [][]int
	}{
		{1, []int{1, 2, 3}, [][]int{{1}, {2}, {3}}},
		{2, []int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{3, []int{1, 2, 3}, [][]int{{1, 2, 3}}},
		{0, []int{1, 2, 3}, [][]int{{}}},
		{4, []int{1, 2, 3}, [][]int{}},
	}

	for _, c := range cases {
		result := Combinations(c.k, c.l, func([]int) bool { return true })
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Combinations(%d, %v, all) -> %v, expected %v", c.k, c.l, result, c.expected)
		}
	}
}

func TestCombinationsSumToThreeOrSix(t *testing.T) {
	cases := []struct {
		k        int
		l        []int
		expected [][]int
	}{
		{1, []int{1, 2, 3}, [][]int{{3}}},
		{2, []int{1, 2, 3}, [][]int{{1, 2}}},
		{3, []int{1, 2, 3}, [][]int{{1, 2, 3}}},
		{0, []int{1, 2, 3}, [][]int{}},
		{4, []int{1, 2, 3}, [][]int{}},
	}

	for _, c := range cases {
		result := Combinations(c.k, c.l, func(l []int) bool {
			return arithmetic.SumTo(3, l) || arithmetic.SumTo(6, l)
		})
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Combinations(%d, %v, sumToThreeOrSix) -> %v, expected %v", c.k, c.l, result, c.expected)
		}
	}
}
