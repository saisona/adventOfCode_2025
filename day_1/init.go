package day1

import (
	"strconv"
	"strings"

	"github.com/saisona/AdventOfCode2025/helper"
)

func turn(p byte, rangeC int, position *int, finalPassword *int) {
	switch p {
	case 'L':
		isgreatherThan := rangeC > *position
		if isgreatherThan {
			tmpPostToZero := rangeC - *position
			tmpPostToZero = tmpPostToZero - (tmpPostToZero/100)*100

			if tmpPostToZero == 0 {
				*position = 0
			} else {
				*position = 100 - tmpPostToZero
			}
		} else {
			*position -= rangeC
		}

	case 'R':
		*position += rangeC
		if *position >= 100 {
			*position = *position % 100
		}
	}

	if *position == 0 {
		*finalPassword = *finalPassword + 1
	}
}

func Solver(path string) int {
	iter, err := helper.GetInput(1, path)
	if err != nil {
		panic(err)
	}

	finalPassword := 0
	startingPoint := 50

	for newString := range iter {
		position := newString[0]
		rangeC, _ := strconv.Atoi(strings.TrimRight(newString[1:], "\n"))
		turn(position, rangeC, &startingPoint, &finalPassword)
	}
	return finalPassword
}
