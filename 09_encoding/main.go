package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/maximepeschard/adventofcode2020/09_encoding/encryption"
	"github.com/maximepeschard/adventofcode2020/util"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(flag.Arg(0))
	util.Panic(err)
	defer util.CloseFile(f)

	numbers, err := util.ReadInts(f)
	util.Panic(err)
	preambleLength := 25

	index, err := encryption.FirstInvalidIndex(numbers, preambleLength)
	util.Panic(err)
	firstInvalidNumber := numbers[index]
	fmt.Println(firstInvalidNumber)

	startIndex, endIndex, err := encryption.FindContiguousSetWithSum(numbers, firstInvalidNumber)
	util.Panic(err)
	rangeNumbers := make([]int, endIndex-startIndex+1)
	copy(rangeNumbers, numbers[startIndex:endIndex+1])
	sort.Ints(rangeNumbers)
	encryptionWeakness := rangeNumbers[0] + rangeNumbers[len(rangeNumbers)-1]
	fmt.Println(encryptionWeakness)
}
