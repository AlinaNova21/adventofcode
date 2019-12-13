package assembler

import (
	"fmt"
	"math"
	"strings"

	"github.com/ags131/adventofcode/2019/aoc/intcode"
)

func getMode(op int, param int) intcode.ParameterMode {
	t1 := int(math.Pow(10, float64(param+2)))
	t2 := int(math.Pow(10, float64(param+1)))
	return intcode.ParameterMode((op % t1) / t2)
}

func Disassemble(p intcode.Program) []string {
	ret := make([]string, 0)
	revOpMap := map[int]string{}
	for k, v := range opMap {
		revOpMap[int(v)] = k
	}
	paramMap := map[intcode.OPCode]int{
		intcode.OPCodeAdd:         3,
		intcode.OPCodeMul:         3,
		intcode.OPCodeInput:       1,
		intcode.OPCodeOutput:      1,
		intcode.OPCodeJumpIfTrue:  2,
		intcode.OPCodeJumpIfFalse: 2,
		intcode.OPCodeLessThan:    3,
		intcode.OPCodeEquals:      3,
		intcode.OpCodeSetRelBase:  1,
		intcode.OPCodeHalt:        0,
	}
	off := 0
	for off < len(p) {
		op := intcode.OPCode(p[off] % 100)
		paramCnt := paramMap[op]
		params := make([]string, paramCnt)
		for i := 0; i < paramCnt; i++ {
			pm := getMode(p[off], i+1)
			prefix := ""
			switch pm {
			case intcode.ParameterModePosition:
				prefix = "@"
			case intcode.ParameterModeRelative:
				prefix = "%"
			}
			params[i] = fmt.Sprintf("%s%d", prefix, p[off+i+1])
		}
		out := append([]string{revOpMap[int(op)]}, params...)
		ret = append(ret, strings.Join(out, " "))
		off += paramCnt + 1
		if op == intcode.OPCodeHalt {
			break
		}
	}
	return ret
}
