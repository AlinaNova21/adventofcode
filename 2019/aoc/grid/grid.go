package grid

import (
	"math"

	"github.com/ags131/adventofcode/2019/aoc/point"
)

type Grid struct {
	width, height int
	data          []interface{}
}

func New(w, h int) *Grid {
	return &Grid{w, h, make([]interface{}, w*h)}
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) Height() int {
	return g.height
}

func (g *Grid) Len() int {
	return g.width * g.height
}

func (g *Grid) toIndex(p point.Point2D) int {
	return (p.Y * g.width) + p.X
}

func (g *Grid) fromIndex(i int) point.Point2D {
	return point.New2D(i%g.width, int(math.Floor(float64(i)/float64(g.width))))
}

func (g *Grid) Get(p point.Point2D) interface{} {
	ind := g.toIndex(p)
	if ind >= len(g.data) {
		return nil
	}
	return g.data[ind]
}

func (g *Grid) Set(p point.Point2D, v interface{}) {
	ind := g.toIndex(p)
	g.data[ind] = v
}

func (g *Grid) Do(fn func(point.Point2D, interface{})) {
	for y := 0; y < g.width; y++ {
		for x := 0; x < g.width; x++ {
			p := point.New2D(x, y)
			fn(p, g.Get(p))
		}
	}
}
