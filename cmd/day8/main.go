package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gitchander/permutation"
	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 8

const (
	DIGIT_1   = 2
	DIGIT_4   = 4
	DIGIT_7   = 3
	DIGIT_8   = 7
	DIGIT_069 = 6
	DIGIT_235 = 5
)

func decodeDigit(digit string) int {
	digit = strings.TrimSpace(digit)
	switch len(digit) {
	case DIGIT_1:
		return 1
	case DIGIT_4:
		return 4
	case DIGIT_7:
		return 7
	case DIGIT_8:
		return 8
	default:
		return -1
	}
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			digits := strings.Fields(parts[1])
			for _, number := range digits {
				if decodeDigit(number) > 0 {
					sum++
				}
			}
		}
	}
	return sum
}

func tryPermutation(input []string) {
	for _, line := range input {
		parts := strings.Split(" | ")
		if len(parts) == 2 {
			digits := strings.Fields(parts[0])
			perm := permutation.New(permutation.StringSlice{digits})
		}
	}
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := part1(lines)
	part2 := 0

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
