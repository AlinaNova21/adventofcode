package day05

import (
	"fmt"

	"github.com/ags131/adventofcode/2019/aoc"
	"github.com/ags131/adventofcode/2019/aoc/intcode"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	var v int
	vals := make([]int, 0)
	for {
		n, _ := fmt.Fscanf(input, "%d", &v)
		if n == 0 {
			break
		}
		vals = append(vals, v)
	}
	clone := func() []int {
		ret := make([]int, len(vals))
		for i, v := range vals {
			ret[i] = v
		}
		return ret
	}
	part1 := 0
	part2 := 0
	m1 := intcode.Machine{
		Memory: clone(),
		Input: func() int {
			return 1
		},
		Output: func(v int) {
			part1 = v
		},
	}
	m1.Run()
	m2 := intcode.Machine{
		Memory: clone(),
		Input: func() int {
			return 5
		},
		Output: func(v int) {
			part2 = v
		},
	}
	m2.Run()
	return aoc.Output{Part1: part1, Part2: part2}
}
