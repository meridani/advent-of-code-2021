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
	{"smallgrid", pkg.Input(`123
	456
	789`), 0, 0},
	{"step 1", pkg.Input(`11111
	19991
	19191
	19991
	11111`), 0, 0},

	{"step 2", pkg.Input(`3454340004500054000434543`), 0, 0},

	{"largeTest 1", pkg.Input(`5483143223
	2745854711
	5264556173
	6141336146
	6357385478
	4167524645
	2176841721
	6882881134
	4846848554
	5283751526`), 0, 0},

	{"largeTest 2", pkg.Input(`6594254334
	3856965822
	6375667284
	7252447257
	7468496589
	5278635756
	3287952832
	7993992245
	5957959665
	6394862637`), 0, 0},

	{name: "Example 1", input: pkg.Input(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`), want1: 1656, want2: 0},
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

func BenchmarkRun(b *testing.B) {
	input, _ := pkg.ReadInput("input.txt")
	for n := 0; n < b.N; n++ {
		run(input)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
