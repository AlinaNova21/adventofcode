package day01

import (
	"io/ioutil"

	"github.com/ags131/adventofcode/2015/aoc"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	inp, _ := ioutil.ReadAll(input)

	part1 := 0
	part2 := 0
	for i, r := range string(inp) {
		if r == '(' {
			part1++
		} else {
			part1--
		}
		if part1 == -1 && part2 == 0 {
			part2 = i + 1
		}
	}

	return aoc.Output{Part1: part1, Part2: part2}
}
