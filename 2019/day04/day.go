package day04

import (
	"fmt"
	"strconv"

	"github.com/ags131/adventofcode/2019/aoc"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	var start, end int
	fmt.Fscanf(input, "%d-%d", &start, &end)
	part1 := 0
	part2 := 0
	for n := start; n <= end; n++ {
		str := strconv.Itoa(n)
		if str[0] <= str[1] && str[1] <= str[2] && str[2] <= str[3] && str[3] <= str[4] && str[4] <= str[5] {
			cnt := map[rune]int{}
			for _, c := range str {
				if v, ok := cnt[c]; ok {
					cnt[c] = v + 1
				} else {
					cnt[c] = 1
				}
			}
			hasTwo := false
			hasMore := false
			for _, n := range cnt {
				if n == 2 {
					hasTwo = true
				}
				if n >= 2 {
					hasMore = true
				}
			}
			if hasMore {
				part1++
				if hasTwo {
					part2++
				}
			}
		}
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
