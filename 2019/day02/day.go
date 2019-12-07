package day02

import (
	"github.com/ags131/adventofcode/2019/aoc"
	"github.com/ags131/adventofcode/2019/aoc/intcode"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	program := intcode.ReadInput(input)

	p1vals := program.Clone()
	p1vals[1] = 12
	p1vals[2] = 2
	m := intcode.NewMachine(p1vals)
	m.Wait()
	part1 := m.Memory[0]
	part2 := 0
outer:
	for v := 0; v < 100; v++ {
		for n := 0; n < 100; n++ {
			p2vals := program.Clone()
			p2vals[1] = n
			p2vals[2] = v
			m := intcode.NewMachine(p2vals)
			m.Wait()
			if m.Memory[0] == 19690720 {
				part2 = 100*n + v
				break outer
			}
		}
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
