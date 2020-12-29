package code

// An Execution holds information about a Program execution.
type Execution struct {
	program     *Program
	accumulator int
}

// NewExecution returns a new execution of p.
func NewExecution(p *Program) *Execution {
	return &Execution{program: p, accumulator: 0}
}

// Accumulator returns the current value of the execution accumulator.
func (e Execution) Accumulator() int {
	return e.accumulator
}

// Run runs the program until it terminates, or until a loop is detected.
func (e *Execution) Run() (bool, error) {
	iptr := 0
	loopDetected := false
	executed := make(map[int]bool)

	for !loopDetected && iptr >= 0 && iptr < len(e.program.Instructions) {
		instruction := e.program.Instructions[iptr]
		executed[iptr] = true

		switch instruction.Operation {
		case OpAccumulator:
			e.accumulator += instruction.Argument
			iptr++
		case OpJump:
			iptr += instruction.Argument
		case OpNoOperation:
			iptr++
		}

		_, loopDetected = executed[iptr]
	}

	return loopDetected, nil
}
