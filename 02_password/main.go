package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maximepeschard/adventofcode2020/02_password/password"
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

	entries, err := util.ReadLines(f)
	util.Panic(err)

	validPart1, validPart2 := 0, 0
	for _, entry := range entries {
		policy, password, err := password.Parse(entry)
		util.Panic(err)

		if policy.CheckCountRule(password) {
			validPart1++
		}
		if policy.CheckPositionRule(password) {
			validPart2++
		}
	}

	fmt.Println(validPart1)
	fmt.Println(validPart2)
}
