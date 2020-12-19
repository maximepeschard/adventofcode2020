package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maximepeschard/adventofcode2020/05_boarding/boarding"
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

	passes, err := util.ReadLines(f)
	util.Panic(err)

	highestSeatID := -1
	presentSeats := make(map[int]bool)
	for _, pass := range passes {
		seatID, err := boarding.SeatID(pass)
		util.Panic(err)

		presentSeats[seatID] = true

		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Println(highestSeatID)

	missingSeatID := -1
	for seatID := highestSeatID - 1; seatID > 0; seatID-- {
		if !presentSeats[seatID] && presentSeats[seatID-1] && presentSeats[seatID+1] {
			missingSeatID = seatID
			break
		}
	}
	fmt.Println(missingSeatID)
}
