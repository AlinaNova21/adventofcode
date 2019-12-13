package assembler

import (
	"io"
	"math"

	"github.com/ags131/adventofcode/2019/aoc/intcode"
)

/*
OPCodeNop OPCode = iota
OPCodeAdd
OPCodeMul
OPCodeInput
OPCodeOutput
OPCodeJumpIfTrue
OPCodeJumpIfFalse
OPCodeLessThan
OPCodeEquals
OpCodeSetRelBase
OPCodeHalt OPCode = 99
*/

var opMap = map[string]intcode.OPCode{
	"NOP":  intcode.OPCodeNop,
	"ADD":  intcode.OPCodeAdd,
	"MUL":  intcode.OPCodeMul,
	"IN":   intcode.OPCodeInput,
	"OUT":  intcode.OPCodeOutput,
	"JNZ":  intcode.OPCodeJumpIfTrue,
	"JZ":   intcode.OPCodeJumpIfFalse,
	"LT":   intcode.OPCodeLessThan,
	"EQ":   intcode.OPCodeEquals,
	"REL":  intcode.OpCodeSetRelBase,
	"HALT": intcode.OPCodeHalt,
}

// Assemble assembles IntCode assembly into an IntCode program
func Assemble(r io.Reader) (intcode.Program, error) {
	var err error
	parser, err := GetParser(&Program{})
	if err != nil {
		return nil, err
	}
	// fmt.Println("PARSER", parser.String())
	prog := &Program{}
	err = parser.Parse(r, prog)
	if err != nil {
		return nil, err
	}
	instructions := make([][]int, len(prog.Instructions))
	labels := map[string]int{}
	data := []int{}
	length := 0
	for _, ins := range prog.Instructions {
		if ins.Label != nil {
			labels[ins.Label.Name] = length
		}
		length += 1 + len(ins.Params)
	}
	varOffset := int(math.Ceil(float64(length)/100) * 100)
	for i, ins := range prog.Instructions {
		if ins.Data != "" {
			v := 0
			if ins.Params[0].Identifier != nil {
				v = labels[*ins.Params[0].Identifier]
			}
			// if ins.Params[0].Boolean != nil {
			// 	v = int(*ins.Params[0].Boolean)
			// }
			if ins.Params[0].Number != nil {
				v = *ins.Params[0].Number
			}
			data = append(data, v)
			length++
			continue
		}
		if ins.Params != nil {
			for _, p := range ins.Params {
				p.Mode = (p.Mode + 1) % 3
			}
		}
		// fmt.Print(ins.Op, ins.Params, ins.Data)
		op := int(opMap[ins.Op])
		params := make([]int, len(ins.Params))
		for j, param := range ins.Params {
			v := 0
			if param.Identifier != nil {
				if _, ok := labels[*param.Identifier]; !ok {
					labels[*param.Identifier] = varOffset
					varOffset++
					data = append(data, 0)
				}
				v = labels[*param.Identifier]
			}
			if param.Number != nil {
				v = *param.Number
			}
			op += int(param.Mode) * int(math.Pow10(j+2))
			params[j] = v
		}
		inst := []int{op}
		inst = append(inst, params...)
		instructions[i] = inst
	}
	ret := make(intcode.Program, length)
	off := 0
	for _, ins := range instructions {
		for i, v := range ins {
			ret[off+i] = v
		}
		off += len(ins)
	}
	return ret, nil
}
