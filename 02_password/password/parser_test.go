package password

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	validCases := []struct {
		s                string
		expectedPolicy   Policy
		expectedPassword string
	}{
		{"1-3 a: abcde", Policy{'a', 1, 3}, "abcde"},
		{"1-3 a  :   abcde", Policy{'a', 1, 3}, "abcde"},
		{"1-3 b: cdefg", Policy{'b', 1, 3}, "cdefg"},
		{"2-9 c: ccccccccc", Policy{'c', 2, 9}, "ccccccccc"},
	}

	invalidCases := []string{
		"",
		"1-3:",
		"1-3: ",
		"1-3 2: abcde",
	}

	for _, c := range validCases {
		policy, password, err := Parse(c.s)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		if !reflect.DeepEqual(*policy, c.expectedPolicy) ||
			password != c.expectedPassword {
			t.Errorf("Parse(%s) = %v, %s, expected %v, %s", c.s, *policy, password, c.expectedPolicy, c.expectedPassword)
		}
	}

	for _, c := range invalidCases {
		_, _, err := Parse(c)
		if err == nil {
			t.Errorf("expected error for %s, got none", c)
		}
	}
}
