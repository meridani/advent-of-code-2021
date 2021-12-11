package main

import (
	"fmt"
	"log"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 11

type xy struct {
	x int
	y int
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0

	grid := make(map[xy]int)
	for y, line := range lines {
		nums := pkg.ToIntSliceCharacter(line)
		for x, level := range nums {
			grid[xy{x, y}] = level
		}
	}

	var flashes int
	step := 0
	for {
		step++
		flashed := make(map[xy]bool)
		var flash func(pos xy)
		flash = func(pos xy) {
			if flashed[pos] {
				return
			}
			flashed[pos] = true
			flashes++
			// adj
			for _, dx := range []int{1, 0, -1} {
				for _, dy := range []int{1, 0, -1} {
					if dx == 0 && dy == 0 {
						continue
					}
					np := xy{pos.x + dx, pos.y + dy}
					if _, ok := grid[np]; !ok {
						continue
					}
					if flashed[np] {
						continue
					}
					grid[np] = grid[np] + 1
					if grid[np] > 9 {
						flash(np)
					}
				}
			}
		}
		for pos, val := range grid {
			grid[pos] = val + 1
			if grid[pos] > 9 {
				flash(pos)
			}
		}
		if len(flashed) == len(grid) {
			part2 = step
			break
		}
		for p := range flashed {
			grid[p] = 0
		}
		if step == 100 {
			part1 = flashes
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
