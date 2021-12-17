package main

import (
	"fmt"
	"log"
	"math"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 17

type xy struct {
	x, y int
}

func vmaxy(y int) int {
	// return int(0.5 * float32(-y) * float32(-y-1))
	return -y * (-y - 1) / 2
}

func reaches(v xy, minX, minY, maxX, maxY int) bool {

	pos := xy{}
	for {
		pos.x += v.x
		pos.y += v.y

		if pos.x >= minX && pos.x <= maxX && pos.y >= minY && pos.y <= maxY {
			return true
		}
		if pos.x > maxX || pos.y < minY {
			break
		}
		if v.x > 0 {
			v.x--
		}
		// if v.x < 0 {
		// 	v.x++
		// }
		v.y--
	}

	return false
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	// lines := input.AsStringSlice()

	part1 := 0
	part2 := 0
	var minX, minY, maxX, maxY int
	fmt.Sscanf(string(input), "target area: x=%d..%d, y=%d..%d", &minX, &maxX, &minY, &maxY)
	vxmin := int(math.Sqrt(float64(minX * 2)))
	vxmax := maxX
	vymin := minY
	vymax := (-minY)

	sum := 0
	for x := vxmin; x <= vxmax; x++ {
		for y := vymin; y <= vymax; y++ {
			p := xy{x, y}
			if reaches(p, minX, minY, maxX, maxY) {
				sum++
			}
		}
	}

	part1 = vmaxy(minY)
	part2 = sum
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
