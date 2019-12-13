package day12

import (
	"bufio"
	"fmt"
	"math"

	"github.com/ags131/adventofcode/2019/aoc"
	"gonum.org/v1/gonum/spatial/r3"
)

type Moon struct {
	Pos, Vel r3.Vec
}

func (m *Moon) PotentialEnergy() float64 {
	return math.Abs(m.Pos.X) + math.Abs(m.Pos.Y) + math.Abs(m.Pos.Z)
}
func (m *Moon) KineticEnergy() float64 {
	return math.Abs(m.Vel.X) + math.Abs(m.Vel.Y) + math.Abs(m.Vel.Z)
}
func (m *Moon) TotalEnergy() float64 {
	return m.PotentialEnergy() * m.KineticEnergy()
}

func (m Moon) String() string {
	return fmt.Sprintf("pos=<x=%.0f, y=%.0f, z=%.0f>, vel=<x=%.0f, y=%.0f, z=%.0f>",
		m.Pos.X, m.Pos.Y, m.Pos.Z,
		m.Vel.X, m.Vel.Y, m.Vel.Z,
	)
}

func doStep(moons []*Moon) {
	pairs := [6][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}
	for _, pair := range pairs {
		m1 := moons[pair[0]]
		m2 := moons[pair[1]]
		if m1.Pos.X > m2.Pos.X {
			m1.Vel.X--
			m2.Vel.X++
		} else if m1.Pos.X < m2.Pos.X {
			m1.Vel.X++
			m2.Vel.X--
		}
		if m1.Pos.Y > m2.Pos.Y {
			m1.Vel.Y--
			m2.Vel.Y++
		} else if m1.Pos.Y < m2.Pos.Y {
			m1.Vel.Y++
			m2.Vel.Y--
		}
		if m1.Pos.Z > m2.Pos.Z {
			m1.Vel.Z--
			m2.Vel.Z++
		} else if m1.Pos.Z < m2.Pos.Z {
			m1.Vel.Z++
			m2.Vel.Z--
		}
	}
	for _, m := range moons {
		m.Pos = m.Pos.Add(m.Vel)
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func dbg(step int, moons []*Moon) {
	fmt.Printf("After %d step:\n", step)
	for _, m := range moons {
		fmt.Println(m)
	}
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	moons := make([]*Moon, 0)
	moons2 := make([]*Moon, 0)
	s := bufio.NewScanner(input)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		v := r3.Vec{}
		fmt.Sscanf(s.Text(), "<x=%f, y=%f, z=%f>", &v.X, &v.Y, &v.Z)
		moons = append(moons, &Moon{v, r3.Vec{}})
		moons2 = append(moons2, &Moon{v.Add(r3.Vec{}), r3.Vec{}})
	}

	// dbg(0, moons)
	for step := 0; step < 10; step++ {
		doStep(moons)
		// dbg(step + 1, moons)
	}
	part1 := 0
	for _, m := range moons {
		part1 += int(m.TotalEnergy())
	}
	seenX, seenY, seenZ := map[string]struct{}{}, map[string]struct{}{}, map[string]struct{}{}
	repeatX, repeatY, repeatZ := 0, 0, 0
	for i := 0; i < 1e8; i++ {
		if repeatX != 0 && repeatY != 0 && repeatZ != 0 {
			break
		}
		doStep(moons2)
		if i == 0 {
			// dbg(i, moons2)
		}
		if repeatX == 0 {
			key := ""
			for _, m := range moons2 {
				key += fmt.Sprintf("%.0f %.0f ", m.Pos.X, m.Vel.X)
			}
			if _, ok := seenX[key]; ok {
				repeatX = i
				// fmt.Printf("Match X: %d %s\n", i, key)
			}
			seenX[key] = struct{}{}
		}
		if repeatY == 0 {
			key := ""
			for _, m := range moons2 {
				key += fmt.Sprintf("%.0f %.0f ", m.Pos.Y, m.Vel.Y)
			}
			if _, ok := seenY[key]; ok {
				repeatY = i
				// fmt.Printf("Match Y: %d %s\n", i, key)
			}
			seenY[key] = struct{}{}
		}
		if repeatZ == 0 {
			key := ""
			for _, m := range moons2 {
				key += fmt.Sprintf("%.0f %.0f ", m.Pos.Z, m.Vel.Z)
			}
			if _, ok := seenZ[key]; ok {
				repeatZ = i
				// fmt.Printf("Match Z: %d %s\n", i, key)
			}
			seenZ[key] = struct{}{}
		}
	}
	part2 := LCM(repeatX, repeatY, repeatZ)
	return aoc.Output{Part1: part1, Part2: part2}
}
