package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 16

const (
	literal = 4 // single binary number
)

const (
	VERSION_OFF     = 0
	VERSION_LEN     = 3
	TYPE_OFF        = 3
	TYPE_LEN        = 3
	LENGTH_TYPE_OFF = 6
	LENGTH_TYPE_LEN = 1
	LENGTH_OFF      = 7
	LENGTH_BITS     = 15
	LENGTH_PACKETS  = 11
	LITERAL_OFFSET  = 6
)

type bitsPacket struct {
	version byte
	tID     byte
	literal uint64
	ltID    byte
	length  uint64
	content []bitsPacket
}

func stringToNum(i string) uint64 {
	num, _ := strconv.ParseUint(i, 2, 64)
	return num
}

func decodePacket(binary string) (bitsPacket, int) {
	packet := bitsPacket{}
	if len(binary) < TYPE_OFF+TYPE_LEN {
		return packet, -1
	}
	packet.version = byte(stringToNum(binary[VERSION_OFF : VERSION_OFF+VERSION_LEN]))
	packet.tID = byte(stringToNum(binary[TYPE_OFF : TYPE_OFF+TYPE_LEN]))
	i := 0
	switch packet.tID {

	case literal:
		var numberBuilder strings.Builder
		i = LITERAL_OFFSET
		for {
			numberBuilder.WriteString(binary[i+1 : i+5])
			if len(binary) < i+5 {
				return bitsPacket{}, -1
			}
			if binary[i] == '0' {
				break
			}
			i += 5
		}

		var err error
		packet.literal = stringToNum(numberBuilder.String())

		if err != nil {
			fmt.Println("invalid literal found!")
		}
		i += 5

	default:
		packet.ltID = byte(stringToNum(string(binary[LENGTH_TYPE_OFF])))
		i = LENGTH_OFF
		if packet.ltID == 0 {
			packet.length = stringToNum(binary[i : i+LENGTH_BITS])
			i += LENGTH_BITS
		} else {
			packet.length = stringToNum(binary[i : i+LENGTH_PACKETS])
			i += LENGTH_PACKETS
		}
	}

	return packet, i
}

func toBinaryString(line string) string {
	var binary strings.Builder

	for _, c := range line {

		switch c {
		case '0':
			binary.WriteString("0000")
		case '1':
			binary.WriteString("0001")
		case '2':
			binary.WriteString("0010")
		case '3':
			binary.WriteString("0011")
		case '4':
			binary.WriteString("0100")
		case '5':
			binary.WriteString("0101")
		case '6':
			binary.WriteString("0110")
		case '7':
			binary.WriteString("0111")
		case '8':
			binary.WriteString("1000")
		case '9':
			binary.WriteString("1001")
		case 'A':
			binary.WriteString("1010")
		case 'B':
			binary.WriteString("1011")
		case 'C':
			binary.WriteString("1100")
		case 'D':
			binary.WriteString("1101")
		case 'E':
			binary.WriteString("1110")
		case 'F':
			binary.WriteString("1111")
		}
	}
	return binary.String()
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

	binary := toBinaryString(strings.TrimSpace(string(input)))

	for {
	}

	fmt.Println(packets)

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
