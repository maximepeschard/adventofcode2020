package password

import (
	"fmt"
	"strings"
)

// Parse returns a Policy and a password parsed from s.
func Parse(s string) (*Policy, string, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return nil, "", fmt.Errorf("invalid password entry : %s", s)
	}

	policy, err := ParsePolicy(parts[0])
	if err != nil {
		return nil, "", err
	}

	password := strings.TrimSpace(parts[1])
	if len(password) == 0 {
		return nil, "", fmt.Errorf("invalid password length")
	}

	return policy, password, nil
}
