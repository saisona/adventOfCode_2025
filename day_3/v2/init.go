package day_3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/saisona/AdventOfCode2025/helper"
)

// bestNumberInStringV2 is the same function as for day_3.bestNumberInString but for part II
func bestNumberInString(line string) int {
	line = line[:len(line)-1]

	dPtr := 0
	uPtr := 1

	ptr2 := 1
	strBuilder := &strings.Builder{}

	for ptr2 < len(line) {
		// Check ^10

		if (ptr2 < len(line)-1) && (line[ptr2] > line[dPtr]) {
			dPtr = ptr2
			uPtr = dPtr + 1
			ptr2 = uPtr
			fmt.Printf("d=%d u=%d\n", dPtr, uPtr)
			continue
		}

		// Check units
		if ptr2 < len(line)-12 && line[ptr2] > line[uPtr] {
			uPtr = ptr2
			fmt.Printf("u=%d\n", uPtr)
		}

		ptr2++
		// fmt.Printf("ptr2=%d\n", ptr2)
	}

	strBuilder.WriteRune(rune(line[dPtr]))
	strBuilder.WriteRune(rune(line[uPtr]))

	intVal, _ := strconv.Atoi(strBuilder.String())
	strBuilder.Reset()
	return intVal
}

// Solver part II of the day_3 of AdventOfCode2025
func Solver(path string) int {
	values, err := helper.GetInput(3, path)
	if err != nil {
		panic(err)
	}

	sum := 0
	for rank := range values {
		sum += bestNumberInString(rank)
	}
	return sum
}
