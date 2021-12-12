package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/meridani/advent-of-code-2021/pkg"
	"github.com/meridani/advent-of-code-2021/pkg/execute"
)

var DAY = 12

type cave struct {
	id          string
	connections map[string]*cave
	big         bool
}

func parseCaves(lines []string) map[string]*cave {
	caves := make(map[string]*cave)
	for _, line := range lines {

		parts := strings.Split(strings.TrimSpace(line), "-")
		lid, rid := parts[0], parts[1]

		if _, lok := caves[lid]; !lok {
			big := strings.ToUpper(lid) == lid
			c := &cave{lid, make(map[string]*cave), big}
			caves[c.id] = c
		}
		if _, rok := caves[rid]; !rok {
			big := strings.ToUpper(rid) == rid
			c := &cave{rid, make(map[string]*cave), big}
			caves[c.id] = c
		}

		lc := caves[lid]
		rc := caves[rid]

		if _, ok := lc.connections[rid]; !ok {
			lc.connections[rid] = rc
		}
		if _, ok := rc.connections[lid]; !ok {
			rc.connections[lid] = lc
		}
	}
	return caves
}

func findPath(caves map[string]*cave, modify bool) *[][]*cave {
	paths := [][]*cave{}
	finalPaths := [][]*cave{}

	start := caves["start"]
	// for _, c := range caves {
	// 	if c.id == "start" {
	// 		start = c
	// 	}
	// }

	for _, c := range start.connections {
		p := []*cave{}
		// p = append(p, start)
		p = append(p, c)
		paths = append(paths, p)
	}

	for len(paths) > 0 {
		current := paths[0]
		paths = paths[1:]

		for _, cn := range current[len(current)-1].connections {
			if cn.id == "start" {
				continue
			}
			if cn.id == "end" {
				fp := make([]*cave, len(current))
				copy(fp, current)
				fp = append(fp, cn)
				finalPaths = append(finalPaths, fp)
				continue
			}
			if !modify {
				if pathContains(cn, current) > 0 && !cn.big {
					continue
				}
			} else {
				mto := moreThanOnce(current)
				if mto.id == cn.id && !cn.big {
					continue
				}
			}
			newPath := make([]*cave, len(current))
			copy(newPath, current)
			newPath = append(newPath, cn)
			paths = append(paths, newPath)
		}
	}

	if false {

		for _, p := range finalPaths {
			for _, c := range p {
				fmt.Print(c.id)
				if c.id != "end" {
					fmt.Print("->")
				}
			}
			fmt.Println()
		}
	}

	return &finalPaths
}

func moreThanOnce(p []*cave) *cave {
	highest := cave{}
	for _, v := range p {
		if cur := pathContains(v, p); cur > 1 && !v.big {
			highest = *v
		}
	}
	return &highest
}

func pathContains(c *cave, p []*cave) int {
	sum := 0
	for _, current := range p {
		if current.id == c.id {
			sum++
		}
	}
	return sum
}

func run(input pkg.Input) (interface{}, interface{}) {

	// numbers := input.AsIntSlice()
	lines := input.AsStringSlice()

	part1 := 0
	part2 := 0

	caves := parseCaves(lines)
	fp := findPath(caves, false)
	part1 = len(*fp)
	fp = nil
	fp = findPath(caves, true)
	part2 = len(*fp)
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
