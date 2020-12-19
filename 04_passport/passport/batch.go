package passport

// ProcessBatchFile processes a batch file of passport data by applying f to each parsed Passport.
func ProcessBatchFile(lines []string, f func(p *Passport) error) error {
	var buffer []string

	for _, line := range lines {
		if line == "" {
			p, err := ParsePassport(buffer)
			if err != nil {
				return err
			}
			buffer = nil
			err = f(p)
			if err != nil {
				return err
			}
			continue
		}

		buffer = append(buffer, line)
	}

	if len(buffer) > 0 {
		p, err := ParsePassport(buffer)
		if err != nil {
			return err
		}
		err = f(p)
		if err != nil {
			return err
		}
	}

	return nil
}
