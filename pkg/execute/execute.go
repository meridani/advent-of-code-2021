package execute

import (
	"fmt"
	"time"

	"github.com/meridani/advent-of-code-2021/pkg"
)

func Run(run func(pkg.Input) (interface{}, interface{}), puzzle pkg.Input) {

	start := time.Now()
	part1, part2 := run(puzzle)

	elapsed := time.Since(start)

	fmt.Printf("PART1: %v\nPART2: %v\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
