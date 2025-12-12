// Package day1
package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/saisona/AdventOfCode2025/helper"
)

func addMinusToSign(p byte) rune {
	if p == 'L' {
		return '-'
	} else {
		return ' '
	}
}

func turn(p byte, rangeC int, position *int, finalPassword *int) {
	oldP := *position
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

	fmt.Printf("[%c] turned %c%d times\n", rune(p), addMinusToSign(p), rangeC)

	updateFinalPassword(position, finalPassword, rangeC, oldP)
}

func updateFinalPassword(position *int, finalPassword *int, rangeC int, oldP int) {
	if *position == 0 {
		if rangeC > 100 {
			*finalPassword += rangeC / 100
		} else {
			*finalPassword += 1
		}
	} else if *position > oldP {
		if rangeC > 100 {
			*finalPassword += rangeC / 100
		} else {
			*finalPassword += 1
		}
	}
}

func Solver(path string) int {
	iter, err := helper.GetInput(1, path)
	if err != nil {
		panic(err)
	}

	finalPassword := 0
	startingPoint := 50

	fmt.Printf("initial position of pointer=%d\n", startingPoint)
	for newString := range iter {
		position := newString[0]
		rangeC, _ := strconv.Atoi(strings.TrimRight(newString[1:], "\n"))
		turn(position, rangeC, &startingPoint, &finalPassword)
	}
	return finalPassword
}
