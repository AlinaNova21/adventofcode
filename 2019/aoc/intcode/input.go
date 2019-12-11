package intcode

import (
	"strconv"
	"strings"

	"github.com/ags131/adventofcode/2019/aoc"
)

// Program is an intcode program
type Program []int

// ReadInput parses IntCode program input
func ReadInput(input *aoc.Input) Program {
	return aoc.ReadIntSlice(input)
}

// Clone copies the Program
func (p Program) Clone() Program {
	return aoc.CloneIntSlice(p)
}

// String returns the Program as a string
func (p Program) String() string {
	tmp := make([]string, len(p))
	for i, v := range p {
		tmp[i] = strconv.Itoa(v)
	}
	return strings.Join(tmp, ",")
}
