package day05

import (
	"bufio"
	"math"
	"strings"
	"sync"

	"github.com/ags131/adventofcode/2018/go/aoc"
)

type polymer struct {
	Units []unit
}

func newPolymer(poly string) *polymer {
	units := make([]unit, len(poly))
	for i, c := range poly {
		neg := c <= 90
		if neg {
			c += 32
		}
		// fmt.Printf("%v %v\n", c, neg)
		units[i] = unit{c, neg}
	}
	return &polymer{units}
}

func (p polymer) Clone() *polymer {
	units := make([]unit, len(p.Units))
	for i, u := range p.Units {
		units[i] = u
	}
	return &polymer{units}
}

func (p polymer) Filter(fn func(unit) bool) *polymer {
	ret := p.Clone()
	ret.Units = make([]unit, 0)
	for _, u := range p.Units {
		if fn(u) {
			ret.Units = append(ret.Units, u)
		}
	}
	return ret
}

func (p *polymer) React() bool {
	for i, a := range p.Units[1:] {
		b := p.Units[i]
		// fmt.Printf("%v %v %v\n", i, a, b)
		if a.Type == b.Type && a.Negative != b.Negative {
			p.Units = append(p.Units[:i], p.Units[i+2:]...)
			return true
		}
	}
	return false
}

func (p *polymer) ReactAll() {
	for {
		if !p.React() {
			break
		}
	}
}

func (p *polymer) String() string {
	str := ""
	for _, u := range p.Units {
		if u.Negative {
			str += strings.ToUpper(string(u.Type))
		} else {
			str += string(u.Type)
		}
	}
	return str
}

type unit struct {
	Type     rune
	Negative bool
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	p := newPolymer(scanner.Text())

	p1 := p.Clone()
	part1 := 0
	p1.ReactAll()
	part1 = len(p1.Units)

	wg := sync.WaitGroup{}
	wg.Add(26)
	ret := make(chan int)
	for l := 'a'; l <= 'z'; l++ {
		go func(p *polymer, l rune, ret chan int) {
			fp := p.Filter(func(u unit) bool {
				return u.Type != l
			})
			fp.ReactAll()
			ret <- len(fp.Units)
			wg.Done()
		}(p, l, ret)
	}
	part2 := math.MaxInt32
	for i := 0; i < 26; i++ {
		l := <-ret
		// fmt.Println(l)
		if l < part2 {
			part2 = l
		}
	}
	wg.Wait()
	return aoc.Output{Part1: part1, Part2: part2}
}
