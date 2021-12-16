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
	{"all 1", pkg.Input(`111
111
111`), 4, 136},
	{"smol", pkg.Input(`123
456
789`), 20, 97},
	{"bigur", pkg.Input(`1111
2221
2221
2221`), 6, 153},
	{"more bigur", pkg.Input(`19111
19191
19191
19191
11191`), 16, 212},

	{name: "Example 1", input: pkg.Input(`1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`), want1: 40, want2: 315},
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
