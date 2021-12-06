package main

import (
	"fmt"
	"log"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 6

func simulateDays(school *[]int, days int) {
	for i := 0; i < days; i++ {
		newFish := (*school)[0]
		// using shift is around 5us slower than manual indexing
		// pkg.ShiftSliceLeft(school, 1)
		// (*school)[6] += newFish
		(*school)[0] = (*school)[1]
		(*school)[1] = (*school)[2]
		(*school)[2] = (*school)[3]
		(*school)[3] = (*school)[4]
		(*school)[4] = (*school)[5]
		(*school)[5] = (*school)[6]
		(*school)[6] = (*school)[7] + newFish
		(*school)[7] = (*school)[8]
		(*school)[8] = newFish
	}
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	// lines := input.AsStringSlice()
	fish := pkg.ToIntSlice(string(input), ",")
	school := make([]int, 9)
	for _, f := range fish {
		school[f]++
	}
	simulateDays(&school, 80)

	part1 := pkg.Sum(school...)
	simulateDays(&school, 256-80)
	part2 := pkg.Sum(school...)

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
