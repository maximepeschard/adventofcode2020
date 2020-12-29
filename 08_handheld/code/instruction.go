package code

import (
	"fmt"
	"strconv"
	"strings"
)

// An Operation is an action type.
type Operation string

// Valid Operation values.
const (
	OpAccumulator Operation = "acc"
	OpJump        Operation = "jmp"
	OpNoOperation Operation = "nop"
)

// ParseOperation returns a valid Operation parsed from s.
func ParseOperation(s string) (Operation, error) {
	switch s {
	case string(OpAccumulator):
		return OpAccumulator, nil
	case string(OpJump):
		return OpJump, nil
	case string(OpNoOperation):
		return OpNoOperation, nil
	default:
		return "", fmt.Errorf("invalid operation: %q", s)
	}
}

// An Instruction consists of an Operation and its argument.
type Instruction struct {
	Operation Operation
	Argument  int
}

// ParseInstruction returns a valid instruction parsed from s.
func ParseInstruction(s string) (*Instruction, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid instruction: %q", s)
	}

	op, err := ParseOperation(parts[0])
	if err != nil {
		return nil, err
	}

	arg, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &Instruction{Operation: op, Argument: arg}, nil
}
