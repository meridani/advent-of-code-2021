package main

import (
	"fmt"
	"log"
	"math"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 15

type point struct {
	x, y int
}

var directions = []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func lowestPath(cave [][]int) int {
	queue := [][]point{}
	shortest := math.MaxInt

	queue = append(queue, []point{{0, 1}})
	queue = append(queue, []point{{1, 0}})
	maxX := len(cave[0])
	maxY := len(cave)

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]
		newQueue := make([][]point, 0, len(queue)+4)
		last := current[len(current)-1]

		for _, p := range directions {
			np := point{last.x + p.x, last.y + p.y}
			if np.x == maxX-1 && np.y == maxY-1 {
				current = append(current, np)
				len := pathLength(current, cave)
				if len < shortest {
					shortest = len
				}
				break
			}
			if np.x == 0 && np.y == 0 {
				continue
			}

			if np.x < 0 || np.y < 0 || np.x >= maxX || np.y >= maxY {
				continue
			}

			if !contains(np, current) {
				newCurrent := make([]point, len(current), len(current)+1)
				copy(newCurrent, current)
				newCurrent = append(newCurrent, np)
				newQueue = append(newQueue, newCurrent)
			}
		}
		newQueue = append(newQueue, queue...)
		queue = newQueue
	}
	return shortest
}

func pathLength(p []point, cave [][]int) int {
	sum := 0
	for _, v := range p {
		sum += cave[v.y][v.x]
	}
	return sum
}

func contains(p point, c []point) bool {
	for _, v := range c {
		if p.x == v.x && p.y == v.y {
			return true
		}
	}
	return false
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0
	cave := [][]int{}
	for _, line := range lines {
		cave = append(cave, pkg.ToIntSliceCharacter(line))
	}
	part1 = lowestPath(cave)

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
