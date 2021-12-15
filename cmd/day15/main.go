package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 15

type point struct {
	x, y int
}
type node struct {
	p   point
	val int
}

type pQueue []node

func (pq pQueue) Len() int { return len(pq) }

func (pq pQueue) Less(i, j int) bool {
	return pq[i].val < pq[j].val
}

func (pq pQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pQueue) Push(x interface{}) {
	*pq = append(*pq, x.(node))
}

func (pq *pQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return item
}

func lowestPath(cave *[][]int) int {

	maxX := len((*cave)[0])
	maxY := len((*cave))
	dist := make([][]int, 0, maxY)
	seen := make([][]bool, 0, maxY)

	for i := 0; i < maxY; i++ {
		cd := make([]int, 0, maxX)
		cs := make([]bool, 0, maxX)
		for j := 0; j < maxX; j++ {
			cd = append(cd, math.MaxInt)
			cs = append(cs, false)
		}
		dist = append(dist, cd)
		seen = append(seen, cs)
	}

	h := make(pQueue, 1, 100)
	h[0] = node{point{0, 0}, 0}

	for {
		cur := heap.Pop(&h).(node)
		seen[cur.p.y][cur.p.x] = true
		if cur.p.x == maxX-1 && cur.p.y == maxY-1 {
			return cur.val
		}

		for _, near := range [][2]int{
			{cur.p.x + 1, cur.p.y}, {cur.p.x - 1, cur.p.y},
			{cur.p.x, cur.p.y + 1}, {cur.p.x, cur.p.y - 1}} {
			if !(near[0] >= 0 && near[0] < maxX && near[1] >= 0 && near[1] < maxY) ||
				seen[near[1]][near[0]] {
				continue
			}
			val := cur.val + (*cave)[near[1]][near[0]]
			if val >= dist[near[1]][near[0]] {
				continue
			}
			dist[near[1]][near[0]] = val
			heap.Push(&h, node{point{near[0], near[1]}, val})
		}
	}
}

func increase(line *[]int, n int) *[]int {

	new := []int{}

	for _, v := range *line {
		new = append(new, v+n)
		if new[len(new)-1] > 9 {
			new[len(new)-1] %= 9
		}
	}
	return &new
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0
	cave := [][]int{}
	for _, line := range lines {
		cave = append(cave, pkg.ToIntSliceCharacter(line))
	}
	newCave := make([][]int, 0, len(cave)*5)

	for _, row := range cave {
		newLine := []int{}
		for i := 0; i < 5; i++ {
			newLine = append(newLine, *increase(&row, i)...)
		}
		newCave = append(newCave, newLine)
	}
	newnewCave := make([][]int, 0, len(cave)*5)
	for j := 0; j < 5; j++ {
		for _, row := range newCave {
			inc := increase(&row, j)
			newnewCave = append(newnewCave, *inc)
		}
	}

	part1 = lowestPath(&cave)
	part2 = lowestPath(&newnewCave)

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
