package rule

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var colorPattern = `[a-z]+ [a-z]+`
var rulePattern = `^(?P<out>` + colorPattern + `) bags contain (?P<ins>[0-9a-z,\s]+).$`
var emptyPattern = `no other bags`

var colorRegexp = regexp.MustCompile(colorPattern)
var colorQuantityRegexp = regexp.MustCompile(`(?P<qty>\d+) (?P<in>` + colorPattern + `)`)
var ruleRegexp = regexp.MustCompile(rulePattern)

// A Rule specifies the contents of a colored bag.
type Rule struct {
	Color      string
	Quantities map[string]int
}

// ParseRule returns a new Rule parsed from s.
func ParseRule(s string) (*Rule, error) {
	rule := &Rule{}

	matches := ruleRegexp.FindStringSubmatch(s)
	if matches == nil {
		return rule, fmt.Errorf("invalid rule: %s", s)
	}

	out := matches[ruleRegexp.SubexpIndex("out")]
	ins := matches[ruleRegexp.SubexpIndex("ins")]
	rule.Color = out
	rule.Quantities = make(map[string]int)

	if ins == emptyPattern {
		return rule, nil
	}

	for _, item := range strings.Split(ins, ",") {
		matches := colorQuantityRegexp.FindStringSubmatch(item)
		if matches == nil {
			return rule, fmt.Errorf("invalid quantity: %s", item)
		}

		qty, _ := strconv.Atoi(matches[colorQuantityRegexp.SubexpIndex("qty")])
		in := matches[colorQuantityRegexp.SubexpIndex("in")]
		rule.Quantities[in] = qty
	}

	return rule, nil
}
