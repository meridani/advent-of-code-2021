package main

import (
	"fmt"
	"log"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 4

type bingoBoard struct {
	board       [][]int
	drawnNumers [][]bool
	win         bool
	winningNr   int
}

func (b *bingoBoard) buildBoard(nums []int) {
	if len(b.board) <= 5 {
		b.board = append(b.board, nums)
	}
}

func getEmptyBingoBoard() bingoBoard {
	board := new(bingoBoard)
	return *board
}

func readDrawNumbers(line string) (numbers []int) {
	if line == "" {
		return []int{}
	}
	numbers = pkg.ToIntSlice(line, ",")
	return numbers
}

func buildBoards(lines []string) (boards []bingoBoard) {
	if len(lines) < 5 {
		return nil
	}
	board := bingoBoard{}
	for _, line := range lines {

	}

	return boards
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0

	if len(lines) > 0 {
		drawNumbers := readDrawNumbers(lines[0])
		boards := buildBoards(lines[2:])
	}

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
