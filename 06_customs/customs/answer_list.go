package customs

// ProcessAnswerList processes a list of customs answers by applying f to each parsed AnswerGroup.
func ProcessAnswerList(lines []string, f func(*AnswerGroup) error) error {
	var buffer []string

	for _, line := range lines {
		if line == "" {
			ag, err := ParseAnswerGroup(buffer)
			if err != nil {
				return err
			}
			buffer = nil
			if err := f(ag); err != nil {
				return err
			}
			continue
		}

		buffer = append(buffer, line)
	}

	if len(buffer) > 0 {
		ag, err := ParseAnswerGroup(buffer)
		if err != nil {
			return err
		}
		buffer = nil
		if err := f(ag); err != nil {
			return err
		}
	}

	return nil
}
