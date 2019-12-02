package day02

import (
	"fmt"

	"github.com/ags131/adventofcode/2019/aoc"
)

type Machine struct {
	IP int
	Memory []int
}

func (m *Machine) Step() bool {
	op := m.Memory[m.IP]
	switch op {
	case 1:
		m.Memory[m.Memory[m.IP + 3]] = m.Memory[m.Memory[m.IP + 1]] + m.Memory[m.Memory[m.IP + 2]]
		m.IP += 4
	case 2:
		m.Memory[m.Memory[m.IP + 3]] = m.Memory[m.Memory[m.IP + 1]] * m.Memory[m.Memory[m.IP + 2]]
		m.IP += 4
	case 99:
		return false
	}
	return true
}

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
	m := Machine{
		Memory: p1vals,
	}
	
	// vals = []int{1,9,10,3,2,3,11,0,99,30,40,50}
	for {
		if !m.Step() {
			break
		}
	}
	part1 := m.Memory[0]
	part2 := 0
outer:
	for v := 0; v < 100; v++ {
		for n := 0; n < 100; n++ {
			p2vals := clone()
			p2vals[1] = n
			p2vals[2] = v
			m := Machine{
				Memory: p2vals,
			}
			for {
				if !m.Step() {
					break
				}
			}
			if m.Memory[0] == 19690720 {
				part2 = 100 * n + v
				break outer
			}
		}
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
