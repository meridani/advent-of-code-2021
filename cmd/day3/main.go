package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 3

func countAtPos(lines []string, pos int) (int, int) {
	sum0, sum1 := 0, 0
	for _, line := range lines {
		if line[pos:pos+1] == fmt.Sprint(1) {
			sum1++
		} else {
			sum0++
		}
	}
	return sum1, sum0
}

func numberContainsAtPosition(lines []string, pos int, criteria int) (num string) {
	if len(lines) == 0 {
		return "error"
	}
	if len(lines) == 1 {
		return lines[0]
	}
	ones, zeroes := countAtPos(lines, pos)
	relevant := 0
	switch criteria {
	case 0:
		if zeroes > ones {
			relevant = 1
		} else {
			relevant = 0
		}
	case 1:
		if ones >= zeroes {
			relevant = 1
		} else {
			relevant = 0
		}
	}

	newLines := []string{}
	for _, line := range lines {
		if line[pos:pos+1] == fmt.Sprint(relevant) {
			newLines = append(newLines, line)
		}
	}
	pos++
	return numberContainsAtPosition(newLines, pos, criteria)
}

func findPowerConsumption(lines []string) int {
	if len(lines) < 2 {
		return 0
	}
	gamma, epsilon := []string{}, []string{}
	for i := 0; i < len(lines[0]); i++ {
		ones, zeroes := countAtPos(lines, i)
		if ones >= zeroes {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		} else {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		}
	}
	eps, err := strconv.ParseUint(strings.Join(epsilon, ""), 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	gam, err := strconv.ParseUint(strings.Join(gamma, ""), 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(gam * eps)
}

func verifyLifeSupport(lines []string) int {
	oxygen := numberContainsAtPosition(lines, 0, 1)
	co2 := numberContainsAtPosition(lines, 0, 0)

	oxi, err := strconv.ParseUint(oxygen, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	co, err := strconv.ParseUint(co2, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(oxi * co)
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := findPowerConsumption(lines)
	part2 := verifyLifeSupport(lines)

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
