package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/maximepeschard/adventofcode2020/01_report/arithmetic"
	"github.com/maximepeschard/adventofcode2020/01_report/combination"
	"github.com/maximepeschard/adventofcode2020/util"
)

var usage = `Usage: report <input file>`

func main() {
	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(flag.Arg(0))
	util.Panic(err)
	defer util.CloseFile(f)

	entries, err := util.ReadInts(f)
	util.Panic(err)

	sumTo2020 := func(l []int) bool { return arithmetic.SumTo(2020, l) }

	part1 := combination.Combinations(2, entries, sumTo2020)
	if len(part1) != 1 {
		util.Panic(errors.New("invalid result for part 1"))
	}
	fmt.Println(part1[0][0] * part1[0][1])

	part2 := combination.Combinations(3, entries, sumTo2020)
	if len(part2) != 1 {
		util.Panic(errors.New("invalid result for part 2"))
	}
	fmt.Println(part2[0][0] * part2[0][1] * part2[0][2])
}
