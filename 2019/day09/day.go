package day09

import (
	"github.com/ags131/adventofcode/2019/aoc"
	"github.com/ags131/adventofcode/2019/aoc/intcode"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	program := intcode.ReadInput(input)
	m1 := intcode.NewMachine(program)
	m1.Input <- 1
	part1 := <-m1.Output
	m2 := intcode.NewMachine(program)
	m2.Input <- 2
	part2 := <-m2.Output
	return aoc.Output{Part1: part1, Part2: part2}
}
