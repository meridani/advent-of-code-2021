package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 16

func decode(s *Stream) Packet {
	s.curr = 0
	version := s.feed(3, false)
	typeID := s.feed(3, false)
	p := Packet{version, TypeID(typeID), -1, -1, -1, version, 0, make([]Packet, 0)}
	if p.typeID == Literal {
		for s.feed(1, false) == 1 {
			s.feed(4, true)
		}
		s.feed(4, true)
		p.bits = s.curr
		p.value = s.accept()
	} else {
		p.ltypeID = LTypeID(s.feed(1, false))
		p.max = s.feed(p.ltypeID.getBits(), false)
		p.bits = s.curr
		keepGoing := true
		for keepGoing {
			subpacket := decode(s)
			p.bits += subpacket.bits
			p.cumsum += subpacket.cumsum
			p.subpackets = append(p.subpackets, subpacket)

			if p.ltypeID == BITCOUNT {
				keepGoing = p.bits-22 < p.max
			} else if p.ltypeID == PACKETCOUNT {
				keepGoing = len(p.subpackets) < p.max
			}
		}
		p.applyOperand()
	}
	return p
}

// ======
// PACKET
// ======
type TypeID int64

const (
	Sum TypeID = iota
	Product
	Min
	Max
	Literal
	Gneq
	Lneq
	Eq
)

type LTypeID int64

const (
	BITCOUNT LTypeID = iota
	PACKETCOUNT
)

func (lt LTypeID) getBits() int {
	switch lt {
	case BITCOUNT:
		return 15
	case PACKETCOUNT:
		return 11
	}
	return 0
}

type Packet struct {
	version    int
	typeID     TypeID
	ltypeID    LTypeID
	max        int // the max length, either of bits or of subpackets
	bits       int // bits used to represent this packet (including subpackets)
	cumsum     int // accumulative version sum of packet + subpackets
	value      int // value of packet, according to operand rules
	subpackets []Packet
}

const MAX_INT = int(^uint(0) >> 1)

func (p *Packet) applyOperand() {
	switch p.typeID {
	case Sum:
		p.value = 0
		for _, sp := range p.subpackets {
			p.value += sp.value
		}
	case Product:
		p.value = 1
		for _, sp := range p.subpackets {
			p.value *= sp.value
		}
	case Min:
		p.value = MAX_INT
		for _, sp := range p.subpackets {
			if sp.value < p.value {
				p.value = sp.value
			}
		}
	case Max:
		p.value = -1
		for _, sp := range p.subpackets {
			if sp.value > p.value {
				p.value = sp.value
			}
		}
	case Gneq:
		p.value = 0
		if p.subpackets[0].value > p.subpackets[1].value {
			p.value = 1
		}
	case Lneq:
		p.value = 0
		if p.subpackets[0].value < p.subpackets[1].value {
			p.value = 1
		}
	case Eq:
		p.value = 0
		if p.subpackets[0].value == p.subpackets[1].value {
			p.value = 1
		}
	}
}

// ============
// INPUT STREAM
// ============
// Input stream allows us to read in hex as binary data, to either discard immediately
// and return the value, or keep in the buffer which can be evaluated later on.
type Stream struct {
	scanner *bufio.Scanner
	buffer  string   // binary string of 0s and 1s
	backlog []uint64 // the next values to feed in
	max     int      // maximum number of bits to read before hard stop
	curr    int      // number of bits read so far
}

// Feeds the next n bytes in to the stream
// keep = false:
//		X where X is the int value of the stream
//		Stream is emptied.
//
// keep = true:
//		returns -1, because stream is not ready to be read.
//		Stream is not emptied
func (s *Stream) feed(n int, keep bool) int {
	// Count the number of bits we feed in
	s.curr += n
	for len(s.backlog) < n {
		r, _ := pkg.Read(s.scanner)
		s.backlog = append(s.backlog, pkg.Hex2Bits(r)...)
	}

	tempBuffer := ""
	toFeed := s.backlog[:n]
	for _, f := range toFeed {
		tempBuffer += strconv.FormatUint(f, 2)
	}
	s.backlog = s.backlog[n:]

	if !keep {
		literal := pkg.Bin2Dec(tempBuffer)
		return literal
	} else {
		s.buffer += tempBuffer
		return -1
	}
}

func (s *Stream) accept() int {
	val := pkg.Bin2Dec(s.buffer)
	s.buffer = ""
	return val
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	// lines := input.AsStringSlice()

	part1 := 0
	part2 := 0
	scanner := pkg.RuneScanner("input.txt")
	s := &Stream{scanner, "", make([]uint64, 0), 0, 0}
	p := decode(s)
	part1 = p.cumsum
	part2 = p.value

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
