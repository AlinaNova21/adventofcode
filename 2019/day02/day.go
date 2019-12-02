package day02

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
	p1vals := clone()
	p1vals[1] = 12
	p1vals[2] = 2
	m := intcode.Machine{
		Memory: p1vals,
	}
	m.Run()
	part1 := m.Memory[0]
	part2 := 0
outer:
	for v := 0; v < 100; v++ {
		for n := 0; n < 100; n++ {
			p2vals := clone()
			p2vals[1] = n
			p2vals[2] = v
			m := intcode.Machine{
				Memory: p2vals,
			}
			m.Run()
			if m.Memory[0] == 19690720 {
				part2 = 100*n + v
				break outer
			}
		}
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
