package intcode

import (
	"context"
	"fmt"
	"math"
)

// OPCode determines the operation to execute
type OPCode int

const (
	OPCodeNop OPCode = iota
	OPCodeAdd
	OPCodeMul
	OPCodeInput
	OPCodeOutput
	OPCodeJumpIfTrue
	OPCodeJumpIfFalse
	OPCodeLessThan
	OPCodeEquals
	OPCodeHalt OPCode = 99
)

// ParameterMode changes how parameters are read
type ParameterMode int

const (
	ParameterModePosition = iota
	ParameterModeImmediate
)

type InputFunc func() int
type OutputFunc func(v int)

// Machine is an intcode interpreter
type Machine struct {
	IP      int
	Memory  []int
	Input   chan int
	Output  chan int
	Result  int
	haltCtx context.Context
}

func NewMachine(p Program) *Machine {
	chIn := make(chan int, 1)
	chOut := make(chan int, 1)
	ctx, cancel := context.WithCancel(context.Background())
	m := &Machine{
		Memory:  []int(p.Clone()),
		Input:   chIn,
		Output:  chOut,
		haltCtx: ctx,
	}
	go func() {
		m.Run()
		cancel()
		close(chOut)
	}()
	return m
}

func (m *Machine) Wait() int {
	<-m.haltCtx.Done()
	return m.Result
}

func (m *Machine) Halted() <-chan struct{} {
	return m.haltCtx.Done()
}

// Run executes until halt
func (m *Machine) Run() {
	for m.Step() {
	}
}

func (m *Machine) Get(addr int) int {
	return m.Memory[addr]
}

func (m *Machine) Set(addr, v int) {
	m.Memory[addr] = v
}

func (m *Machine) getParam(off int) int {
	v := m.Memory[m.IP]
	t1 := int(math.Pow(10, float64(off+2)))
	t2 := int(math.Pow(10, float64(off+1)))
	pm := ParameterMode((v % t1) / t2)
	addr := m.IP + off
	switch pm {
	case ParameterModePosition:
		return m.Memory[m.Memory[addr]]
	case ParameterModeImmediate:
		return m.Memory[addr]
	}
	return 0
}

func (m *Machine) setResult(off, v int) {
	addr := m.IP + off
	m.Memory[m.Memory[addr]] = v
}

// Step executes one instruction
func (m *Machine) Step() bool {
	op := OPCode(m.Memory[m.IP] % 100)
	switch op {
	case OPCodeAdd:
		m.setResult(3, m.getParam(1)+m.getParam(2))
		m.IP += 4
	case OPCodeMul:
		m.setResult(3, m.getParam(1)*m.getParam(2))
		m.IP += 4
	case OPCodeInput:
		m.setResult(1, <-m.Input)
		m.IP += 2
	case OPCodeOutput:
		v := m.getParam(1)
		m.Output <- v
		m.Result = v
		m.IP += 2
	case OPCodeJumpIfTrue:
		if m.getParam(1) > 0 {
			m.IP = m.getParam(2)
		} else {
			m.IP += 3
		}
	case OPCodeJumpIfFalse:
		if m.getParam(1) == 0 {
			m.IP = m.getParam(2)
		} else {
			m.IP += 3
		}
	case OPCodeLessThan:
		v := 0
		if m.getParam(1) < m.getParam(2) {
			v = 1
		}
		m.setResult(3, v)
		m.IP += 4
	case OPCodeEquals:
		v := 0
		if m.getParam(1) == m.getParam(2) {
			v = 1
		}
		m.setResult(3, v)
		m.IP += 4
	case OPCodeHalt:
		return false
	default:
		panic(fmt.Errorf("Unknown opcode! %d", op))
	}
	return true
}
