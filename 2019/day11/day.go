package day11

import (
	"fmt"
	"math"

	"github.com/ags131/adventofcode/2019/aoc"
	"github.com/ags131/adventofcode/2019/aoc/grid"
	"github.com/ags131/adventofcode/2019/aoc/intcode"
	"github.com/ags131/adventofcode/2019/aoc/point"
)

func GridString(g *grid.Grid) string {
	minX, minY := 1.0e9, 1.0e9
	maxX, maxY := 0.0, 0.0
	g.Do(func(p point.Point2D, v interface{}) {
		if v.(int) == 1 {
			minX = math.Min(minX, float64(p.X))
			minY = math.Min(minY, float64(p.Y))
			maxX = math.Max(maxX, float64(p.X))
			maxY = math.Max(maxY, float64(p.Y))
		}
	})
	out := "\n"
	for y := int(maxY); y >= int(minY); y-- {
		for x := int(minX); x <= int(maxX); x++ {
			v := g.Get(point.New2D(x, y))
			clr := 40 + (7 * v.(int))
			out += fmt.Sprintf("\x1b[%dm ", clr)
		}
		out += fmt.Sprint("\x1b[0m\n")
	}
	return out
}

func initGrid(g *grid.Grid) {
	g.Do(func(p point.Point2D, _ interface{}) {
		g.Set(p, 0)
	})
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	program := intcode.ReadInput(input)
	grids := make([]*grid.Grid, 2)
	grids[0] = grid.New(500, 500)
	grids[1] = grid.New(500, 500)
	initGrid(grids[0])
	initGrid(grids[1])
	touched := make(map[int]struct{}, 0)
	for i := 0; i < 2; i++ {
		g := grids[i]
		pos := point.New2D(g.Width()/2, g.Height()/2)
		dir := 0
		g.Set(pos, i)
		m := intcode.NewMachine(program)
		for !m.Stopped {
			m.Input <- g.Get(pos).(int)
			clr := <-m.Output
			g.Set(pos, clr)
			if i == 0 {
				touched[pos.ID()] = struct{}{}
			}
			switch <-m.Output {
			case 0:
				dir--
			case 1:
				dir++
			}
			if dir < 4 {
				dir += 4
			}
			dir %= 4
			switch dir {
			case 0:
				pos.Y++
			case 1:
				pos.X++
			case 2:
				pos.Y--
			case 3:
				pos.X--
			}
		}
	}

	part1 := len(touched)
	part2 := GridString(grids[1])
	return aoc.Output{Part1: part1, Part2: part2}
}
