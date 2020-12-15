package password

import (
	"fmt"
	"regexp"
	"strconv"
)

var policyPattern = regexp.MustCompile(`(?P<lowest>\d+)-(?P<highest>\d+)\s+(?P<letter>[[:alpha:]])`)

// Policy describes a password policy.
type Policy struct {
	letter  rune
	lowest  int
	highest int
}

// ParsePolicy returns a Policy parsed from s.
func ParsePolicy(s string) (*Policy, error) {
	matches := policyPattern.FindStringSubmatch(s)
	if matches == nil {
		return nil, fmt.Errorf("invalid policy: %s", s)
	}

	letter := rune(matches[policyPattern.SubexpIndex("letter")][0])
	lowest, _ := strconv.Atoi(matches[policyPattern.SubexpIndex("lowest")])
	highest, _ := strconv.Atoi(matches[policyPattern.SubexpIndex("highest")])

	if lowest > highest {
		return nil, fmt.Errorf("invalid policy: %s", s)
	}

	return &Policy{letter, lowest, highest}, nil
}

// CheckCountRule indicates if s verifies Policy p with a count rule.
func (p Policy) CheckCountRule(s string) bool {
	if p.lowest > len(s) {
		return false
	}

	count := 0
	for _, l := range s {
		if l == p.letter {
			count++
		}
	}

	return p.lowest <= count && count <= p.highest
}

// CheckPositionRule indicates if s verifies Policy p with a position rule.
func (p Policy) CheckPositionRule(s string) bool {
	if p.lowest > len(s) || p.highest > len(s) {
		return false
	}

	lowestMatch := rune(s[p.lowest-1]) == p.letter
	highestMatch := rune(s[p.highest-1]) == p.letter

	return (lowestMatch || highestMatch) && !(lowestMatch && highestMatch)
}
