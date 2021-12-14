package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 14

func readInsertions(input []string) map[string]string {
	insertions := map[string]string{}

	for _, line := range input {
		parts := strings.Split(strings.TrimSpace(line), " -> ")
		if len(parts) == 2 {
			insertions[parts[0]] = parts[1]
		}
	}
	return insertions
}

func insert(polymer string, ins map[string]string, times int) map[string]uint64 {
	pairs := map[string]uint64{}

	for i := 0; i < len(polymer)-1; i++ {
		pairs[polymer[i:i+2]]++
	}
	for i := 0; i < times; i++ {
		newPairs := map[string]uint64{}
		for k, v := range pairs {
			newPairs[k[:1]+ins[k]] += v
			newPairs[ins[k]+k[1:]] += v
		}
		pairs = newPairs
	}
	return pairs
}

func maxmin(m map[string]uint64, last string) uint64 {
	p := map[string]uint64{}
	for k, v := range m {
		p[k[:1]] += v
	}
	p[last]++

	var min, max uint64

	min, max = math.MaxUint64, 0
	for _, v := range p {
		if min > v {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	// part1 := 0
	// part2 := 0

	polymer := lines[0]
	insertions := readInsertions(lines[2:])

	pairs := insert(polymer, insertions, 10)

	part1 := maxmin(pairs, polymer[len(polymer)-1:])

	pairs = insert(polymer, insertions, 40)

	part2 := maxmin(pairs, polymer[len(polymer)-1:])

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
