package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maximepeschard/adventofcode2020/07_haversacks/rule"
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

	ruleSet, err := rule.ParseSet(lines)
	util.Panic(err)

	containers, err := ruleSet.CountContainers("shiny gold")
	util.Panic(err)
	fmt.Println(containers)

	contents, err := ruleSet.CountContents("shiny gold")
	util.Panic(err)
	fmt.Println(contents)
}
