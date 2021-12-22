package main

import (
	"fmt"
	"log"
	"math"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 21

var quantumDie map[int]int = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

type player struct {
	id    int
	pos   int
	score int
}

func (p *player) play(step int) bool {
	p.pos += step
	p.pos = (p.pos-1)%10 + 1
	p.score += p.pos
	return p.score >= 1000
}

type deterministicDice struct {
	rolls int
}

func (d *deterministicDice) roll(n int) int {
	val := 0
	for i := 0; i < n; i++ {
		d.rolls++
		val += d.rolls
	}
	return val
}

type diracDice struct {
	cache map[[4]int][2]int
}

func (d *diracDice) play(currentPos, otherPos, currentPoints, otherPoints int) [2]int {
	if currentPoints > 20 {
		return [2]int{1, 0}
	}
	if otherPoints > 20 {
		return [2]int{0, 1}
	}
	if res, ok := d.cache[[4]int{currentPos, otherPos, currentPoints, otherPoints}]; ok {
		return res
	}
	res := [2]int{0, 0}
	for roll, additional := range quantumDie {
		newPos := (currentPos + roll) % 10
		newPoints := currentPoints + newPos + 1
		wins := d.play(otherPos, newPos, otherPoints, newPoints)
		res[0] += wins[1] * additional
		res[1] += wins[0] * additional
	}
	d.cache[[4]int{currentPos, otherPos, currentPoints, otherPoints}] = res
	return res
}

func playWithDeterministicDice(players []player) int {
	d := deterministicDice{}
	rolls := 0

	for {
		for i := 0; i < len(players); i++ {
			rolls += d.roll(3)

			if players[i].play(rolls) {
				goto win
			}
			rolls = 0
		}
	}
win:
	min := math.MaxInt
	for _, p := range players {
		min = pkg.Min(p.score, min)
	}
	return min * d.rolls
}

func playWithQuantumDice(players []player) uint64 {
	game := diracDice{cache: map[[4]int][2]int{}}
	wins := game.play(players[0].pos-1, players[1].pos-1, 0, 0)
	fmt.Println(wins)
	if wins[0] > wins[1] {
		return uint64(wins[0])
	}
	return uint64(wins[1])
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0
	players := []player{}
	for _, line := range lines {
		newplayer := player{}
		id, pos := 0, 0
		fmt.Sscanf(line, "Player %d starting position: %d", &id, &pos)
		newplayer.id = id
		newplayer.pos = pos
		players = append(players, newplayer)
	}
	players2 := make([]player, 2)
	copy(players2, players)
	part1 = playWithDeterministicDice(players)
	part2 = int(playWithQuantumDice(players2))
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
