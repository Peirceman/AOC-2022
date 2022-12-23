package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/Peirceman/AOC-2022/days"
)

func main() {
	var day, part uint
	var test bool

	flag.UintVar(&day, "day", 0, "The day to run")
	flag.UintVar(&part, "part", 0, "The part to run")
	flag.BoolVar(&test, "test", false, "if true, run the test case, if false, run the real input")

	flag.Parse()

	if day < 1 {
		fmt.Fprintln(os.Stderr, "please select a day")
		os.Exit(1)
	}

	if part < 1 {
		fmt.Fprintln(os.Stderr, "please select a part")
		os.Exit(1)
	}

	filePath := "./inputs/day" + strconv.FormatUint(uint64(day), 10)

	if test {
		filePath += "_test.txt"
	} else {
		filePath += ".txt"
	}

	days.Run(day, part, filePath)
}
