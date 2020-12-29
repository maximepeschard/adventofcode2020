package main

import (
	"testing"

	"github.com/maximepeschard/adventofcode2020/08_handheld/code"
)

var testLines = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}

func TestPart1(t *testing.T) {
	program, _ := code.ParseProgram(testLines)
	execution := code.NewExecution(program)
	loopDetected, err := execution.Run()

	if err != nil {
		t.Error(err)
	}

	if !loopDetected {
		t.Errorf("expected loop detection")
	}

	if execution.Accumulator() != 5 {
		t.Errorf("expected accumulator = %d, found %d", 5, execution.Accumulator())
	}
}

func TestPart2(t *testing.T) {
	program, _ := code.ParseProgram(testLines)
	accumulator, err := program.Fix()

	if err != nil {
		t.Error(err)
	}

	if accumulator != 8 {
		t.Errorf("expected accumulator = %d, found %d", 8, accumulator)
	}
}
