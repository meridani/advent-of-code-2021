package main

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 7

func factor(n int) int {
	ret := 0
	for i := 1; i <= n; i++ {
		ret += i
	}
	return ret
}
func calculateFuelUsage(pos *[]int, useFactor bool) int {

	min := math.MaxInt
	for i := (*pos)[0]; i < (*pos)[len(*pos)-1]; i++ {
		total := 0
		for _, v := range *pos {
			if useFactor {
				total += factor(pkg.Abs(v - i))
			} else {
				total += pkg.Abs(v - i)
			}
		}
		if total < min {
			min = total
		}
	}
	return min
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	// lines := input.AsStringSlice()
	crabs := pkg.ToIntSlice(string(input), ",")
	part1 := 0
	part2 := 0
	sort.Ints(crabs)
	part1 = calculateFuelUsage(&crabs, false)
	part2 = calculateFuelUsage(&crabs, true)

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
