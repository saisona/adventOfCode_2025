package helper

import (
	"fmt"
	"iter"
	"os"
	"strings"
)

func GetInput(day int, input string) (iter.Seq[string], error) {
	bytes, err := os.ReadFile(fmt.Sprintf("day_%d/%s", day, input))
	if err != nil {
		return nil, err
	}
	iter := strings.Lines(string(bytes))
	return iter, nil
}

func GetInputToArr(day int, input string) ([]string, error) {
	bytes, err := os.ReadFile(fmt.Sprintf("day_%d/%s", day, input))
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bytes), ","), nil
}
