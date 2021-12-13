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
	{name: "Example small", input: pkg.Input(`start-A
	start-b
	A-c
	A-b
	b-d
	A-end
	b-end`), want1: 10, want2: 36},

	{
		"Example bigger", pkg.Input(`dc-end
	HN-start
	start-kj
	dc-start
	dc-HN
	LN-dc
	HN-end
	kj-sa
	kj-HN
	kj-dc`), 19, 103},

	{
		"Example BIG", pkg.Input(`fs-end
	he-DX
	fs-he
	start-DX
	pj-DX
	end-zg
	zg-sl
	zg-pj
	pj-he
	RW-he
	fs-DX
	pj-RW
	zg-RW
	start-pj
	he-WI
	zg-he
	pj-fs
	start-RW`), 226, 3509},
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
