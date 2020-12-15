package password

import (
	"reflect"
	"testing"
)

func TestParsePolicy(t *testing.T) {
	validCases := []struct {
		s        string
		expected Policy
	}{
		{"1-3 a", Policy{'a', 1, 3}},
		{"1-3 b", Policy{'b', 1, 3}},
		{"2-9 c", Policy{'c', 2, 9}},
	}

	invalidCases := []string{
		"",
		"123",
		"abc",
		"1-3",
		"1-3 4",
		"1-a b",
	}

	for _, c := range validCases {
		policy, err := ParsePolicy(c.s)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		if !reflect.DeepEqual(*policy, c.expected) {
			t.Errorf("ParsePolicy(%s) = %v, expected %v", c.s, policy, c.expected)
		}
	}

	for _, c := range invalidCases {
		_, err := ParsePolicy(c)
		if err == nil {
			t.Errorf("expected error for %s, got none", c)
		}
	}
}

func TestCheckCountRule(t *testing.T) {
	cases := []struct {
		p        Policy
		s        string
		expected bool
	}{
		{Policy{'a', 1, 3}, "abcde", true},
		{Policy{'a', 1, 3}, "aaade", true},
		{Policy{'a', 1, 3}, "aaaaa", false},
		{Policy{'a', 1, 3}, "zbcde", false},
		{Policy{'a', 8, 10}, "abcde", false},
	}

	for _, c := range cases {
		result := c.p.CheckCountRule(c.s)
		if result != c.expected {
			t.Errorf("%v.CheckCountRule(%s) = %t, expected %t", c.p, c.s, result, c.expected)
		}
	}
}

func TestCheckPositionRule(t *testing.T) {
	cases := []struct {
		p        Policy
		s        string
		expected bool
	}{
		{Policy{'a', 1, 3}, "abcde", true},
		{Policy{'a', 1, 3}, "aaade", false},
		{Policy{'a', 1, 3}, "aaaaa", false},
		{Policy{'a', 1, 3}, "zbcde", false},
		{Policy{'a', 8, 10}, "abcde", false},
	}

	for _, c := range cases {
		result := c.p.CheckPositionRule(c.s)
		if result != c.expected {
			t.Errorf("%v.CheckPositionRule(%s) = %t, expected %t", c.p, c.s, result, c.expected)
		}
	}
}
