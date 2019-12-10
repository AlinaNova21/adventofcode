package day10

import (
	"bufio"
	"fmt"
	"math"
	"sort"

	"github.com/ags131/adventofcode/2019/aoc"
)

type Grid struct {
	Data          []byte
	Width, Height int
}

func ReadInput(input *aoc.Input) *Grid {
	b := make([]byte, 0)
	s := bufio.NewScanner(input)
	width, height := 0, 0
	for s.Scan() {
		bytes := s.Bytes()
		width = len(bytes)
		height++
		b = append(b, bytes...)
	}
	return &Grid{b, width, height}
}

func (g *Grid) ToIndex(x, y int) int {
	return (y * g.Width) + x
}

func (g *Grid) FromIndex(i int) (int, int) {
	return i % g.Width, int(math.Floor(float64(i) / float64(g.Width)))
}

func (g *Grid) Get(x, y int) (byte, error) {
	ind := g.ToIndex(x, y)
	if ind >= len(g.Data) {
		return 0, fmt.Errorf("Index %d >= %d", ind, len(g.Data))
	}
	return g.Data[ind], nil
}

func (g *Grid) Set(x, y int, v byte) error {
	ind := g.ToIndex(x, y)
	if ind >= len(g.Data) {
		return fmt.Errorf("Index %d >= %d", ind, len(g.Data))
	}
	g.Data[ind] = v
	return nil
}

type Ray struct {
	Angle, Dist    float64
	Source, Target int
}

func (g *Grid) NewRay(src, dst int) Ray {
	// y = mx+b
	// y - b = mx
	// (y - b) / x = m
	x1, y1 := g.FromIndex(src)
	x2, y2 := g.FromIndex(dst)
	ang := math.Atan2(float64(y2-y1), float64(x2-x1))
	dist := math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1))
	// ang -= math.Pi / 2
	ang = ang * (180 / math.Pi)
	ang -= 90 + 180
	for ang < 0 {
		ang += 360
	}
	ang = math.Round(ang*100) / 100
	return Ray{ang, dist, src, dst}
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	grid := ReadInput(input)
	asteroids := make([]int, 0)
	for i, v := range grid.Data {
		if v == '#' {
			asteroids = append(asteroids, i)
		}
	}
	maxHits := 0.0
	srcAst := 0
	for _, a := range asteroids {
		rays := make([]Ray, 0)
		for _, b := range asteroids {
			if a == b {
				continue
			}
			rays = append(rays, grid.NewRay(a, b))
		}
		angles := map[float64]int{}
		for _, ra := range rays {
			angles[ra.Angle] = ra.Target
		}
		hits := float64(len(angles))
		x, y := grid.FromIndex(a)
		if x == 11 && y == 13 {
			fmt.Println(hits, maxHits, a, x, y)
		}
		if hits > maxHits {
			maxHits = hits
			srcAst = a
			fmt.Printf("SrcAst: %d,%d\n", x, y)
		}
	}
	part1 := int(maxHits)
	fmt.Println("part1 done")
	ast := 0
	part2 := 0
	destroyed := map[int]struct{}{}
	// srcAst = grid.ToIndex(8, 3)
	for ast < 200 {
		fmt.Println(ast)
		rays := make([]Ray, 0)
		for _, b := range asteroids {
			if srcAst == b {
				continue
			}
			if _, ok := destroyed[b]; ok {
				continue
			}
			rays = append(rays, grid.NewRay(srcAst, b))
		}
		if len(rays) == 0 {
			break
		}
		angles := map[float64]Ray{}
		for _, ra := range rays {
			ang := ra.Angle
			if r, ok := angles[ang]; ok {
				if ra.Dist < r.Dist {
					angles[ang] = ra
				}
			} else {
				angles[ang] = ra
			}
		}
		angs := make([]float64, 0)
		for ang := range angles {
			angs = append(angs, ang)
		}
		sort.Slice(angs, func(i, j int) bool {
			return angs[i] < angs[j]
		})
		fmt.Println(angs)
		for _, ang := range angs {
			destroyed[angles[ang].Target] = struct{}{}
			ast++
			x, y := grid.FromIndex(angles[ang].Target)
			fmt.Printf("Boom #%d %d,%d %.3f %d\n", ast, x, y, ang, angles[ang].Target)
			if ast == 200 {
				part2 = (x * 100) + y
				break
			}
		}
	}

	return aoc.Output{Part1: part1, Part2: part2}
}
