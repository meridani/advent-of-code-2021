package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 5

type Point struct {
	x, y int
}

func parseLine(line string) (Point, Point) {
	parts := strings.Split(line, " -> ")
	return parsePoint(parts[0]), parsePoint(parts[1])
}

func parsePoint(point string) Point {
	parts := strings.Split(point, ",")
	return Point{x: pkg.MustAtoi(parts[0]), y: pkg.MustAtoi(parts[1])}
}

func fillFloor(lines []string, diagonal bool) *map[Point]int {
	oceanFloor := make(map[Point]int)
	for _, line := range lines {
		start, end := parseLine(line)
		dx, dy := 0, 0
		if end.x < start.x {
			dx = -1
		} else if start.x < end.x {
			dx = 1
		}
		if end.y < start.y {
			dy = -1
		} else if start.y < end.y {
			dy = 1
		}
		if !diagonal && (start.x != end.x && start.y != end.y) {
			continue
		}
		for {
			oceanFloor[start]++
			if start.x == end.x && start.y == end.y {
				break
			}
			start.x += dx
			start.y += dy
		}
	}

	return &oceanFloor
}

func countDanger(oceanFloor *map[Point]int, danger int) int {
	sum := 0
	for _, point := range *oceanFloor {
		if point > danger {
			sum++
		}
	}
	return sum
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0
	oceanFloor := fillFloor(lines, false)
	part1 = countDanger(oceanFloor, 1)
	oceanFloor = nil
	oceanFloor = fillFloor(lines, true)
	part2 = countDanger(oceanFloor, 1)

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
