package intcode

import (
	"github.com/ags131/adventofcode/2019/aoc"
)

// Program is an intcode program
type Program []int

// ReadInput parses IntCode program input
func ReadInput(input *aoc.Input) Program {
	return aoc.ReadIntSlice(input)
}

func (p Program) Clone() Program {
	return aoc.CloneIntSlice(p)
}
