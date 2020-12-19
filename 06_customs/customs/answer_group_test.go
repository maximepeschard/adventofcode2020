package customs

import (
	"reflect"
	"testing"
)

func TestParseAnswerGroup(t *testing.T) {
	cases := []struct {
		lines    []string
		expected AnswerGroup
	}{
		{[]string{}, AnswerGroup{0, map[rune]int{}}},
		{[]string{"ab"}, AnswerGroup{1, map[rune]int{'a': 1, 'b': 1}}},
		{[]string{"ab", "bc"}, AnswerGroup{2, map[rune]int{'a': 1, 'b': 2, 'c': 1}}},
		{[]string{"a", "b"}, AnswerGroup{2, map[rune]int{'a': 1, 'b': 1}}},
	}

	for _, c := range cases {
		parsed, err := ParseAnswerGroup(c.lines)
		if err != nil {
			t.Errorf("unexpected error for ParseAnswerGroup(%v) : %s", c.lines, err)
		}
		if !reflect.DeepEqual(c.expected, *parsed) {
			t.Errorf("ParseAnswerGroup(%v) = %v, expected %v", c.lines, *parsed, c.expected)
		}
	}
}
