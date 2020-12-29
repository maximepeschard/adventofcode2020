package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maximepeschard/adventofcode2020/08_handheld/code"
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

	lines, err := util.ReadLines(f)
	util.Panic(err)

	program, err := code.ParseProgram(lines)
	util.Panic(err)

	execution := code.NewExecution(program)
	_, err = execution.Run()
	util.Panic(err)
	fmt.Println(execution.Accumulator())

	accumulator, err := program.Fix()
	util.Panic(err)
	fmt.Println(accumulator)
}
