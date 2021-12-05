package pkg

import (
	"strconv"
)

func MustAtoi(s string) int {
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
