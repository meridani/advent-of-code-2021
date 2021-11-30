package pkg

import (
	"errors"
	"strconv"
)

func MustAtoi(s string) (int, error) {
	if s == "" {
		return 0, errors.New("s can't be empty")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n, nil
}
