package day01

import (
	"fmt"
	"log"

	"github.com/ags131/adventofcode/2018/go/aoc"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	var v int
	vals := make([]int, 0)
	for {
		n, _ := fmt.Fscan(input, &v)
		if n == 0 {
			break
		}
		vals = append(vals, v)
	}

	part1 := 0
	part2 := 0
	freqSeen := make(map[int]struct{}, 0)
	for _, v := range vals {
		part1 += v
	}
outer:
	for {
		for _, v := range vals {
			part2 += v
			if _, ok := freqSeen[part2]; ok {
				log.Printf("Freq Rep %d", part2)
				break outer
			}
			freqSeen[part2] = struct{}{}
		}
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
