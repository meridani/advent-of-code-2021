package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 8

const (
	DIGIT_1   = 2
	DIGIT_4   = 4
	DIGIT_7   = 3
	DIGIT_8   = 7
	DIGIT_069 = 6
	DIGIT_235 = 5
)

var digitStrings = map[string]int{
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
	"abcefg":  0,
}

type decoder struct {
	secret map[rune]string
}

//  aaaa
// b    c
// b    c
//  dddd
// e    f
// e    f
//  gggg

func (d *decoder) updateKeys(values string) {
	if d.secret == nil {
		d.secret = make(map[rune]string)
	}
	values = pkg.SortString(values)

	switch len(values) {
	case DIGIT_1:
		d.reduceKeys('c', values)
		d.reduceKeys('f', values)
	case DIGIT_4:
		d.reduceKeys('b', values)
		d.reduceKeys('c', values)
		d.reduceKeys('d', values)
		d.reduceKeys('f', values)
	case DIGIT_7:
		d.reduceKeys('a', values)
		d.reduceKeys('c', values)
		d.reduceKeys('f', values)
	case DIGIT_8:
		d.reduceKeys('a', values)
		d.reduceKeys('b', values)
		d.reduceKeys('c', values)
		d.reduceKeys('d', values)
		d.reduceKeys('e', values)
		d.reduceKeys('f', values)
		d.reduceKeys('g', values)
	case DIGIT_069:
		d.reduceKeys('a', values)
		d.reduceKeys('b', values)
		d.reduceKeys('f', values)
		d.reduceKeys('g', values)
	case DIGIT_235:
		d.reduceKeys('a', values)
		d.reduceKeys('d', values)
		d.reduceKeys('g', values)
	}
}

func (d *decoder) removeDuplicates() {
	valuesToRemove := []string{}
	for c_k, c_v := range d.secret {
		if len(c_v) == 1 {
			for i_k, i_v := range d.secret {
				if c_k != i_k {
					if strings.Contains(i_v, c_v) {
						d.secret[i_k] = strings.Replace(i_v, c_v, "", -1)
					}
				}
			}
		}
		if len(c_v) == 2 {
			for i_k, i_v := range d.secret {
				if c_k != i_k && c_v == i_v {
					if !pkg.SliceContainsString(valuesToRemove, c_v) {
						valuesToRemove = append(valuesToRemove, c_v)
					}
				}
			}
		}
		// Including this slows down the execution from 7ms to 11ms
		// if len(c_v) == 3 {
		// 	for i_k, i_v := range d.secret {
		// 		for j_k, j_v := range d.secret {
		// 			if c_k != i_k && i_k != j_k && c_k != j_k &&
		// 				c_v == i_v && i_v == j_v && c_v == j_v {
		// 				if !pkg.SliceContainsString(valuesToRemove, c_v) {
		// 					valuesToRemove = append(valuesToRemove, c_v)
		// 				}
		// 			}
		// 		}
		// 	}
		// }
	}
	if len(valuesToRemove) > 0 {

		for c_k, c_v := range d.secret {
			for _, value := range valuesToRemove {
				if c_v != value {
					for _, letter := range value {
						if strings.Contains(c_v, string(letter)) {
							d.secret[c_k] = strings.Replace(d.secret[c_k], string(letter), "", -1)
						}
					}
				}
			}
		}
	}
}

func (d *decoder) getKeyFromValue(search string) rune {
	for k, v := range d.secret {
		if v == search {
			return k
		}
	}
	return '0'
}

func (d *decoder) decode(input string) string {
	output := []rune{}
	for _, v := range input {
		decodedLetter := d.getKeyFromValue(string(v))
		output = append(output, rune(decodedLetter))
	}
	return pkg.SortString(string(output))
}

func (d *decoder) reduceKeys(key rune, values string) {
	if d.secret[key] == "" {
		d.secret[key] = values
	}
	newValue := []rune{}
	for _, v := range values {
		if strings.Contains(d.secret[key], string(v)) {
			newValue = append(newValue, v)
		}
	}
	d.secret[key] = string(newValue)
	d.removeDuplicates()
}

func (d *decoder) checkKeys() bool {
	if len(d.secret) != 7 {
		return false
	}
	for _, key := range d.secret {
		if len(key) != 1 {
			return false
		}
	}
	return true
}

func solve(input []string) (int, int) {
	sum := 0
	sumDigits := 0
	for _, line := range input {
		keyDecoder := decoder{}
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			digits := strings.Fields(parts[0])
			for _, number := range digits {
				keyDecoder.updateKeys(number)
				if keyDecoder.checkKeys() {
					break
				}
			}
			output := strings.Fields(parts[1])
			var outputDigits strings.Builder
			for _, number := range output {
				digit := digitStrings[keyDecoder.decode(number)]
				outputDigits.WriteString(fmt.Sprint(digit))
				if digit == 1 || digit == 4 || digit == 7 || digit == 8 {
					sumDigits++
				}
			}
			sum += pkg.MustAtoi(outputDigits.String())
		}
	}
	return sumDigits, sum
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()
	part1, part2 := solve(lines)

	return part1, part2
}

func main() {
	pkg.CheckAndDownloadFile("input.txt", fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", DAY))
	input, err := pkg.ReadInput("input.txt")
	if err != nil {
		log.Fatal("Input file does not exists")
	}
	execute.Run(run, input)
}
