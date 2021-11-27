package pkg

import "testing"

func TestMax(t *testing.T) {
	got := Max([]int{1, 2, 3, 4, 5}...)
	if got != 5 {
		t.Errorf("Max(1,2,3,4,5) = %v; want 5", got)
	}
}
