package stack

import "errors"

// Stack is a generic stack type.
type Stack struct {
	items []interface{}
}

// New returns a new empty Stack.
func New() *Stack {
	return &Stack{}
}

// Push adds an element onto the Stack.
func (s *Stack) Push(e interface{}) {
	s.items = append(s.items, e)
}

// Pop removes and returns the first element on the Stack.
func (s *Stack) Pop() (interface{}, error) {
	size := len(s.items)

	if size == 0 {
		return nil, errors.New("empty stack")
	}

	popped := s.items[size-1]
	s.items = s.items[:size-1]

	return popped, nil
}

// Size returns the number of elements in the Stack.
func (s Stack) Size() int {
	return len(s.items)
}
