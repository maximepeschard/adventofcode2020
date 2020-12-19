package boarding

import "testing"

func TestParseRow(t *testing.T) {
	validCases := []struct {
		chars    string
		expected int
	}{
		{"FBFBBFF", 44},
		{"BFFFBBF", 70},
		{"FFFBBBF", 14},
		{"BBFFBBF", 102},
	}

	for _, c := range validCases {
		row, err := ParseRow(c.chars)
		if err != nil {
			t.Errorf("unexpected error for ParseRow(%s) : %s", c.chars, err)
		}
		if row != c.expected {
			t.Errorf("ParseRow(%s) = %d, expected %d", c.chars, row, c.expected)
		}
	}

	invalidCases := []string{
		"",
		"FBF",
		"FBFBBFL",
		"FBFBBFFFFF",
	}

	for _, c := range invalidCases {
		_, err := ParseRow(c)
		if err == nil {
			t.Errorf("expected error for ParseRow(%s), got none", c)
		}
	}
}

func TestParseColumn(t *testing.T) {
	validCases := []struct {
		chars    string
		expected int
	}{
		{"RLR", 5},
		{"RRR", 7},
		{"RLL", 4},
	}

	for _, c := range validCases {
		column, err := ParseColumn(c.chars)
		if err != nil {
			t.Errorf("unexpected error for ParseColumn(%s) : %s", c.chars, err)
		}
		if column != c.expected {
			t.Errorf("ParseColumn(%s) = %d, expected %d", c.chars, column, c.expected)
		}
	}

	invalidCases := []string{
		"",
		"RL",
		"RLF",
		"RLRR",
	}

	for _, c := range invalidCases {
		_, err := ParseColumn(c)
		if err == nil {
			t.Errorf("expected error for ParseColumn(%s), got none", c)
		}
	}
}

func TestSeatID(t *testing.T) {
	cases := []struct {
		pass     string
		expected int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, c := range cases {
		seatID, err := SeatID(c.pass)
		if err != nil {
			t.Errorf("unexpected error for SeatID(%s) : %s", c.pass, err)
		}
		if seatID != c.expected {
			t.Errorf("SeatID(%s) = %d, expected %d", c.pass, seatID, c.expected)
		}
	}
}
