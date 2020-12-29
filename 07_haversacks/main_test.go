package main

import (
	"testing"

	"github.com/maximepeschard/adventofcode2020/07_haversacks/rule"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}
	ruleSet, _ := rule.ParseSet(lines)
	containers, err := ruleSet.CountContainers("shiny gold")
	if err != nil || containers != 4 {
		t.Error()
	}
}

func TestPart2(t *testing.T) {
	lines := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}
	ruleSet, _ := rule.ParseSet(lines)
	contents, err := ruleSet.CountContents("shiny gold")
	if err != nil || contents != 32 {
		t.Error()
	}

	lines = []string{
		"shiny gold bags contain 2 dark red bags.",
		"dark red bags contain 2 dark orange bags.",
		"dark orange bags contain 2 dark yellow bags.",
		"dark yellow bags contain 2 dark green bags.",
		"dark green bags contain 2 dark blue bags.",
		"dark blue bags contain 2 dark violet bags.",
		"dark violet bags contain no other bags.",
	}
	ruleSet, _ = rule.ParseSet(lines)
	contents, err = ruleSet.CountContents("shiny gold")
	if err != nil || contents != 126 {
		t.Error()
	}
}
