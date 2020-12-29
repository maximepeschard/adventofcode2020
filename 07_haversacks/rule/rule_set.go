package rule

import (
	"github.com/maximepeschard/adventofcode2020/07_haversacks/stack"
)

// A Set stores a set of Rule.
type Set struct {
	contents   map[string]map[string]int
	containers map[string][]string
	colors     map[string]bool
}

// ParseSet returns a new Set parsed from lines.
func ParseSet(lines []string) (*Set, error) {
	rs := &Set{
		make(map[string]map[string]int),
		make(map[string][]string),
		make(map[string]bool),
	}

	for _, line := range lines {
		rule, err := ParseRule(line)
		if err != nil {
			return rs, nil
		}

		rs.Add(rule)
	}

	return rs, nil
}

// Add adds a new Rule to the Set.
func (rs *Set) Add(rule *Rule) {
	_, exists := rs.contents[rule.Color]
	if !exists {
		rs.contents[rule.Color] = make(map[string]int)
	}
	rs.colors[rule.Color] = true

	for color, qty := range rule.Quantities {
		rs.contents[rule.Color][color] = qty

		_, exists := rs.containers[color]
		if !exists {
			rs.containers[color] = []string{rule.Color}
		} else {
			rs.containers[color] = append(rs.containers[color], rule.Color)
		}

		rs.colors[color] = true
	}
}

// Roots returns the colors the cannot be contained by any other color.
func (rs Set) Roots() []string {
	var roots []string
	for c := range rs.colors {
		if _, hasContainers := rs.containers[c]; !hasContainers {
			roots = append(roots, c)
		}
	}

	return roots
}

// CountContainers returns the number of colors that can contain at least one bag with color.
func (rs Set) CountContainers(color string) (int, error) {
	visited := make(map[string]bool)
	containers := make(map[string]bool)

	type stackFrame struct {
		color string
		path  []string
	}

	for _, root := range rs.Roots() {
		colorStack := stack.New()
		colorStack.Push(stackFrame{root, []string{root}})

		for colorStack.Size() > 0 {
			popped, err := colorStack.Pop()
			if err != nil {
				return -1, err
			}
			frame := popped.(stackFrame)

			_, isContainer := containers[frame.color]
			if frame.color == color {
				for _, c := range frame.path[:len(frame.path)-1] {
					containers[c] = true
				}
			} else if isContainer {
				for _, c := range frame.path {
					containers[c] = true
				}
			}

			if _, alreadyVisited := visited[frame.color]; alreadyVisited {
				continue
			}

			visited[frame.color] = true
			for childColor := range rs.contents[frame.color] {
				newPath := make([]string, len(frame.path)+1)
				copy(newPath, frame.path)
				newPath[len(frame.path)] = childColor
				colorStack.Push(stackFrame{childColor, newPath})
			}
		}
	}

	return len(containers), nil
}

// CountContents returns the number of bags contained by one bag with color.
func (rs Set) CountContents(color string) (int, error) {
	total := 0

	type stackFrame struct {
		color string
		qty   int
	}

	colorStack := stack.New()
	colorStack.Push(stackFrame{color, 1})

	for colorStack.Size() > 0 {
		popped, err := colorStack.Pop()
		if err != nil {
			return -1, err
		}
		frame := popped.(stackFrame)

		for childColor, childQty := range rs.contents[frame.color] {
			total += frame.qty * childQty
			colorStack.Push(stackFrame{childColor, frame.qty * childQty})
		}
	}

	return total, nil
}
