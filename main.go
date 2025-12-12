package main

import (
	"flag"
	"fmt"

	// day1 "github.com/saisona/AdventOfCode2025/day_1"
	day1 "github.com/saisona/AdventOfCode2025/day_1/v2"
)

var inputfile = flag.String("input", "", "file to read input")

func main() {
	flag.Parse()

	if inputfile == nil || *inputfile == "" {
		panic("input file should be given")
	}

	res := day1.Solver(*inputfile)
	// res := day2.Solver(*inputfile)
	// res := day3.Solver(*inputfile)
	// res := day4.Solver(*inputfile)
	fmt.Println("Res=", res)
}
