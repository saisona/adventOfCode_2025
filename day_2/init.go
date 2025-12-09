package day_2

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/saisona/AdventOfCode2025/helper"
)

func isValidInput(val int) bool {
	s := fmt.Sprintf("%d", val)
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%2 == 0 {
			block := s[:i]
			repeated := strings.Repeat(block, len(s)/i)

			isSameStr := strings.Compare(repeated, s) == 0
			blockSizeIsValid := len(block) == len(s)/2

			if isSameStr && blockSizeIsValid {
				return false
			}
		}
	}

	return true
}

func Worker(min int, max int, out chan int, wg *sync.WaitGroup) error {
	defer wg.Done()
	for i := min; i <= max; i++ {
		if !isValidInput(i) {
			out <- i
		}
	}
	return nil
}

func fromStringToMinMax(input string) (int, int) {
	minmaxStr := strings.Split(input, "-")
	min, _ := strconv.Atoi(minmaxStr[0])
	max, _ := strconv.Atoi(minmaxStr[1])
	return min, max
}

func Solver(path string) int {
	iter, err := helper.GetInputToArr(2, path)
	if err != nil {
		panic(err)
	}

	resChan := make(chan int)
	wg := sync.WaitGroup{}
	finalPassword := 0
	wg.Add(len(iter))

	for _, v := range iter {
		min, max := fromStringToMinMax(v)
		go Worker(min, max, resChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	for output := range resChan {
		finalPassword += output
	}

	return finalPassword
}
