package day19

import (
	"fmt"
	"math"
	"sync"

	"github.com/ags131/adventofcode/2019/aoc"
	"github.com/ags131/adventofcode/2019/aoc/grid"
	"github.com/ags131/adventofcode/2019/aoc/intcode"
	"github.com/ags131/adventofcode/2019/aoc/point"
)

func GridString(g *grid.Grid) string {
	minX, minY := 1.0e9, 1.0e9
	maxX, maxY := 0.0, 0.0
	g.Do(func(p point.Point2D, v interface{}) {
		if v.(int) != 0 {
			minX = math.Min(minX, float64(p.X))
			minY = math.Min(minY, float64(p.Y))
			maxX = math.Max(maxX, float64(p.X))
			maxY = math.Max(maxY, float64(p.Y))
		}
	})
	// out := "\x1b[2;0H"
	out := "\n"
	for y := int(minY); y <= int(maxY); y++ {
		for x := int(minX); x <= int(maxX); x++ {
			v := g.Get(point.New2D(x, y))
			clr := 40
			ch := " "
			switch v.(int) {
			case 1:
				clr = 40
				ch = "."
			case 2:
				clr = 47
				ch = "#"
			}
			out += fmt.Sprintf("\x1b[%dm%s", clr, string(ch))
		}
		out += fmt.Sprint("\x1b[0m\n")
	}
	return out
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	prog := intcode.ReadInput(input)

	g := grid.New(50, 50)
	g.Do(func(p point.Point2D, v interface{}) {
		g.Set(p, 0)
	})
	part1 := 0
	wg := sync.WaitGroup{}
	test := func(x, y int) int {
		m := intcode.NewMachine(prog)
		m.Input <- x
		m.Input <- y
		v := <-m.Output
		return v
	}
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			wg.Add(1)
			p := point.New2D(x, y)
			v := test(p.X, p.Y)
			g.Set(p, v+1)
			part1 += v
		}
	}

	y := 900
	part2 := 0
	for {
		xo := 0
		for x := 0; x < y; x += y / 4 {
			if test(x, y) == 1 {
				xo = x
				break
			}
		}
		ls := xo
		le := xo
		for test(ls, y) != 0 {
			ls--
		}
		ls++
		for test(le, y) != 0 {
			le++
		}
		le--
		if le-ls < 100 {
			y += 100
			continue
		}
		xo = le - 99
		if test(xo, y+99) == 1 {
			part2 = (xo)*10000 + y
			break
		}
		y++
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
