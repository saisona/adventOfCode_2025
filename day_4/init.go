package day4

import (
	"fmt"

	"github.com/saisona/AdventOfCode2025/helper"
)

func Solver(path string) int {
	values, err := helper.GetInput(3, path)
	if err != nil {
		panic(err)
	}

	for t := range values {
		fmt.Println("t=", t)
	}

	return 0
}
