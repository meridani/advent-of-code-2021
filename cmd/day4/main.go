package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 4

type bingoBoard struct {
	board         []int
	markedNumbers map[int]bool
	winner        bool
}

func (b *bingoBoard) createMap() {
	b.markedNumbers = make(map[int]bool)
}

func (b *bingoBoard) markNumber(num int) {
	if b.markedNumbers != nil {
		b.markedNumbers[num] = true
	}
}

func (b *bingoBoard) isComplete() bool {
	return len(b.board) == 25
}

func (b *bingoBoard) updateBoard(newline []int) bool {
	if len(b.board)+len(newline) <= 25 {
		b.board = append(b.board, newline...)
		return true
	}
	return false
}
func (b *bingoBoard) drawNumber(number int) {
	if contains(b.board, number) {
		b.markNumber(number)
	}
}

func (b *bingoBoard) checkWin() bool {

	for i := 0; i < 5; i++ {
		if isSliceWinner(b.board[(i*5):(i*5+5)], b.markedNumbers) {
			b.winner = true
			return true
		}
	}
	temp := []int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			temp = append(temp, b.board[i+j*5])
		}
		if isSliceWinner(temp, b.markedNumbers) {
			b.winner = true
			return true
		} else {
			temp = nil
		}
	}
	return false
}

func isSliceWinner(slice []int, numbers map[int]bool) bool {
	for _, elem := range slice {
		if !numbers[elem] {
			return false
		}
	}
	return true
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func createBoard() *bingoBoard {
	board := new(bingoBoard)
	board.createMap()
	return board
}

func readBoards(lines []string) (boards []bingoBoard) {
	board := createBoard()
	for i, line := range lines {
		lineSlice := strings.Fields(line)
		ints := []int{}
		for _, num := range lineSlice {
			number, err := strconv.Atoi(num)
			if err == nil {
				ints = append(ints, number)
			}
		}
		if len(ints) == 5 && i > 1 {
			board.updateBoard(ints)
			if board.isComplete() {
				boards = append(boards, *board)
			}
		} else {
			board = createBoard()
		}
	}

	return boards
}

func readDrawNumbers(line string) (numbers []int) {
	if line == "" {
		return []int{}
	}
	numbers = pkg.ToIntSlice(line, ",")
	return numbers
}

func playGame(drawNumbers []int, boards []bingoBoard) (bingoBoard, int, bingoBoard, int) {
	winningBoard := bingoBoard{}
	lastBoard := bingoBoard{}
	firstWinner := false
	winningNumber, lastNumber := 0, 0
	winners := 0
	for _, num := range drawNumbers {
		for i := 0; i < len(boards); i++ {
			if !boards[i].winner {
				boards[i].drawNumber(num)
				if boards[i].checkWin() {
					if !firstWinner {
						winningBoard = boards[i]
						winningNumber = num
						firstWinner = true
					}
					winners++
				}
			}
		}
		if winners == len(boards)-1 {
			for _, loserBoard := range boards {
				if !loserBoard.winner {
					// loserBoard.drawNumber(num)
					lastBoard = loserBoard
					// lastNumber = num
				}
			}
		}
		if winners == len(boards) {
			lastNumber = num
			break
		}
	}
	return winningBoard, winningNumber, lastBoard, lastNumber
}

func getSumBoard(board *bingoBoard) (sum int) {
	for _, number := range board.board {
		if !board.markedNumbers[number] {
			sum += number
		}
	}
	return sum
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0

	if len(lines) > 2 {
		drawNumbers := readDrawNumbers(lines[0])
		boards := readBoards(lines)
		board, winningNumber, losingBoard, losingNumber := playGame(drawNumbers, boards)
		part1 = getSumBoard(&board) * winningNumber
		part2 = getSumBoard(&losingBoard) * losingNumber
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
