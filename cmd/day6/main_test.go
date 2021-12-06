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
	{name: "No input", input: pkg.Input(""), want1: 0, want2: 0},
	{name: "Example 18 days", input: pkg.Input(`3,4,3,1,2`), want1: 26, want2: 0},
	{name: "Example 80", input: pkg.Input(`3,4,3,1,2`), want1: 5934, want2: 0},
}

var SimulateDaysTests = []struct {
	name  string
	input []int
	day   int
	want  int
}{
	{name: "1day", input: []int{3, 4, 3, 0, 2}, day: 1, want: 6},
	{name: "2day", input: []int{3, 4, 3, 1, 2}, day: 2, want: 6},
	{name: "3day", input: []int{3, 4, 3, 1, 2}, day: 3, want: 7},
	{name: "4day", input: []int{3, 4, 3, 1, 2}, day: 4, want: 9},
	{name: "5day", input: []int{3, 4, 3, 1, 2}, day: 5, want: 10},
	{name: "6day", input: []int{3, 4, 3, 1, 2}, day: 6, want: 10},
	{name: "7day", input: []int{3, 4, 3, 1, 2}, day: 7, want: 10},
	{name: "8day", input: []int{3, 4, 3, 1, 2}, day: 8, want: 10},
	{name: "9day", input: []int{3, 4, 3, 1, 2}, day: 9, want: 11},
	{name: "10day", input: []int{3, 4, 3, 1, 2}, day: 10, want: 12},
	{name: "10day", input: []int{3, 4, 3, 1, 2}, day: 80, want: 5934},
}

func TestSimulateDays(t *testing.T) {
	for _, test := range SimulateDaysTests {
		simulateDays(&test.input, test.day)
		if pkg.Sum(test.input...) != test.want {
			t.Errorf("%v failed: got %v want %v", test.name, len(test.input), test.want)
		}
	}
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
