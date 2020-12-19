package passport

import (
	"fmt"
	"strings"
)

// A Passport represents a set of fields.
type Passport struct {
	fields map[string]string
}

// ParsePassport returns a Passport parsed from one or more lines of text.
func ParsePassport(lines []string) (*Passport, error) {
	fields := make(map[string]string)
	for _, line := range lines {
		for _, f := range strings.Fields(line) {
			parts := strings.SplitN(f, ":", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid field %s", f)
			}

			fields[parts[0]] = parts[1]
		}
	}

	return &Passport{fields}, nil
}

// IsValid returns whether a passport is valid.
func (p Passport) IsValid(validators []Validator) bool {
	for _, v := range validators {
		if !v(p) {
			return false
		}
	}

	return true
}
