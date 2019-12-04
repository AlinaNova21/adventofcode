package day02

import (
	"fmt"
	"math"

	"github.com/ags131/adventofcode/2015/aoc"
)

type Box struct {
	L, W, H float64
}

func (b Box) Area() float64 {
	return (2 * b.L * b.W) + (2 * b.W * b.H) + (2 * b.H * b.L)
}

func (b Box) SmallestArea() float64 {
	return math.Min(b.L*b.W, math.Min(b.W*b.H, b.H*b.L))
}

func (b Box) SmallestPerimeter() float64 {
	return math.Min(2*(b.L+b.W), math.Min(2*(b.W+b.H), 2*(b.H+b.L)))
}

func (b Box) Volume() float64 {
	return b.L * b.W * b.H
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	part1 := 0.0
	part2 := 0.0
	for {
		box := Box{}
		n, _ := fmt.Fscanf(input, "%fx%fx%f", &box.L, &box.W, &box.H)
		if n == 0 {
			break
		}
		part1 += box.Area() + box.SmallestArea()
		part2 += box.SmallestPerimeter() + box.Volume()
	}
	return aoc.Output{Part1: int(part1), Part2: int(part2)}
}
