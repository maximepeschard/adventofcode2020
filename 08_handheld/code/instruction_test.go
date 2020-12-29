package code

import (
	"reflect"
	"testing"
)

func TestParseOperation(t *testing.T) {
	validCases := []struct {
		text     string
		expected Operation
	}{
		{"acc", OpAccumulator},
		{"jmp", OpJump},
		{"nop", OpNoOperation},
	}

	invalidCases := []string{"", "abc", "acca"}

	for _, c := range validCases {
		parsed, err := ParseOperation(c.text)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		if parsed != c.expected {
			t.Errorf("ParseOperation(%q) = %v, expected %v", c.text, parsed, c.expected)
		}
	}

	for _, c := range invalidCases {
		_, err := ParseOperation(c)
		if err == nil {
			t.Errorf("expected error for %q, found none", c)
		}
	}
}

func TestParseInstruction(t *testing.T) {
	validCases := []struct {
		text     string
		expected Instruction
	}{
		{"nop +0", Instruction{OpNoOperation, 0}},
		{"acc +1", Instruction{OpAccumulator, 1}},
		{"jmp -3", Instruction{OpJump, -3}},
	}

	invalidCases := []string{"", "nop", "acc", "jmp", "jmp foo"}

	for _, c := range validCases {
		parsed, err := ParseInstruction(c.text)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		if !reflect.DeepEqual(*parsed, c.expected) {
			t.Errorf("ParseInstruction(%q) = %v, expected %v", c.text, *parsed, c.expected)
		}
	}

	for _, c := range invalidCases {
		_, err := ParseInstruction(c)
		if err == nil {
			t.Errorf("expected error for %q, found none", c)
		}
	}
}
