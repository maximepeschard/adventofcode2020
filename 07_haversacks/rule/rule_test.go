package rule

import (
	"reflect"
	"testing"
)

func TestParseRule(t *testing.T) {
	validCases := []struct {
		s        string
		expected Rule
	}{
		{"light red bags contain 1 bright white bag, 2 muted yellow bags.", Rule{"light red", map[string]int{"bright white": 1, "muted yellow": 2}}},
		{"bright white bags contain 1 shiny gold bag.", Rule{"bright white", map[string]int{"shiny gold": 1}}},
		{"faded blue bags contain no other bags.", Rule{"faded blue", map[string]int{}}},
	}

	invalidCases := []string{
		"",
		"light red bags",
		"light red bags contain 1 foo.",
	}

	for _, c := range validCases {
		parsed, err := ParseRule(c.s)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		if !reflect.DeepEqual(*parsed, c.expected) {
			t.Errorf("ParseRule(%q) = %v, expected %v", c.s, *parsed, c.expected)
		}
	}

	for _, c := range invalidCases {
		_, err := ParseRule(c)
		if err == nil {
			t.Errorf("expected error for %q, found none", c)
		}
	}
}
