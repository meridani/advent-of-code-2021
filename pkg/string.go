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
	s = strings.TrimSpace(s)
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
		num := MustAtoi(cur)
		numbers = append(numbers, num)
	}
	return numbers
}
