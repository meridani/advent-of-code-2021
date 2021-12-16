package main

import (
	"fmt"
	"log"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 16

const (
	literal = 4 // single binary number

)

type bitsPacket struct {
	version byte
	typeID  byte
	packet  []byte
}

func decodePacket(line string) bitsPacket {
	binary := []byte{}
	for _, c := range line {
		if c >= 'A' && c <= 'F' {
			binary = append(binary, byte(c-'A'+10))
		} else {
			binary = append(binary, byte(c-'0'))
		}
	}

	packet := bitsPacket{}
	packet.version = (binary[0] >> 1) & 0b111
	packet.typeID = (binary[0]&0b1)<<2 | (binary[1] >> 2)
	return packet
}

func sumPart1(packets []bitsPacket) int {
	sum := 0
	for _, p := range packets {
		sum += int(p.version)
	}
	return sum
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	// lines := input.AsStringSlice()

	part1 := 0
	part2 := 0

	// part1 = sumPart1(packets)
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
