package pkg

import (
	"math"
	"testing"
)

var MaxTests = []struct {
	name  string
	input []int
	want  int
}{
	{name: "Empty slice", input: []int{}, want: 0},
	{name: "1pcs 0", input: []int{0}, want: 0},
	{name: "2pcs 0", input: []int{0, 0}, want: 0},
	{name: "1,2,3", input: []int{1, 2, 3}, want: 3},
	{name: "Big num", input: []int{math.MaxInt, 0, 1}, want: math.MaxInt},
	{name: "Empty slice", input: []int{}, want: 0},
}

func TestMax(t *testing.T) {
	for _, test := range MaxTests {
		defer func() {
			if r := recover(); r != nil {
				t.Log("Recovered from Max panic")
			}
		}()
		got := Max(test.input...)
		if got != test.want {
			t.Errorf("%v failed: got %v want %v", test.name, test.want, got)
		}
	}
}
