package passport

import (
	"reflect"
	"testing"
)

func TestParsePassport(t *testing.T) {
	validCases := []struct {
		lines    []string
		expected Passport
	}{
		{[]string{}, Passport{map[string]string{}}},
		{[]string{"abc:foo def:123"}, Passport{map[string]string{"abc": "foo", "def": "123"}}},
		{[]string{"abc:foo", "def:123"}, Passport{map[string]string{"abc": "foo", "def": "123"}}},
	}

	for _, c := range validCases {
		parsed, err := ParsePassport(c.lines)
		if err != nil {
			t.Errorf("unexpected error : %s", err)
		}

		if !reflect.DeepEqual(c.expected, *parsed) {
			t.Errorf("ParsePassport(%v) = %v, expected %v", c.lines, *parsed, c.expected)
		}
	}

	invalidCases := [][]string{
		{"abc"},
		{"abc def"},
		{"abc:foo", "def"},
	}

	for _, c := range invalidCases {
		_, err := ParsePassport(c)
		if err == nil {
			t.Errorf("expected error for %v, got none", c)
		}
	}
}
