package day01

import (
	"fmt"
	"math"

	"github.com/ags131/adventofcode/2019/aoc"
)

func calc(mass int) int {
	return int(math.Floor(float64(mass)/3) - 2)
}

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
	for _, mass := range vals {
		part1 += calc(mass)
	}
	for _, mass := range vals {
		lastMass := mass
		for {
			fuel := calc(lastMass)
			if fuel <= 0 {
				break
			}
			part2 += fuel
			lastMass = fuel
		}
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
