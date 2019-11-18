package day09

import (
	"container/ring"
	"fmt"

	"github.com/ags131/adventofcode/2018/go/aoc"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	players := 10
	last := 1618
	f := "%d players; last marble is worth %d points"
	f1 := f + "\n"
	f2 := f + ": high score is %d\n"
	fmt.Fscanf(input, f, &players, &last)
	fmt.Printf(f1, players, last)

	part1 := runMarbles(players, last)
	fmt.Printf(f2, players, last, part1)
	part2 := runMarbles(players, last*100)
	fmt.Printf(f2, players, last*100, part2)
	return aoc.Output{Part1: part1, Part2: part2}
}

func runMarbles(players, last int) int {
	mkMarble := func(v int) *ring.Ring {
		r := ring.New(1)
		r.Value = v
		return r
	}

	next := 1
	current := mkMarble(0)
	highScore := 0
	scores := make(map[int]int, players)

	player := 0
	for i := 0; i < last; i++ {
		if next%23 == 0 {
			for m := 0; m < 8; m++ {
				current = current.Prev()
			}
			pop := current.Unlink(1).Value.(int)
			scores[player] += next + pop
			if scores[player] > highScore {
				highScore = scores[player]
			}
			current = current.Next()
		} else {
			nm := mkMarble(next)
			current.Next().Link(nm)
			current = nm
		}
		next++
		player++
		if player >= players {
			player = 0
		}
		if last < 50 {
			marbleDisplay(player, current)
		}
	}
	return highScore
}

func marbleDisplay(p int, r *ring.Ring) {
	fmt.Printf("[%d] ", p)
	r.Do(func(p interface{}) {
		fmt.Printf("%d ", p.(int))
	})
	fmt.Println()
}
