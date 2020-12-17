package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/maximepeschard/adventofcode2020/03_toboggan/toboggan"
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

	grid, err := util.ReadLines(f)
	util.Panic(err)

	travelMap, err := toboggan.ParseMap(grid)
	util.Panic(err)

	treesPart1 := travelMap.TreesOnSlope(3, 1)
	fmt.Println(treesPart1)

	treeCounts := make(chan int)
	countForSlope := func(right int, down int, wg *sync.WaitGroup) {
		defer wg.Done()
		treeCounts <- travelMap.TreesOnSlope(right, down)
	}
	var wg sync.WaitGroup
	wg.Add(5)
	go countForSlope(1, 1, &wg)
	go countForSlope(3, 1, &wg)
	go countForSlope(5, 1, &wg)
	go countForSlope(7, 1, &wg)
	go countForSlope(1, 2, &wg)
	go func() {
		wg.Wait()
		close(treeCounts)
	}()

	treesPart2 := 1
	for count := range treeCounts {
		treesPart2 *= count
	}
	fmt.Println(treesPart2)
}
