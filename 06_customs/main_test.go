package main

import (
	"testing"

	"github.com/maximepeschard/adventofcode2020/06_customs/customs"
)

var testLines = []string{
	"abc",
	"",
	"a",
	"b",
	"c",
	"",
	"ab",
	"ac",
	"",
	"a",
	"a",
	"a",
	"a",
	"",
	"b",
}

func TestPart1(t *testing.T) {
	sum := 0
	err := customs.ProcessAnswerList(testLines, func(ag *customs.AnswerGroup) error {
		sum += len(ag.AnswersPerQuestion)
		return nil
	})

	if err != nil || sum != 11 {
		t.Error()
	}
}

func TestPart2(t *testing.T) {
	sum := 0
	err := customs.ProcessAnswerList(testLines, func(ag *customs.AnswerGroup) error {
		for _, n := range ag.AnswersPerQuestion {
			if n == ag.People {
				sum++
			}
		}
		return nil
	})

	if err != nil || sum != 6 {
		t.Error()
	}
}
