package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maximepeschard/adventofcode2020/04_passport/passport"
	"github.com/maximepeschard/adventofcode2020/util"
)

var validatorsPart1 = []passport.Validator{
	passport.PresentValidator("byr"),
	passport.PresentValidator("iyr"),
	passport.PresentValidator("eyr"),
	passport.PresentValidator("hgt"),
	passport.PresentValidator("hcl"),
	passport.PresentValidator("ecl"),
	passport.PresentValidator("pid"),
}
var validatorsPart2 = []passport.Validator{
	passport.IntValidator("byr", 1920, 2002),
	passport.IntValidator("iyr", 2010, 2020),
	passport.IntValidator("eyr", 2020, 2030),
	passport.Any(passport.SizeValidator("hgt", "cm", 150, 193), passport.SizeValidator("hgt", "in", 59, 76)),
	passport.PatternValidator("hcl", `^#[0-9a-f]{6}$`),
	passport.PatternValidator("ecl", `^(amb|blu|brn|gry|grn|hzl|oth)$`),
	passport.PatternValidator("pid", `^\d{9}$`),
}

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

	validPart1, validPart2 := 0, 0
	err = passport.ProcessBatchFile(lines, func(p *passport.Passport) error {
		if p.IsValid(validatorsPart1) {
			validPart1++
		}
		if p.IsValid(validatorsPart2) {
			validPart2++
		}
		return nil
	})
	util.Panic(err)
	fmt.Println(validPart1)
	fmt.Println(validPart2)
}
