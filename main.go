package main

import (
	"flag"
	"fmt"

	_ "github.com/saisona/AdventOfCode2025/day_1"
	_ "github.com/saisona/AdventOfCode2025/day_2"
	day3 "github.com/saisona/AdventOfCode2025/day_3"
)

var inputfile = flag.String("input", "", "file to read input")

func main() {
	flag.Parse()

	if inputfile == nil || *inputfile == "" {
		panic("input file should be given")
	}

	// _ = day1.Solver(os.Args[1])
	// res := day2.Solver(*inputfile)
	res := day3.Solver(*inputfile)
	fmt.Println("Res=", res)
}
