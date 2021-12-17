package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 16

type Packet struct {
	version    byte
	typeID     byte
	ltypeID    byte
	literal    uint64
	length     uint64
	subpackets []Packet
}

type Stream struct {
	scanner *bufio.Scanner
	buff    string
	backlog []uint64
	max     int
	cur     int
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	// lines := input.AsStringSlice()

	part1 := 0
	part2 := 0

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
