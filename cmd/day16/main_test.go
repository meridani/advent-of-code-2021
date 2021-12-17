package main

import (
	"os"
	"testing"

	"github.com/meridani/advent-of-code-2021/pkg"
)

var RunTests = []struct {
	name         string
	input        pkg.Input
	want1, want2 int
}{
	{name: "Example 1", input: pkg.Input(`D2FE28`), want1: 0, want2: 0},
}

func TestRun(t *testing.T) {
	for _, test := range RunTests {
		got1, got2 := run(test.input)
		if got1 != test.want1 {
			t.Errorf("%v part 1 failed: got %v want %v", test.name, got1, test.want1)
		}
		if got2 != test.want2 {
			t.Errorf("%v part 2 failed: got %v want %v", test.name, got2, test.want2)
		}
	}

}

func TestRun2(t *testing.T) {
	for _, test := range BitPacketDecoderTests {
		run(pkg.Input(test.input))
	}
}

var BitPacketDecoderTests = []struct {
	name  string
	input string
	want  Packet
}{
	{"1", "D2FE28", Packet{version: 6, typeID: 4}},
	{"2", "38006F45291200", Packet{version: 1, typeID: 6, ltypeID: 0, length: 27,
		subpackets: []Packet{
			{version: 6, typeID: 4, literal: 10},
			{version: 2, typeID: 4, literal: 20},
		}}},
	{"3", "EE00D40C823060", Packet{version: 7, typeID: 3, ltypeID: 1, length: 3,
		subpackets: []Packet{
			{version: 2, typeID: 4, literal: 1},
			{version: 4, typeID: 4, literal: 2},
			{version: 1, typeID: 4, literal: 3},
		}}},
	{"4", "8A004A801A8002F478", Packet{version: 4, typeID: 1}},
	{"5", "620080001611562C8802118E34", Packet{version: 3, typeID: 3}},
	{"6", "C0015000016115A2E0802F182340", Packet{version: 7, typeID: 3}},
	{"7", "A0016C880162017C3686B18A3D4780", Packet{version: 7, typeID: 3}},
}

func TestDecodePacket(t *testing.T) {
	for _, test := range BitPacketDecoderTests {
		binary, _ := toBinaryString(test.input)
		got, _ := decodePacket(binary)
		if got.version != test.want.version {
			t.Errorf("%v failed: got %v want %v", test.name, got, test.want)
		}
	}
}

func BenchmarkRun(b *testing.B) {
	input, _ := pkg.ReadInput("input.txt")
	for n := 0; n < b.N; n++ {
		run(input)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
