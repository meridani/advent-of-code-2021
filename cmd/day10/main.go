package main

import (
	"fmt"
	"log"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 10

var points = map[rune]int{')': 3, '}': 1197, ']': 57, '>': 25137}
var openings = []rune{'(', '{', '[', '<'}
var closings = []rune{')', '}', ']', '>'}

func isOpening(c rune) bool {
	for _, o := range openings {
		if c == o {
			return true
		}
	}
	return false
}

func getOpeningPos(c rune) int {
	for i, o := range openings {
		if c == o {
			return i
		}
	}
	return -1
}

func isClosing(c rune) bool {
	for _, o := range closings {
		if c == o {
			return true
		}
	}
	return false
}

func getClosingPos(c rune) int {
	for i, o := range closings {
		if c == o {
			return i
		}
	}
	return -1
}

func checkLine(line string) int {

	opens := []rune{}

	for _, c := range line {
		if isOpening(c) {
			opens = append(opens, c)
		} else {
			expected := getOpeningPos(opens[len(opens)-1])
			if c != closings[expected] {
				return points[c]
			}
			opens = opens[:len(opens)-1]
		}
		if len(opens) != 0 {

		}
	}
	return 0
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()
	part1 := 0
	part2 := 0

	for _, line := range lines {
		part1 += checkLine(line)
	}

	return part1, part2
}

func main() {
	pkg.CheckAndDownloadFile("input.txt", fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", DAY))
	input, err := pkg.ReadInput("input.txt")
	if err != nil {
		log.Fatal("Input file does not exists")
	}
	execute.Run(run, input)
}
