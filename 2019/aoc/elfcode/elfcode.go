package elfcode

import "fmt"

type Op func(a, b int) int

const OP_COUNT = 16

type Machine struct {
	Register [4]int
	OpMap    []int
	Ops      []Op
	Memory   [][4]int
	PC       int
}

func NewMachine() *Machine {
	m := &Machine{
		OpMap: make([]int, OP_COUNT),
	}
	for i := 0; i < OP_COUNT; i++ {
		m.OpMap[i] = i
	}
	m.Ops = []Op{
		m.addr,
		m.addi,
		m.mulr,
		m.muli,
		m.banr,
		m.bani,
		m.borr,
		m.bori,
		m.setr,
		m.seti,
		m.gtir,
		m.gtri,
		m.gtrr,
		m.eqir,
		m.eqri,
		m.eqrr,
	}
	return m
}

func (m *Machine) String() string {
	return fmt.Sprintf("[%d %d %d %d]", m.Register[0], m.Register[1], m.Register[2], m.Register[3])
}

func (m *Machine) Step() error {
	if m.PC >= len(m.Memory) {
		return fmt.Errorf("Memory Access Error")
	}
	o := m.Memory[m.PC]
	fn := m.Ops[m.OpMap[o[0]]]
	m.Register[o[3]] = fn(o[1], o[2])
	m.PC++
	return nil
}

func (m *Machine) Reset() {
	m.PC = 0
	m.Register = [4]int{}
	m.Memory = make([][4]int, 0)
}

func (m *Machine) addr(a, b int) int {
	return m.Register[a] + m.Register[b]
}

func (m *Machine) addi(a, b int) int {
	return m.Register[a] + b
}

func (m *Machine) mulr(a, b int) int {
	return m.Register[a] * m.Register[b]
}

func (m *Machine) muli(a, b int) int {
	return m.Register[a] * b
}

func (m *Machine) banr(a, b int) int {
	return m.Register[a] & m.Register[b]
}

func (m *Machine) bani(a, b int) int {
	return m.Register[a] & b
}

func (m *Machine) borr(a, b int) int {
	return m.Register[a] | m.Register[b]
}

func (m *Machine) bori(a, b int) int {
	return m.Register[a] | b
}

func (m *Machine) setr(a, _ int) int {
	return m.Register[a]
}

func (m *Machine) seti(a, _ int) int {
	return a
}

func (m *Machine) gtir(a, b int) int {
	if a > m.Register[b] {
		return 1
	}
	return 0
}

func (m *Machine) gtri(a, b int) int {
	if m.Register[a] > b {
		return 1
	}
	return 0
}

func (m *Machine) gtrr(a, b int) int {
	if m.Register[a] > m.Register[b] {
		return 1
	}
	return 0
}

func (m *Machine) eqir(a, b int) int {
	if a == m.Register[b] {
		return 1
	}
	return 0
}

func (m *Machine) eqri(a, b int) int {
	if m.Register[a] == b {
		return 1
	}
	return 0
}

func (m *Machine) eqrr(a, b int) int {
	if m.Register[a] == m.Register[b] {
		return 1
	}
	return 0
}
