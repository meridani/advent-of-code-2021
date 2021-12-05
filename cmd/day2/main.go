package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 2

type Direction int

const (
	Forward = iota
	Up
	Down
)

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	horizontalPos, verticalPos, aim := 0, 0, 0
	depthWithAim := 0

	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), " ")
		value := pkg.MustAtoi(parts[1])

		switch parts[0] {
		case "forward":
			horizontalPos += value
			depthWithAim += value * aim
		case "up":
			verticalPos -= value
			aim -= value
		case "down":
			verticalPos += value
			aim += value
		}
	}

	part1 := horizontalPos * verticalPos
	part2 := depthWithAim * horizontalPos

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
