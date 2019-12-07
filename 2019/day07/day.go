package day07

import (
	"github.com/ags131/adventofcode/2019/aoc"
	"github.com/ags131/adventofcode/2019/aoc/intcode"
	prmt "github.com/gitchander/permutation"
)

var program intcode.Program

func CreateAmp(phase int) *intcode.Machine {
	m := intcode.NewMachine(program)
	m.InCh <- phase
	return m
}

func RunAmp(phase, input int) int {
	m := CreateAmp(phase)
	m.InCh <- input
	return <-m.OutCh
}

func RunAmps(phase []int) int {
	chCopy := func(chOut, chIn chan int) {
		for v := range chOut {
			chIn <- v
		}
	}
	mA := CreateAmp(phase[0])
	mB := CreateAmp(phase[1])
	mC := CreateAmp(phase[2])
	mD := CreateAmp(phase[3])
	mE := CreateAmp(phase[4])
	go chCopy(mA.OutCh, mB.InCh)
	go chCopy(mB.OutCh, mC.InCh)
	go chCopy(mC.OutCh, mD.InCh)
	go chCopy(mD.OutCh, mE.InCh)
	mA.InCh <- 0
	chCopy(mE.OutCh, mA.InCh)
	return mE.Result
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
