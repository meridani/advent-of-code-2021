package main

import (
	"fmt"
	"log"
	"math"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 20

type xy struct {
	x, y int
}

var neighbours = []xy{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 0},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0

	part1 = enhance(&lines, 2)
	part2 = enhance(&lines, 50)

	return part1, part2
}

func enhance(lines *[]string, n int) int {
	algo := (*lines)[0]
	img := make(map[xy]bool)
	for x, line := range (*lines)[2:] {
		for y, pixel := range line {
			img[xy{x, y}] = pixel == '#'
		}
	}

	var maxX, maxY int
	var minX, minY = math.MaxInt, math.MaxInt
	var arg int
	for i := 0; i < n; i++ {

		output := make(map[xy]bool, len(img))

		for pixel := range img {
			maxX = pkg.Max(maxX, pixel.x)
			maxY = pkg.Max(maxY, pixel.y)
			minX = pkg.Min(minX, pixel.x)
			minY = pkg.Min(minY, pixel.y)
		}

		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for _, n := range neighbours {
					arg = arg << 1
					if v, ok := img[xy{x + n.x, y + n.y}]; ok {
						if v {
							arg |= 1
						}
					} else {
						if i%2 == 1 {
							arg |= 1
						}
					}
				}
				if algo[arg] == '#' {
					output[xy{x, y}] = true
				} else {
					output[xy{x, y}] = false
				}
				arg = 0
			}
		}
		img = output
	}

	return count(&img)
}

func count(in *map[xy]bool) (res int) {
	for _, i := range *in {
		if i {
			res++
		}
	}
	return
}

func main() {
	pkg.CheckAndDownloadFile("input.txt", fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", DAY))
	input, err := pkg.ReadInput("input.txt")
	if err != nil {
		log.Fatal("Input file does not exists")
	}
	execute.Run(run, input)
}
