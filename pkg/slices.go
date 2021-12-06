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
