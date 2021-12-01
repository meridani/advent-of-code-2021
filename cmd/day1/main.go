package main

import (
	"fmt"
	"log"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 1

func run(input pkg.Input) (interface{}, interface{}) {

	numbers := input.AsIntSlice()
	// lines := input.AsStringSlice()

	part1 := 0
	part2 := 0

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			part1++
		}
	}
	for i := 3; i < len(numbers); i++ {
		if numbers[i]+numbers[i-1]+numbers[i-2] > numbers[i-1]+numbers[i-2]+numbers[i-3] {
			part2++
		}
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
