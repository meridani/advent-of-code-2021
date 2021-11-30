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

	part1 := 0
	part2 := 0

	part1NotFound, part2NotFound := true, true

	for i := 0; i < len(numbers)-2; i++ {
		if !part1NotFound && !part2NotFound {
			break
		}
		for j := i + 1; j < len(numbers)-1; j++ {
			if pkg.Sum(numbers[i], numbers[j]) == 2020 && part1NotFound {
				part1 = pkg.Multiply(numbers[i], numbers[j])
				part1NotFound = false
			}
			for k := j + 1; k < len(numbers); k++ {
				if pkg.Sum(numbers[i], numbers[j], numbers[k]) == 2020 && part2NotFound {
					part2 = pkg.Multiply(numbers[i], numbers[j], numbers[k])
					part2NotFound = false
				}
			}
		}
	}
	return part1, part2
}

func main() {
	pkg.CheckAndDownloadFile("input.txt", fmt.Sprintf("https://adventofcode.com/2020/day/%v/input", DAY))
	input, err := pkg.ReadInput("input.txt")
	if err != nil {
		log.Fatal("Input file does not exists")
	}
	execute.Run(run, input)
}
