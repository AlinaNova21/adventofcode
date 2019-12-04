package day03

import (
	"io/ioutil"

	"github.com/ags131/adventofcode/2015/aoc"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	inp, _ := ioutil.ReadAll(input)
	part1 := 0
	part2 := 0
	visited1 := map[int]interface{}{}
	visited2 := map[int]interface{}{}
	pos1 := 0
	pos2 := [2]int{}
	visited1[0] = nil
	visited2[0] = nil
	for i, dir := range string(inp) {
		switch dir {
		case '^':
			pos1 += 1e8
			pos2[i%2] += 1e8
		case 'v':
			pos1 -= 1e8
			pos2[i%2] -= 1e8
		case '<':
			pos1++
			pos2[i%2]++
		case '>':
			pos1--
			pos2[i%2]--
		}
		visited1[pos1] = nil
		visited2[pos2[i%2]] = nil
	}
	part1 = len(visited1)
	part2 = len(visited2)
	return aoc.Output{Part1: part1, Part2: part2}
}
