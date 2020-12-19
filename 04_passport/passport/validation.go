package passport

import (
	"regexp"
	"strconv"
)

// A Validator is function to check a passport.
type Validator = func(Passport) bool

// PresentValidator returns a Validator to check the presence of a field.
func PresentValidator(field string) Validator {
	return func(p Passport) bool {
		_, present := p.fields[field]
		return present
	}
}

// IntValidator returns a Validator to check an integer field with constraints.
func IntValidator(field string, min int, max int) Validator {
	return func(p Passport) bool {
		value, present := p.fields[field]
		if !present {
			return false
		}

		n, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		return min <= n && n <= max
	}
}

// SizeValidator returns a Validator to check a size field.
func SizeValidator(field string, unit string, min int, max int) Validator {
	return func(p Passport) bool {
		value, present := p.fields[field]
		if !present {
			return false
		}

		sizePattern := regexp.MustCompile(`^(?P<size>\d+)` + unit + `$`)
		matches := sizePattern.FindStringSubmatch(value)
		if matches == nil {
			return false
		}

		size, _ := strconv.Atoi(matches[sizePattern.SubexpIndex("size")])

		return min <= size && size <= max
	}
}

// PatternValidator returns a Validator to check that a field matches a pattern.
func PatternValidator(field string, pattern string) Validator {
	return func(p Passport) bool {
		value, present := p.fields[field]
		if !present {
			return false
		}

		re := regexp.MustCompile(pattern)

		return re.MatchString(value)
	}
}

// Any returns a Validator which is the union of multiple validators.
func Any(validators ...Validator) Validator {
	return func(p Passport) bool {
		for _, v := range validators {
			if v(p) {
				return true
			}
		}

		return false
	}
}
