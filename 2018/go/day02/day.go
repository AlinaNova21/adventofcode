package day02

import (
	"fmt"

	"github.com/ags131/adventofcode/2018/go/aoc"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	var v string
	vals := make([]string, 0)
	for {
		n, _ := fmt.Fscan(input, &v)
		if n == 0 {
			break
		}
		vals = append(vals, v)
	}

	part1 := 0
	part2 := ""

	two := 0
	three := 0
	for _, v := range vals {
		chars := make(map[rune]int)
		for _, c := range v {
			if _, ok := chars[c]; ok {
				chars[c]++
			} else {
				chars[c] = 1
			}
		}
		for _, cnt := range chars {
			if cnt == 2 {
				two++
				break
			}
		}
		for _, cnt := range chars {
			if cnt == 3 {
				three++
				break
			}
		}
	}
	part1 = two * three

outer:
	for _, a := range vals {
		for _, b := range vals {
			diff := 0
			diffStr := ""
			for i := 0; i < len(a); i++ {
				if a[i] != b[i] {
					diff++
				} else {
					diffStr += string(b[i])
				}
			}
			if diff == 1 {
				part2 = diffStr
				break outer
			}
		}
	}

	return aoc.Output{Part1: part1, Part2: part2}
}
