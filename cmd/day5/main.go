package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 5

type point struct {
	x, y int
}

type ventLine struct {
	x1, y1, x2, y2 int
}

func readLineofVents(lines []string) []ventLine {
	ventLines := []ventLine{}
	for _, line := range lines {
		parts := strings.Split(line, "->")
		if len(parts) == 2 {
			x := strings.Split(parts[0], ",")
			y := strings.Split(parts[1], ",")
			x1, err := pkg.MustAtoi(strings.TrimSpace(x[0]))
			if err != nil {
				fmt.Printf("%v cannot convert to x1 %v\n", line, x)
			}
			y1, err := pkg.MustAtoi(strings.TrimSpace(x[1]))
			if err != nil {
				fmt.Printf("%v cannot convert to y1 %v\n", line, x)
			}
			x2, err := pkg.MustAtoi(strings.TrimSpace(y[0]))
			if err != nil {
				fmt.Printf("%v cannot convert to x2 %v\n", line, x)
			}
			y2, err := pkg.MustAtoi(strings.TrimSpace(y[1]))
			if err != nil {
				fmt.Printf("%v cannot convert to y2 %v\n", line, x)
			}
			vent := ventLine{x1, y1, x2, y2}
			ventLines = append(ventLines, vent)
		}
	}
	return ventLines
}

func getLinePoints(vent ventLine, cross bool) []point {
	points := []point{}
	if vent.x1 == vent.x2 {
		if vent.y1 > vent.y2 {
			for p := vent.y2; p <= vent.y1; p++ {
				points = append(points, point{p, vent.x1})
			}
		} else {
			for p := vent.y1; p <= vent.y2; p++ {
				points = append(points, point{p, vent.x1})
			}
		}
	} else if vent.y1 == vent.y2 {
		if vent.x1 > vent.x2 {
			for p := vent.x2; p <= vent.x1; p++ {
				points = append(points, point{vent.y1, p})
			}
		} else {
			for p := vent.x1; p <= vent.x2; p++ {
				points = append(points, point{vent.y1, p})
			}
		}
	} else if cross {
		if pkg.Abs(vent.x2-vent.x1) != pkg.Abs(vent.y2-vent.y1) {
			fmt.Printf("Not 45 degres: %v\n", vent)
		}
		dx := (vent.x2 - vent.x1) / pkg.Abs(vent.x2-vent.x1)
		dy := (vent.y2 - vent.y1) / pkg.Abs(vent.y2-vent.y1)
		dist := pkg.Abs(vent.x2 - vent.x1)
		for x, y, i := 0, 0, 0; i <= dist; i++ {
			points = append(points, point{vent.y1 + y, vent.x1 + x})
			x += dx
			y += dy
		}
	}
	return points
}

func createMapStraight(vents []ventLine) *([1000][1000]int) {
	var oceanFloor [1000][1000]int

	for _, vent := range vents {
		// only straight lines
		points := getLinePoints(vent, false)
		for _, point := range points {
			oceanFloor[point.y][point.x]++
		}
	}
	return (*[1000][1000]int)(&oceanFloor)
}

func createMap(vents []ventLine) *([1000][1000]int) {
	var oceanFloor [1000][1000]int

	for _, vent := range vents {
		points := getLinePoints(vent, true)
		for _, point := range points {
			oceanFloor[point.y][point.x]++
		}
	}
	return (*[1000][1000]int)(&oceanFloor)
}
func countDanger(oceanFloor *[1000][1000]int, danger int) int {
	sum := 0
	for _, row := range *oceanFloor {
		for _, point := range row {
			if point > danger {
				sum++
			}
		}
	}
	return sum
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0

	vents := readLineofVents(lines)
	oceanFloor := createMapStraight(vents)
	part1 = countDanger(oceanFloor, 1)
	oceanFloor = nil
	part2Floor := createMap(vents)
	part2 = countDanger(part2Floor, 1)

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
