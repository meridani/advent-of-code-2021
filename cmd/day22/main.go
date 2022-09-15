package main

import (
	"fmt"
	"log"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 22

type cube struct {
	command                string
	x1, x2, y1, y2, z1, z2 int
	volume                 int
}

func (c *cube) calculateVolume() {
	xlen := c.x2 - c.x1
	ylen := c.y2 - c.y1
	zlen := c.z2 - c.z1

	c.volume = xlen * ylen * zlen
}

func createCube(command string, x1, x2, y1, y2, z1, z2 int) *cube {
	cube := cube{command, x1, x2, y1, y2, z1, z2, 0}
	cube.calculateVolume()
	return &cube
}

func part1(cubes []*cube) {

}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0

	cubes1 := []cube{}
	cubes2 := []cube{}
	// reg := regexp.MustCompile(`^(.*) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)

	for _, line := range lines {
		var x1, x2, y1, y2, z1, z2 int
		var command string
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &command, &x1, &x2, &y1, &y2, &z1, &z2)
		// matches := reg.FindStringSubmatch(line)
		// cube := createCube(matches[1], pkg.MustAtoi(matches[2]), pkg.MustAtoi(matches[3]),
		// 	pkg.MustAtoi(matches[4]), pkg.MustAtoi(matches[5]), pkg.MustAtoi(matches[6]),
		// 	pkg.MustAtoi(matches[7]))

		cube := createCube(command, x1, x2, y1, y2, z1, z2)
		if pkg.Abs(x1) > 50 || pkg.Abs(x2) > 50 ||
			pkg.Abs(y1) > 50 || pkg.Abs(y2) > 50 ||
			pkg.Abs(z1) > 50 || pkg.Abs(z2) > 50 {
			cubes2 = append(cubes2, *cube)
		} else {
			cubes1 = append(cubes1, *cube)
			cubes2 = append(cubes2, *cube)
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
