package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 9

type point struct {
	x   int
	y   int
	val int
}

var direction = []struct {
	x int
	y int
}{
	{x: 1, y: 0},  // Right
	{x: -1, y: 0}, // Left
	{x: 0, y: 1},  // Up
	{x: 0, y: -1}, // Down
}

func (p *point) equal(c *point) bool {
	if p.x != c.x || p.y != c.y {
		return false
	}
	return true
}

func (p *point) visited(list *[]point) bool {
	for _, l := range *list {
		if p.equal(&l) {
			return true
		}
	}
	return false
}

func (p *point) findChildrenNum(input *[][]int) int {

	queue := []point{}
	visited := []point{}
	queue = append(queue, *p)
	xMax := len((*input)[0])
	yMax := len(*input)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, d := range direction {
			newPoint := point{}

			if xMax > current.x+d.x && current.x+d.x >= 0 &&
				yMax > current.y+d.y && current.y+d.y >= 0 {
				newPoint.x = current.x + d.x
				newPoint.y = current.y + d.y

				testPos := (*input)[newPoint.y][newPoint.x]
				if testPos != 9 && !newPoint.visited(&visited) && !newPoint.visited(&queue) {
					newPoint.val = testPos
					queue = append(queue, newPoint)
				}
			}
		}
		visited = append(visited, current)
	}
	return len(visited)
}

func readHeightMap(lines *[]string) (int, int) {

	linesAsInt := [][]int{}

	for _, line := range *lines {
		linesAsInt = append(linesAsInt, pkg.ToIntSliceCharacter(line))
	}
	minimums := []point{}
	for i, line := range linesAsInt {
		for j, height := range line {
			points := []int{}
			points = append(points, height)
			if i == 0 {
				if j == 0 {
					points = append(points, line[j+1])
					points = append(points, linesAsInt[i+1][j])
				} else if j == len(line)-1 {
					points = append(points, line[j-1])
					points = append(points, linesAsInt[i+1][j])
				} else {
					points = append(points, line[j+1])
					points = append(points, line[j-1])
					points = append(points, linesAsInt[i+1][j])
				}

			} else if i == len(*lines)-1 {
				if j == 0 {
					points = append(points, line[j+1])
					points = append(points, linesAsInt[i-1][j])
				} else if j == len(line)-1 {
					points = append(points, line[j-1])
					points = append(points, linesAsInt[i-1][j])
				} else {
					points = append(points, line[j+1])
					points = append(points, line[j-1])
					points = append(points, linesAsInt[i-1][j])
				}
			} else {
				if j == 0 {
					points = append(points, line[j+1])
					points = append(points, linesAsInt[i+1][j])
					points = append(points, linesAsInt[i-1][j])
				} else if j == len(line)-1 {
					points = append(points, line[j-1])
					points = append(points, linesAsInt[i+1][j])
					points = append(points, linesAsInt[i-1][j])
				} else {
					points = append(points, line[j+1])
					points = append(points, line[j-1])
					points = append(points, linesAsInt[i+1][j])
					points = append(points, linesAsInt[i-1][j])
				}
			}
			if height == pkg.Min(points...) && height < pkg.Max(points...) {
				minimums = append(minimums, point{x: j, y: i, val: height})
			}
		}
	}
	sum := 0
	basins := []int{}
	for _, min := range minimums {
		sum += (min.val + 1)
		basins = append(basins, min.findChildrenNum(&linesAsInt))
	}
	sort.Ints(basins)
	part2 := pkg.Multiply(basins[len(basins)-3:]...)

	return sum, part2
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()
	part1, part2 := readHeightMap(&lines)

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
