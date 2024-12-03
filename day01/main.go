package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	ans, err := part1(input)

	if err != nil {
		panic(err)
	}

	fmt.Println("Output:", ans)
}

func part1(input string) (int, error) {
	sum := 0
	input = strings.TrimRight(input, "\n")
	lines := strings.Split(input, "\n")

	var leftNumbers []int
	var rightNumbers []int
	for i := 0; i < len(lines); i++ {

		leftNumber, err := getSideNumber(lines[i], "l")
		if err != nil {
			return 0, err
		}
		leftNumbers = append(leftNumbers, leftNumber)

		rightNumber, err := getSideNumber(lines[i], "r")
		if err != nil {
			return 0, err
		}
		rightNumbers = append(rightNumbers, rightNumber)
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	for i := 0; i < len(lines); i++ {
		sum = sum + absVal(rightNumbers[i]-leftNumbers[i])
	}

	return sum, nil
}

func getSideNumber(s string, side string) (int, error) {
	if side == "l" {
		return strconv.Atoi(
			strings.TrimSpace(strings.TrimRightFunc(s, func(r rune) bool {
				return unicode.IsNumber(r)
			})))
	} else {
		return strconv.Atoi(
			strings.TrimSpace(strings.TrimLeftFunc(s, func(r rune) bool {
				return unicode.IsNumber(r)
			})))
	}

}

func absVal(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
