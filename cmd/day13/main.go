package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 13

type point struct {
	x, y int
}

type instruction struct {
	x, y int
}

func decodePaper(lines []string) ([]point, []instruction) {
	points := []point{}
	instructions := []instruction{}
	for _, line := range lines {
		if len(line) < 3 {
			continue
		}
		parts := strings.Split(strings.TrimSpace(line), ",")
		if len(parts) == 2 {
			x := pkg.MustAtoi(parts[0])
			y := pkg.MustAtoi(parts[1])
			points = append(points, point{x, y})
		} else {
			parts = strings.Split(line, " ")
			parts = strings.Split(parts[2], "=")
			num := pkg.MustAtoi(parts[1])
			if parts[0] == "x" {
				instructions = append(instructions, instruction{num, 0})
			} else {
				instructions = append(instructions, instruction{0, num})
			}
		}
	}
	return points, instructions
}

func fold(points *[]point, instructions *[]instruction) {
	maxX, maxY := 0, 0
	for _, p := range *points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	newpoints := []point{}
	in := (*instructions)[0]
	*instructions = (*instructions)[1:]

	for _, p := range *points {
		dx, dy := 0, 0
		np := point{}
		if in.x != 0 && p.x > in.x {
			dx = in.x - p.x
			np = point{in.x + dx, p.y}
		} else if in.y != 0 && p.y > in.y {
			dy = in.y - p.y
			np = point{p.x, in.y + dy}
		} else {
			np = point{p.x, p.y}
		}
		if !contains(np, newpoints) {
			newpoints = append(newpoints, np)
		}
	}
	*points = newpoints
}

func contains(p point, points []point) bool {
	for _, c := range points {
		if c.x == p.x && c.y == p.y {
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
	points, instructions := decodePaper(lines)
	fold(&points, &instructions)
	part1 = len(points)
	for range instructions {
		fold(&points, &instructions)
	}
	maxX, maxY := 0, 0
	for _, p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if contains(point{x, y}, points) {
				fmt.Print("█")
			} else {
				fmt.Print("░")
			}
		}
		fmt.Println()
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
