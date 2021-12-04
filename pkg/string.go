package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MustAtoi(s string) (int, error) {
	if s == "" {
		return 0, errors.New("s can't be empty")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Print("Invalid MustAtoi input...")
	}
	return n, nil
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
