package pkg

func ShiftSliceLeft(s *[]int, num int) {
	for {
		first := (*s)[0]
		*s = (*s)[1:]
		(*s) = append((*s), first)
		num--
		if num == 0 {
			break
		}
	}
}
func ShiftSliceRight(s *[]int, num int) {
	for {
		last := (*s)[len((*s))-1]
		*s = (*s)[:len(*s)-1]
		(*s) = append([]int{last}, (*s)...)
		num--
		if num == 0 {
			break
		}
	}
}

func ChangeToIntSlice(ar []float64) []int {
	newar := make([]int, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = int(v)
	}
	return newar
}
func ChangeToFloatSlice(ar []int) []float64 {
	newar := make([]float64, len(ar))
	var v int
	var i int
	for i, v = range ar {
		newar[i] = float64(v)
	}
	return newar
}
