package code

import "errors"

// A Program is a sequence of Instruction.
type Program struct {
	Instructions []*Instruction
}

// ParseProgram returns a valid Program parsed from lines.
func ParseProgram(lines []string) (*Program, error) {
	var instructions []*Instruction

	for _, line := range lines {
		itr, err := ParseInstruction(line)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, itr)
	}

	return &Program{Instructions: instructions}, nil
}

// Fix returns the value of the execution accumulator with the fixed program.
func (p Program) Fix() (int, error) {
	for index, instruction := range p.Instructions {
		switch instruction.Operation {
		case OpJump, OpNoOperation:
			newInstructions := make([]*Instruction, len(p.Instructions))
			copy(newInstructions, p.Instructions)

			if instruction.Operation == OpJump {
				newInstructions[index] = &Instruction{Operation: OpNoOperation, Argument: instruction.Argument}
			} else {
				newInstructions[index] = &Instruction{Operation: OpJump, Argument: instruction.Argument}
			}
			newProgram := &Program{Instructions: newInstructions}

			execution := NewExecution(newProgram)
			loopDetected, err := execution.Run()
			if err != nil {
				return -1, err
			}

			if !loopDetected {
				return execution.Accumulator(), nil
			}
		default:
			continue
		}
	}

	return -1, errors.New("no fix found")
}
