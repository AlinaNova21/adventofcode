package day06

import (
	"fmt"
	"math"

	"github.com/ags131/adventofcode/2019/aoc"
)

type Orbit struct {
	Name   string
	Orbits *Orbit
	// Satellites []*Orbit
}

func NewOrbit(name string) *Orbit {
	return &Orbit{
		Name: name,
		// Satellites: make([]*Orbit, 0),
	}
}

func (o *Orbit) Count() int {
	if o.Orbits != nil {
		return o.Orbits.Count() + 1
	}
	return 0
}

func (o *Orbit) Chain() []string {
	if o.Orbits != nil {
		return append(o.Orbits.Chain(), o.Name)
	}
	return []string{o.Name}
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	orbits := map[string]*Orbit{}
	orbits["COM"] = NewOrbit("COM")
	for {
		var orbiting, name string
		n, _ := fmt.Fscanf(input, "%3s)%3s\n", &orbiting, &name)
		if n == 0 {
			break
		}
		var o *Orbit
		if oo, ok := orbits[name]; ok {
			o = oo
		} else {
			o = NewOrbit(name)
			orbits[name] = o
		}
		if or, ok := orbits[orbiting]; ok {
			o.Orbits = or
		} else {
			orbits[orbiting] = NewOrbit(orbiting)
			o.Orbits = orbits[orbiting]
		}
	}
	part1 := 0
	part2 := 0
	fmt.Println(orbits)
	for _, o := range orbits {
		part1 += o.Count()
	}
	you := orbits["YOU"].Orbits.Chain()
	san := orbits["SAN"].Orbits.Chain()
	cnt := int(math.Min(float64(len(you)), float64(len(san))))
	for i := 0; i < cnt; i++ {
		fmt.Printf("%s %s %d\n", you[i], san[i], i)
		if you[i] == san[i] {
			continue
		}
		youLen := len(you)
		sanLen := len(san)
		fmt.Println(youLen, sanLen, i)
		part2 = len(you) + len(san) - (i * 2)
		break
	}
	fmt.Println(you)
	fmt.Println(san)
	return aoc.Output{Part1: part1, Part2: part2}
}
