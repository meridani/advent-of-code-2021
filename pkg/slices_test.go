package pkg

import (
	"reflect"
	"testing"
)

var ShiftSliceLeftTests = []struct {
	name  string
	input []int
	want  []int
	shift int
}{
	{name: "Shift 1", input: []int{1, 2, 3, 4}, want: []int{2, 3, 4, 1}, shift: 1},
	{name: "Shift 2", input: []int{1, 2, 3, 4}, want: []int{3, 4, 1, 2}, shift: 2},
	{name: "Shift 3", input: []int{1, 2, 3, 4}, want: []int{4, 1, 2, 3}, shift: 3},
	{name: "Shift 4", input: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}, shift: 4},
}

func TestShiftSliceLeft(t *testing.T) {
	for _, test := range ShiftSliceLeftTests {
		ShiftSliceLeft(&test.input, test.shift)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("%v failed: got %v want %v", test.name, test.input, test.want)
		}
	}
}

var ShiftSliceRightTests = []struct {
	name  string
	input []int
	want  []int
	shift int
}{
	{name: "Shift 1", input: []int{1, 2, 3, 4}, want: []int{4, 1, 2, 3}, shift: 1},
	{name: "Shift 2", input: []int{1, 2, 3, 4}, want: []int{3, 4, 1, 2}, shift: 2},
	{name: "Shift 3", input: []int{1, 2, 3, 4}, want: []int{2, 3, 4, 1}, shift: 3},
	{name: "Shift 4", input: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}, shift: 4},
}

func TestShiftSliceRight(t *testing.T) {
	for _, test := range ShiftSliceRightTests {
		ShiftSliceRight(&test.input, test.shift)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("%v failed: got %v want %v", test.name, test.input, test.want)
		}
	}
}
