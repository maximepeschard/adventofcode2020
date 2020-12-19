package passport

import "testing"

func TestPresentValidator(t *testing.T) {
	validator := PresentValidator("foo")

	cases := []struct {
		passport Passport
		expected bool
	}{
		{Passport{fields: map[string]string{}}, false},
		{Passport{fields: map[string]string{"foo": "bar"}}, true},
		{Passport{fields: map[string]string{"foo": "bar", "baz": "123"}}, true},
		{Passport{fields: map[string]string{"baz": "123"}}, false},
	}

	for _, c := range cases {
		if validator(c.passport) != c.expected {
			t.Error()
		}
	}
}

func TestIntValidator(t *testing.T) {
	validator := IntValidator("foo", 2, 4)

	cases := []struct {
		passport Passport
		expected bool
	}{
		{Passport{fields: map[string]string{}}, false},
		{Passport{fields: map[string]string{"foo": "bar"}}, false},
		{Passport{fields: map[string]string{"foo": "2"}}, true},
		{Passport{fields: map[string]string{"foo": "3", "baz": "123"}}, true},
		{Passport{fields: map[string]string{"foo": "5"}}, false},
		{Passport{fields: map[string]string{"baz": "123"}}, false},
	}

	for _, c := range cases {
		if validator(c.passport) != c.expected {
			t.Error()
		}
	}
}

func TestSizeValidator(t *testing.T) {
	validator := SizeValidator("foo", "px", 2, 4)

	cases := []struct {
		passport Passport
		expected bool
	}{
		{Passport{fields: map[string]string{}}, false},
		{Passport{fields: map[string]string{"foo": "bar"}}, false},
		{Passport{fields: map[string]string{"foo": "2"}}, false},
		{Passport{fields: map[string]string{"foo": "2px"}}, true},
		{Passport{fields: map[string]string{"foo": "3px", "baz": "123"}}, true},
		{Passport{fields: map[string]string{"foo": "2cm"}}, false},
		{Passport{fields: map[string]string{"baz": "123"}}, false},
	}

	for _, c := range cases {
		if validator(c.passport) != c.expected {
			t.Error()
		}
	}
}

func TestPatternValidator(t *testing.T) {
	validator := PatternValidator("foo", `^[0-4a-c]{3}$`)

	cases := []struct {
		passport Passport
		expected bool
	}{
		{Passport{fields: map[string]string{}}, false},
		{Passport{fields: map[string]string{"foo": "bar"}}, false},
		{Passport{fields: map[string]string{"foo": "abc"}}, true},
		{Passport{fields: map[string]string{"foo": "abcd"}}, false},
		{Passport{fields: map[string]string{"foo": "abb", "baz": "123"}}, true},
		{Passport{fields: map[string]string{"baz": "123"}}, false},
	}

	for _, c := range cases {
		if validator(c.passport) != c.expected {
			t.Error()
		}
	}
}
