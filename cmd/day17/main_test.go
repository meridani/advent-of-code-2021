package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/meridani/advent-of-code-2021/pkg"
)

var RunTests = []struct {
	name         string
	input        pkg.Input
	want1, want2 int
}{
	{name: "Example 1", input: pkg.Input(``), want1: 0, want2: 0},
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

func BenchmarkVmaxy(b *testing.B) {
	input, _ := pkg.ReadInput("input.txt")
	var minX, minY, maxX, maxY int
	fmt.Sscanf(string(input), "target area: x=%d..%d, y=%d..%d", &minX, &maxX, &minY, &maxY)
	for n := 0; n < b.N; n++ {
		vmaxy(minY)
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
