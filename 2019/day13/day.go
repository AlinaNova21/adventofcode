package day13

import (
	"fmt"
	"math"
	"time"

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
	out := "\n\x1b[H"
	for y := int(maxY); y >= int(minY); y-- {
		for x := int(minX); x <= int(maxX); x++ {
			v := g.Get(point.New2D(x, y))
			clr := 47
			ch := " "
			switch v.(int) {
			case 0:
				clr = 40
				ch = " "
			case 1:
				clr = 47
				ch = " "
			case 2:
				clr = 46
				ch = "#"
			case 3:
				clr = 46
				ch = "|"
			case 4:
				clr = 46
				ch = "O"
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
	grid := grid.New(100, 100)
	grid.Do(func(p point.Point2D, v interface{}) {
		grid.Set(p, 0)
	})
	prog[0] = 2
	score := 0
	fmt.Print("\x1b[2J")
	m := intcode.NewMachine(prog)
	done := false
	var ball, paddle point.Point2D
	part1 := 0
	go func() {
		for !done {
			x := <-m.Output
			y := <-m.Output
			t := <-m.Output
			if x == -1 && y == 0 {
				score = t
			} else {
				p := point.New2D(x, y)
				grid.Set(p, t)
				if t == 3 {
					paddle = p
				}
				if t == 4 {
					ball = p
				}
			}
		}
	}()
	go func() {
		for !done {
			in := 0
			// fmt.Scanf("%d", &in)
			if ball.X < paddle.X {
				in = -1
			}
			if ball.X > paddle.X {
				in = 1
			}
			m.Input <- in
			if part1 == 0 {
				grid.Do(func(p point.Point2D, v interface{}) {
					if v.(int) == 2 {
						part1++
					}
				})
			}
			fmt.Print()
			fmt.Println(GridString(grid))
			fmt.Println(score)
		}
	}()
	m.Wait()
	time.Sleep(1 * time.Second)
	done = true
	part2 := score
	return aoc.Output{Part1: part1, Part2: part2}
}
