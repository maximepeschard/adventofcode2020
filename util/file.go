package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// CloseFile closes a file and panics if an error occurs.
func CloseFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Panicln(err)
	}
}

// ReadInts parses a file with one integer per line.
func ReadInts(f *os.File) ([]int, error) {
	scanner := bufio.NewScanner(f)
	var ints []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}

	return ints, scanner.Err()
}
