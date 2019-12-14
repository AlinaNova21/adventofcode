package day14

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	"github.com/ags131/adventofcode/2019/aoc"
)

type Chemical struct {
	Qty  int
	Name string
}

type Reaction struct {
	Consume []Chemical
	Produce Chemical
}

func (r *Reaction) React(qtys map[string]int) {
	mul := int(math.Ceil(float64(-qtys[r.Produce.Name]) / float64(r.Produce.Qty)))
	for _, c := range r.Consume {
		qtys[c.Name] -= c.Qty * mul
	}
	qtys[r.Produce.Name] += r.Produce.Qty * mul
}

var reactions []Reaction

func calc(fuel int) int {
	// fmt.Println("Calculating for", fuel)
	qtys := map[string]int{}
	qtys["FUEL"] = -fuel
	hasNeg := true
	for hasNeg {
		hasNeg = false
		for name, cnt := range qtys {
			if cnt < 0 && name != "ORE" {
				hasNeg = true
				for _, r := range reactions {
					if r.Produce.Name == name {
						r.React(qtys)
					}
				}
			}
		}
	}
	return -qtys["ORE"]
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	s := bufio.NewScanner(input)
	reactions = make([]Reaction, 0)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		chems := make([]Chemical, 0)
		r := strings.NewReader(s.Text())
		for {
			chem := Chemical{}
			n, _ := fmt.Fscanf(r, "%d %s", &chem.Qty, &chem.Name)
			if n == 0 {
				fmt.Fscanf(r, "=>")
				fmt.Fscanf(r, "%d %s", &chem.Qty, &chem.Name)
			}
			chem.Name = strings.TrimRight(chem.Name, ",")
			chems = append(chems, chem)
			if n == 0 {
				break
			}
		}
		reaction := Reaction{chems[:len(chems)-1], chems[len(chems)-1]}
		reactions = append(reactions, reaction)
	}
	part1 := calc(1)
	max := int(1e12)
	low := 0
	high := max
	for low < high {
		mid := (low + high + 1) / 2
		ore := calc(mid)
		// ore := mid  part1
		// fmt.Println("Checking", mid, ore)
		if ore <= max {
			low = mid
		} else {
			high = mid - 1
		}
	}
	part2 := low
	return aoc.Output{Part1: part1, Part2: part2}
}
