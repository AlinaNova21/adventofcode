package intcode

import (
	"context"
	"fmt"
	"math"
)

// OPCode determines the operation to execute
type OPCode int

// OpCode values
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
	OpCodeSetRelBase
	OPCodeHalt OPCode = 99
)

// ParameterMode changes how parameters are read
type ParameterMode int

// ParameterMode values
const (
	ParameterModePosition = iota
	ParameterModeImmediate
	ParameterModeRelative
)

// Machine is an intcode interpreter
type Machine struct {
	IP         int
	RelBase    int
	Memory     []int
	Input      chan int
	Output     chan int
	InputFunc  func() int
	OutputFunc func(int)
	Result     int
	Stopped    bool
	haltCtx    context.Context
}

// NewMachine creates a new IntCode vm
func NewMachine(p Program) *Machine {
	return newMachine(p, 0, 0)
}

func newMachine(p Program, ip int, rel int) *Machine {
	chIn := make(chan int, 0)
	chOut := make(chan int, 1)
	ctx, cancel := context.WithCancel(context.Background())
	mem := make([]int, 1e6)
	copy(mem, p)
	m := &Machine{
		IP:      ip,
		RelBase: rel,
		Memory:  mem,
		Input:   chIn,
		Output:  chOut,
		InputFunc: func() int {
			return <-chIn
		},
		OutputFunc: func(v int) {
			chOut <- v
		},
		haltCtx: ctx,
	}
	go func() {
		m.Run()
		cancel()
		close(chOut)
	}()
	return m
}

// Fork creates a copy of the Machine
func (m *Machine) Fork() *Machine {
	return newMachine(m.Memory, m.IP, m.RelBase)
}

// Wait waits for the vm to halt and returns the last output
func (m *Machine) Wait() int {
	<-m.haltCtx.Done()
	return m.Result
}

// Halted returns a channel indicating when the vm halts
func (m *Machine) Halted() <-chan struct{} {
	return m.haltCtx.Done()
}

// Run executes until halt
func (m *Machine) Run() {
	for m.Step() {
	}
}

// Get returns the value at the specified memory address
func (m *Machine) Get(addr int) int {
	return m.Memory[addr]
}

// Set sets the value at the specified memory address
func (m *Machine) Set(addr, v int) {
	m.Memory[addr] = v
}

func (m *Machine) getMode(off int) ParameterMode {
	v := m.Memory[m.IP]
	t1 := int(math.Pow(10, float64(off+2)))
	t2 := int(math.Pow(10, float64(off+1)))
	return ParameterMode((v % t1) / t2)
}

func (m *Machine) getParam(off int) int {
	addr := m.IP + off
	pm := m.getMode(off)
	switch pm {
	case ParameterModePosition:
		return m.Memory[m.Memory[addr]]
	case ParameterModeImmediate:
		return m.Memory[addr]
	case ParameterModeRelative:
		return m.Memory[m.RelBase+m.Memory[addr]]
	}
	return 0
}

func (m *Machine) setResult(off, v int) {
	addr := m.IP + off
	pm := m.getMode(off)
	switch pm {
	case ParameterModePosition:
		m.Memory[m.Memory[addr]] = v
	case ParameterModeImmediate:
		m.Memory[addr] = v
	case ParameterModeRelative:
		m.Memory[m.RelBase+m.Memory[addr]] = v
	}
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
		m.setResult(1, m.InputFunc())
		m.IP += 2
	case OPCodeOutput:
		v := m.getParam(1)
		m.OutputFunc(v)
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
	case OpCodeSetRelBase:
		m.RelBase += m.getParam(1)
		m.IP += 2
	case OPCodeHalt:
		m.Stopped = true
		return false
	default:
		panic(fmt.Errorf("Unknown opcode! %d", op))
	}
	return true
}
