package encryption

import (
	"errors"

	"github.com/maximepeschard/adventofcode2020/01_report/combination"
)

// FirstInvalidIndex returns the index of the first invalid number.
func FirstInvalidIndex(numbers []int, preambleLength int) (int, error) {
	for i := preambleLength; i < len(numbers); i++ {
		valid := false
		for _, comb := range combination.Combinations(2, numbers[i-preambleLength:i], func([]int) bool { return true }) {
			valid = valid || comb[0]+comb[1] == numbers[i]
		}
		if !valid {
			return i, nil
		}
	}

	return -1, errors.New("no invalid number found")
}

// FindContiguousSetWithSum returns the start and end indexes of a contiguous set of
// items in numbers that sum up to target.
func FindContiguousSetWithSum(numbers []int, target int) (int, int, error) {
	startIndex, endIndex := 0, 1
	sum := numbers[startIndex] + numbers[endIndex]
	for sum != target {
		if sum < target {
			if endIndex == len(numbers)-1 {
				return -1, -1, errors.New("no set found")
			}
			endIndex++
			sum += numbers[endIndex]
		} else {
			sum -= numbers[startIndex]
			startIndex++
		}
	}

	return startIndex, endIndex, nil
}
