package assembler

import (
	"strings"
	"testing"

	"github.com/ags131/adventofcode/2019/aoc/intcode"
)

func TestAssemble(t *testing.T) {
	var err error
	r := strings.NewReader(teststr)
	p, err := Assemble(r)
	if err != nil {
		t.Error(err)
	}
	m := intcode.NewMachine(p)
	for range m.Output {
	}
}

const teststr = `
start: REL 1
OUT %-1
ADD @loop 1 @loop
EQ @loop 16 @cmp
ADD 0 -1 @test
JZ @cmp start
HALT
`

// 109,1,
// 204,-1,
// 1001,100,1,100,
// 1008,100,16,101,
// 1006,101,0,
// 99

// const teststr = `
// # Test Program
// ADD test :6 test
// OUT %test
// HALT
// `
