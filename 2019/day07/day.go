package day07

import (
	"github.com/ags131/adventofcode/2019/aoc"
	"github.com/ags131/adventofcode/2019/aoc/intcode"
	prmt "github.com/gitchander/permutation"
)

var program intcode.Program

func CreateAmp(phase int) (chan int, chan int) {
	m := intcode.NewMachine(program)
	chIn, chOut := m.RunCh()
	chIn <- phase
	return chIn, chOut
}

func RunAmp(phase, input int) int {
	chIn, chOut := CreateAmp(phase)
	chIn <- input
	return <-chOut
}

func RunAmps(phase []int) int {
	chCopy := func(chOut, chIn chan int) {
		for v := range chOut {
			chIn <- v
		}
	}
	chInA, chOutA := CreateAmp(phase[0])
	chInB, chOutB := CreateAmp(phase[1])
	chInC, chOutC := CreateAmp(phase[2])
	chInD, chOutD := CreateAmp(phase[3])
	chInE, chOutE := CreateAmp(phase[4])
	go chCopy(chOutA, chInB)
	go chCopy(chOutB, chInC)
	go chCopy(chOutC, chInD)
	go chCopy(chOutD, chInE)
	chInA <- 0
	chCopy(chOutE, chInA)
	return <-chInA
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	program = intcode.ReadInput(input)

	part1 := 0
	part2 := 0

	phase1slice := []int{0, 1, 2, 3, 4}
	prmt1 := prmt.New(prmt.IntSlice(phase1slice))
	for prmt1.Next() {
		v := 0
		for i := 0; i < 5; i++ {
			v = RunAmp(phase1slice[i], v)
		}
		if v > part1 {
			part1 = v
		}
	}
	phase2slice := []int{5, 6, 7, 8, 9}
	prmt2 := prmt.New(prmt.IntSlice(phase2slice))
	for prmt2.Next() {
		v := RunAmps(phase2slice)
		if v > part2 {
			part2 = v
		}
	}

	return aoc.Output{Part1: part1, Part2: part2}
}
