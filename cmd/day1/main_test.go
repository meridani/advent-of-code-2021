package main

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/meridani/advent-of-code-2021/pkg"
)

type testData struct {
	name  string
	input pkg.Input
	want1 string
	want2 string
	err   error
}

var i pkg.Input = pkg.Input("")

var example string = `1721
979
366
299
675
1456`

var testDatas = []testData{}

func PrepareTestData(t *testing.T) {
	t.Helper()
	i.FromIntSlice([]int{1, 2})
	testDatas = append(testDatas,
		testData{
			name:  "Too Few arguments",
			input: i,
			want1: "",
			want2: "",
			err:   errors.New("error returned mismatch"),
		})
	testDatas = append(testDatas,
		testData{
			name:  "Too much arguments",
			input: pkg.Input(example),
			want1: "",
			want2: "",
			err:   nil,
		})
	testDatas = append(testDatas,
		testData{
			name:  "input file",
			input: "input.txt",
			want1: "0",
			want2: "0",
			err:   nil,
		})
}

func TestRun(t *testing.T) {
	PrepareTestData(t)
	for _, data := range testDatas {

		if data.input == "input.txt" {
			data.input, _ = pkg.ReadInput("input.txt")
		}

		p1, p2, err := run(data.input)
		p1string := fmt.Sprintf("%v", p1)
		p2string := fmt.Sprintf("%v", p2)

		if data.want1 != "" && data.want1 != p1string {
			t.Errorf("In test [%v] Part 1 got: %v, want %v", data.name, p1string, data.want1)
		}
		if data.want2 != "" && data.want2 != p2string {
			t.Errorf("In test [%v] Part 2 got: %v, want %v", data.name, p2string, data.want2)
		}
		if data.err != nil && err != nil && data.err.Error() != err.Error() {
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
