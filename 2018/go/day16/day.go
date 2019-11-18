package day16

import (
	"fmt"

	"github.com/ags131/adventofcode/2018/go/aoc"
	"github.com/ags131/adventofcode/2018/go/aoc/elfcode"
)

type testCase struct {
	Before      [4]int
	After       [4]int
	Instruction [4]int
	ValidOps    []int
}

func newTestCase() *testCase {
	return &testCase{
		ValidOps: make([]int, 0),
	}
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	testCases := make([]*testCase, 0)
	for {
		tc := newTestCase()
		n, _ := fmt.Fscanf(input, "Before: [%d, %d, %d, %d]", &tc.Before[0], &tc.Before[1], &tc.Before[2], &tc.Before[3])
		if n == 0 {
			fmt.Fscan(input)
			fmt.Fscan(input)
			fmt.Fscan(input)
			break
		}
		n, _ = fmt.Fscan(input, &tc.Instruction[0], &tc.Instruction[1], &tc.Instruction[2], &tc.Instruction[3])
		n, _ = fmt.Fscanf(input, "After:  [%d, %d, %d, %d]", &tc.After[0], &tc.After[1], &tc.After[2], &tc.After[3])
		fmt.Fscanln(input)
		fmt.Fscanln(input)
		testCases = append(testCases, tc)
	}
	program := make([][4]int, 0)
	for {
		i := [4]int{}
		n, _ := fmt.Fscan(input, &i[0], &i[1], &i[2], &i[3])
		if n == 0 {
			break
		}
		program = append(program, i)
	}
	opNames := []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "seti", "setr", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"}
	m := elfcode.NewMachine()

	for _, tc := range testCases {
		for op := range opNames {
			m.Reset()
			copy(m.Register[:], tc.Before[:])
			// fmt.Println(m.String())
			ins := [4]int{}
			copy(ins[:], tc.Instruction[:])
			ins[0] = op
			m.Memory = [][4]int{
				ins,
			}
			m.Step()

			// fmt.Printf("[%d, %d, %d, %d] %s %v\n", tc.After[0], tc.After[1], tc.After[2], tc.After[3], op, valid)
			if regMatch(m.Register, tc.After) {
				tc.ValidOps = append(tc.ValidOps, op)
			}
		}
	}
	part1 := 0
	for _, tc := range testCases {
		if len(tc.ValidOps) >= 3 {
			part1++
		}
	}

	opmap := make([]int, elfcode.OP_COUNT)
	filter := make(map[int]struct{}, 0)
	for {
		if len(filter) == elfcode.OP_COUNT {
			break
		}
		for _, tc := range testCases {
			validOps := make([]int, 0)
			for _, op := range tc.ValidOps {
				if _, ok := filter[op]; !ok {
					validOps = append(validOps, op)
				}
			}
			if len(validOps) == 1 {
				op := validOps[0]
				filter[op] = struct{}{}
				opmap[tc.Instruction[0]] = op
			}
		}
	}
	m.Reset()
	m.OpMap = opmap
	m.Memory = program
	for {
		if err := m.Step(); err != nil {
			break
		}
	}
	part2 := m.Register[0]
	return aoc.Output{Part1: part1, Part2: part2}
}

func regMatch(a, b [4]int) bool {
	for i := 0; i < 4; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
