package customs

// An AnswerGroup summarizes the answers given by a group to the customs.
type AnswerGroup struct {
	People             int          // number of people in the group
	AnswersPerQuestion map[rune]int // number of answers per question
}

// ParseAnswerGroup returns an AnswerGroup parsed from lines of text.
func ParseAnswerGroup(lines []string) (*AnswerGroup, error) {
	people := 0
	answers := make(map[rune]int)
	for _, line := range lines {
		people++
		for _, letter := range line {
			answers[letter]++
		}
	}

	return &AnswerGroup{people, answers}, nil
}
