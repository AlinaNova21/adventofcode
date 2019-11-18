package day04

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/ags131/adventofcode/2018/go/aoc"
)

type guard struct {
	ID       int
	Minutes  []int
	Sleeping int
}

type record struct {
	TS   string
	Text string
}

type recordSlice []record

func (rs recordSlice) Sort() {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].TS < rs[j].TS
	})
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	records := make(recordSlice, 0)
	scanner := bufio.NewScanner(input)
	for {
		r := record{}
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		r.TS = text[1:17]
		r.Text = text[19:]
		records = append(records, r)
	}
	records.Sort()
	id := 0
	asleep := 0
	guards := make(map[int]guard, 0)
	for _, r := range records {
		min, _ := strconv.Atoi(r.TS[14:])
		switch r.Text {
		case "wakes up":
			g := guards[id]
			for m := asleep; m < min; m++ {
				g.Minutes[m]++
			}
			g.Sleeping += min - asleep
			guards[id] = g
			asleep = 0
		case "falls asleep":
			asleep = min
		default:
			fmt.Sscanf(r.Text, "Guard #%d", &id)
			if _, ok := guards[id]; !ok {
				guards[id] = guard{id, make([]int, 60), 0}
			}
			asleep = 0
		}
	}
	sleepiestGuard := guard{}
	sleepiestMinute := 0
	for _, g := range guards {
		if sleepiestGuard.ID == 0 || g.Sleeping > sleepiestGuard.Sleeping {
			sleepiestGuard = g
			for min, cnt := range g.Minutes {
				if g.Minutes[sleepiestMinute] < cnt {
					sleepiestMinute = min
				}
			}
		}
	}
	part1 := sleepiestGuard.ID * sleepiestMinute

	sleepiestGuard = guard{}
	sleepiestMinute = 0
	for _, g := range guards {
		max := 0
		for min, cnt := range g.Minutes {
			if g.Minutes[max] < cnt {
				max = min
			}
		}
		if sleepiestGuard.ID == 0 || g.Minutes[max] > sleepiestGuard.Minutes[sleepiestMinute] {
			sleepiestMinute = max
			sleepiestGuard = g
		}
	}
	part2 := sleepiestGuard.ID * sleepiestMinute
	return aoc.Output{Part1: part1, Part2: part2}
}
