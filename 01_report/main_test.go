package main

import (
	"reflect"
	"testing"

	"github.com/maximepeschard/adventofcode2020/01_report/arithmetic"
	"github.com/maximepeschard/adventofcode2020/01_report/combination"
)

func TestPart1(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	result := combination.Combinations(2, entries, func(l []int) bool {
		return arithmetic.SumTo(2020, l)
	})

	if !reflect.DeepEqual(result, [][]int{{1721, 299}}) {
		t.Error()
	}
}

func TestPart2(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	result := combination.Combinations(3, entries, func(l []int) bool {
		return arithmetic.SumTo(2020, l)
	})

	if !reflect.DeepEqual(result, [][]int{{979, 366, 675}}) {
		t.Error()
	}
}
