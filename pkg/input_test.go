package pkg

import (
	"reflect"
	"testing"
)

var FromIntSliceTests = []struct {
	name  string
	input []int
	want  Input
}{
	{name: "Empty slice", input: []int{}, want: Input("")},
	{name: "Slice of 1", input: []int{1}, want: Input("1")},
	{name: "Slice of 2", input: []int{1, 2}, want: Input("1\n2")},
}

func TestFromIntSlice(t *testing.T) {

	for _, test := range FromIntSliceTests {
		got := Input("")
		got.FromIntSlice(test.input)
		if string(test.want) != string(got) {
			t.Errorf("%v failed: got %v want %v", test.name, test.want, got)
		}
	}
}

var AsIntSliceTests = []struct {
	name  string
	input Input
	want  []int
}{
	{name: "Empty input", input: Input(""), want: []int{}},
	{name: "One element", input: Input("1"), want: []int{1}},
	{name: "Two element", input: Input("1\n2"), want: []int{1, 2}},
}

func TestAsIntSlice(t *testing.T) {
	for _, test := range AsIntSliceTests {
		got := test.input.AsIntSlice()
		if len(test.want) == len(got) && len(got) == 0 {
			// pass
		} else if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%v failed: got %v want %v", test.name, test.want, got)
		}
	}
}

var AsStringSliceTests = []struct {
	name  string
	input Input
	want  []string
}{
	{name: "Empty input", input: Input(""), want: []string{}},
	{name: "One element", input: Input("1"), want: []string{"1"}},
	{name: "Two element", input: Input("1\n2"), want: []string{"1", "2"}},
}

func TestAsStringSlice(t *testing.T) {
	for _, test := range AsStringSliceTests {
		got := test.input.AsStringSlice()
		if len(test.want) == len(got) && len(got) == 0 {
			// pass
		} else if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%v failed: got %v want %v", test.name, test.want, got)
		}
	}
}
