package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

func MustAtoi(s string) int {
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Print("Invalid MustAtoi input...")
	}
	return n
}

func ToIntSlice(s string, sep string) []int {
	numbers := []int{}
	split := strings.Split(s, sep)
	for _, cur := range split {
		num, err := MustAtoi(cur)
		if err != nil {
			fmt.Printf("invalid in toIntSlice")
		}
		numbers = append(numbers, num)
	}
	return numbers
}
