package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maximepeschard/adventofcode2020/06_customs/customs"
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

	sumAnsweredByAnyone, sumAnsweredByEveryone := 0, 0

	err = customs.ProcessAnswerList(lines, func(ag *customs.AnswerGroup) error {
		sumAnsweredByAnyone += len(ag.AnswersPerQuestion)

		for _, n := range ag.AnswersPerQuestion {
			if n == ag.People {
				sumAnsweredByEveryone++
			}
		}

		return nil
	})

	fmt.Println(sumAnsweredByAnyone)
	fmt.Println(sumAnsweredByEveryone)
}
