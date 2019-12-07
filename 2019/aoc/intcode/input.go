package intcode

import (
	"fmt"

	"github.com/ags131/adventofcode/2019/aoc"
)

// Program is an intcode program
type Program []int

// ReadInput parses IntCode program input
func ReadInput(input *aoc.Input) Program {
	ret := make(Program, 0)
	var v int
	for {
		n, _ := fmt.Fscanf(input, "%d", &v)
		if n == 0 {
			break
		}
		ret = append(ret, v)
	}
	return ret
}

func (p Program) Clone() Program {
	ret := make(Program, len(p))
	for i, v := range p {
		ret[i] = v
	}
	return ret
}
