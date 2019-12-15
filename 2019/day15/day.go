package day15

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
		if v.(int) != 0 {
			minX = math.Min(minX, float64(p.X)-1)
			minY = math.Min(minY, float64(p.Y)-1)
			maxX = math.Max(maxX, float64(p.X)+1)
			maxY = math.Max(maxY, float64(p.Y)+1)
		}
	})
	out := "\x1b[2;0H"
	for y := int(minY); y <= int(maxY); y++ {
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
				clr = 40
				ch = "O"
			case 3:
				clr = 45
				ch = "X"
			case 4:
				clr = 41
				ch = " "
			case 5:
				clr = 44
				ch = " "
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
	m := intcode.NewMachine(prog)
	g := grid.New(1000, 1000)
	g.Do(func(p point.Point2D, v interface{}) {
		g.Set(p, 0)
	})
	dg := grid.New(1000, 1000)
	dg.Do(func(p point.Point2D, v interface{}) {
		dg.Set(p, 1.0e18)
	})

	pos := point.New2D(500, 500)
	dir := 1
	iter := 0
	nextDirRight := map[int]int{
		1: 4,
		2: 3,
		3: 1,
		4: 2,
	}
	nextDirLeft := map[int]int{
		4: 1,
		3: 2,
		1: 3,
		2: 4,
	}
	var res int
	var location point.Point2D
	// for location.X == 0 && location.Y == 0 {
	dist := 0.0
	for {
		// var dummy int
		// fmt.Scanf("%d", &dummy)
		// _ = dummy
		m.Input <- dir
		res = <-m.Output
		np := pos
		switch dir {
		case 1:
			np.Y--
		case 2:
			np.Y++
		case 3:
			np.X--
		case 4:
			np.X++
		}
		switch res {
		case 0:
			g.Set(np, 1)
			dir = nextDirRight[dir]
		case 1:
			dg.Set(pos, dist)
			dist++
			dist = math.Min(dist, dg.Get(np).(float64))
			pos = np
			g.Set(pos, 4)
			dir = nextDirLeft[dir]
		case 2:
			dg.Set(pos, dist)
			dist++
			dist = math.Min(dist, dg.Get(np).(float64))
			pos = np
			location = pos
			g.Set(pos, 3)
			// break
		}
		if iter > 0 && pos.X == 500 && pos.Y == 500 {
			break
		}
		iter++
		fmt.Printf("\x1b[H\x1b[Kdir=%d iter=%d pos=%v res=%d\n", dir, iter, pos, res)
		ov := g.Get(pos)
		g.Set(pos, 2)
		fmt.Println(GridString(g))
		g.Set(pos, ov)
	}
	fmt.Println(GridString(g))
	_ = location
	part1 := dg.Get(location).(float64)

	cg := g
	cg.Set(location, 5)
	hasAir := true
	steps := 0
	for hasAir {
		hasAir = false
		ng := grid.New(1000, 1000)
		cg.Do(func(p point.Point2D, val interface{}) {
			v := val.(int)
			if v == 4 {
				hasAir = true
				neighbors := []point.Point2D{Top(p), Bottom(p), Left(p), Right(p)}
				for _, n := range neighbors {
					if cg.Get(n).(int) == 5 {
						v = 5
					}
				}
			}
			ng.Set(p, v)
		})
		steps++
		cg = ng
		fmt.Printf("\x1b[H\x1b[Kiter=%d\n", steps)
		fmt.Println(GridString(cg))
	}

	part2 := steps
	return aoc.Output{Part1: part1, Part2: part2}
}

func Left(p point.Point2D) point.Point2D {
	return point.New2D(p.X-1, p.Y)
}
func Right(p point.Point2D) point.Point2D {
	return point.New2D(p.X+1, p.Y)
}
func Top(p point.Point2D) point.Point2D {
	return point.New2D(p.X, p.Y+1)
}
func Bottom(p point.Point2D) point.Point2D {
	return point.New2D(p.X, p.Y-1)
}
