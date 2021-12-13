package pkg

import (
	"fmt"
	"sort"
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

func ToIntSliceCharacter(s string) []int {
	ret := []int{}
	for _, char := range s {
		ret = append(ret, int(char-'0'))
	}
	return ret
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func SliceContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
