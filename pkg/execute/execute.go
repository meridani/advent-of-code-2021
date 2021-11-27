package execute

import (
	"fmt"
	"log"
	"time"

	"github.com/meridani/advent-of-code-2021/pkg"
)

func Run(run func(pkg.Input) (interface{}, interface{}, error), puzzle pkg.Input) {

	start := time.Now()
	part1, part2, err := run(puzzle)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)

	fmt.Printf("PART1: %v\nPART2: %v\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
