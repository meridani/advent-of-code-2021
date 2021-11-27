package main

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/meridani/advent-of-code-2021/pkg"
)

type testData struct {
	name    string
	input   pkg.Input
	output1 string
	output2 string
	err     error
}

var i pkg.Input = pkg.Input("")

var example string = `1721
979
366
299
675
1456`

var testDatas = []testData{
	{
		name:    "Too Few arguments",
		input:   i.FromIntSlice([]int{1, 2}),
		output1: "",
		output2: "",
		err:     errors.New("there are less than 3 input numbers to compare"),
	},
	{
		name:    "example data 1",
		input:   pkg.Input(example),
		output1: "514579",
		output2: "",
		err:     nil,
	},
	{
		name:    "input file",
		input:   "input.txt",
		output1: "1020084",
		output2: "295086480",
		err:     nil,
	},
}

func TestRun(t *testing.T) {
	for _, data := range testDatas {

		if data.input == "input.txt" {
			data.input, _ = pkg.ReadInput("input.txt")
		}

		p1, p2, err := run(data.input)
		p1string := fmt.Sprintf("%v", p1)
		p2string := fmt.Sprintf("%v", p2)

		if data.output1 != "" && data.output1 != p1string {
			t.Errorf("In test [%v] Part 1 got: %v, want %v", data.name, p1string, data.output1)
		}
		if data.output2 != "" && data.output2 != p2string {
			t.Errorf("In test [%v] Part 2 got: %v, want %v", data.name, p2string, data.output2)
		}
		if data.err != nil && data.err.Error() != err.Error() {
			t.Errorf("In test [%v] Error test got: %v, want %v", data.name, err, data.err)
		}
		// t.Logf("Test [%v] passed", data.name)
	}
}

func BenchmarkRun(b *testing.B) {
	input, _ := pkg.ReadInput("input.txt")
	for n := 0; n < b.N; n++ {
		run(input)
	}
}

// func TestMainFunc(t *testing.T) {
// 	main()
// }

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
