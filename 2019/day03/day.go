package day03

import (
	"fmt"
	"math"
	"strings"

	"github.com/ags131/adventofcode/2019/aoc"
)

type Point struct {
	X int
	Y int
}

func (p Point) Dist() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

type Step struct {
	Dir  string
	Dist int
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	var v string
	lines := make([][]Step, 0)
	for {
		n, _ := fmt.Fscan(input, &v)
		if n == 0 {
			break
		}
		steps := make([]Step, 0)
		reader := strings.NewReader(v)
		for {
			step := Step{}
			n, _ := fmt.Fscanf(reader, "%1s%d", &step.Dir, &step.Dist)
			if n == 0 {
				break
			}
			steps = append(steps, step)
		}
		lines = append(lines, steps)
	}
	lines2 := make([][]Point, len(lines))
	for i, line := range lines {
		lines2[i] = make([]Point, 0)
		x := 0
		y := 0
		for _, step := range line {
			for j := 0; j < step.Dist; j++ {
				lines2[i] = append(lines2[i], Point{x, y})
				switch step.Dir {
				case "U":
					y--
				case "D":
					y++
				case "L":
					x--
				case "R":
					x++
				}
			}
		}
	}
	closest1 := 1.000e18
	space1 := make(map[int]int, 0)
	space2 := make(map[int]int, 0)
	dist := 0
	intersecting := make(map[int]interface{})
	for _, p1 := range lines2[0] {
		space1[1e8*p1.X+p1.Y] = dist
		dist++
	}
	dist = 0
	for _, p := range lines2[1] {
		i := 1e8*p.X + p.Y
		space2[i] = dist
		if _, ok := space1[i]; ok {
			intersecting[i] = nil
		}
		dist++
	}
	closest2 := 1.000e18
	delete(intersecting, 0)
	for i := range intersecting {
		fi := float64(i)
		nd := math.Abs(math.Floor(fi/1e8)) + math.Abs(math.Remainder(fi, 1e8))
		closest1 = math.Min(nd, closest1)
		nv := float64(space1[i] + space2[i])
		closest2 = math.Min(nv, closest2)
	}
	// part1 := closest.Dist()
	part1 := closest1
	part2 := closest2
	return aoc.Output{Part1: part1, Part2: part2}
}
