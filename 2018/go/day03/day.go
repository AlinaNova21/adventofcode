package day03

import (
	"fmt"

	"github.com/ags131/adventofcode/2018/go/aoc"
)

type claim struct {
	ID      int
	X       int
	Y       int
	Width   int
	Height  int
	Overlap bool
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	f := "#%d @ %d,%d: %dx%d"
	claims := make(map[int]claim, 0)
	for {
		c := claim{}
		n, _ := fmt.Fscanf(input, f, &c.ID, &c.X, &c.Y, &c.Width, &c.Height)
		if n == 0 {
			break
		}
		claims[c.ID] = c
	}
	cnts := make(map[int]int)
	for ci, c := range claims {
		for x := c.X; x < c.X+c.Width; x++ {
			for y := c.Y; y < c.Y+c.Height; y++ {
				ind := x + (y * 1000)
				if _, ok := cnts[ind]; ok {
					cnts[ind]++
					claims[ci] = c
				} else {
					cnts[ind] = 1
				}
			}
		}
	}
	part1 := 0
	for _, cnt := range cnts {
		if cnt > 1 {
			part1++
		}
	}

outer:
	for ci, c := range claims {
		for x := c.X; x < c.X+c.Width; x++ {
			for y := c.Y; y < c.Y+c.Height; y++ {
				ind := x + (y * 1000)
				if cnts[ind] > 1 {
					c.Overlap = true
					claims[ci] = c
					continue outer
				}
			}
		}
	}
	part2 := 0
	for _, c := range claims {
		if !c.Overlap {
			part2 = c.ID
		}
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
