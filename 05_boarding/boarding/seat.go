package boarding

import (
	"fmt"
	"math"
)

// Boarding pass specification
const (
	NumCharsRow    = 7
	NumCharsColumn = 3
)

// ParseRow returns a row number parsed from characters.
func ParseRow(chars string) (int, error) {
	if len(chars) != NumCharsRow {
		return -1, fmt.Errorf("row must be specified by %d characters", NumCharsRow)
	}

	row := 0
	for i := 0; i < NumCharsRow; i++ {
		n := -1
		switch chars[NumCharsRow-i-1] {
		case 'F':
			n = 0
		case 'B':
			n = 1
		default:
			return -1, fmt.Errorf("invalid character in row specification : %v", chars[NumCharsRow-i-1])
		}

		row += n * int(math.Pow(float64(2), float64(i)))
	}

	return row, nil
}

// ParseColumn returns a column number parsed from characters.
func ParseColumn(chars string) (int, error) {
	if len(chars) != NumCharsColumn {
		return -1, fmt.Errorf("column must be specified by %d characters", NumCharsColumn)
	}

	column := 0
	for i := 0; i < NumCharsColumn; i++ {
		n := -1
		switch chars[NumCharsColumn-i-1] {
		case 'L':
			n = 0
		case 'R':
			n = 1
		default:
			return -1, fmt.Errorf("invalid character in column specification : %v", chars[NumCharsColumn-i-1])
		}

		column += n * int(math.Pow(float64(2), float64(i)))
	}

	return column, nil
}

// SeatID returns a seat ID from a boarding pass.
func SeatID(pass string) (int, error) {
	if len(pass) != NumCharsRow+NumCharsColumn {
		return -1, fmt.Errorf("a boading pass must be specified by %d characters", NumCharsRow+NumCharsColumn)
	}

	row, err := ParseRow(pass[:NumCharsRow])
	if err != nil {
		return -1, err
	}

	column, err := ParseColumn(pass[NumCharsRow:])
	if err != nil {
		return -1, err
	}

	return row*8 + column, nil
}
