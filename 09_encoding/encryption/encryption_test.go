package encryption

import "testing"

var testNumbers = []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
var testPreambleLength = 5

func TestFirstInvalidIndex(t *testing.T) {
	index, err := FirstInvalidIndex(testNumbers, testPreambleLength)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if testNumbers[index] != 127 {
		t.Error("wrong invalid index")
	}
}

func TestFindContiguousSetWithSum(t *testing.T) {
	startIndex, endIndex, err := FindContiguousSetWithSum(testNumbers, 127)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if startIndex != 2 || endIndex != 5 {
		t.Error("wrong contiguous set")
	}
}
