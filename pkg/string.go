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

func ToFloatSlice(s string, sep string) []float64 {
	numbers := []float64{}
	split := strings.Split(s, sep)
	for _, cur := range split {
		cur = strings.TrimSpace(cur)
		num, err := strconv.ParseFloat(cur, 64)
		if err != nil {
			fmt.Println("Invalid float")
		}
		numbers = append(numbers, num)
	}
	return numbers
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
