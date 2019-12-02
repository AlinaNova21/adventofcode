package intcode

// OPCode determines the operation to execute
type OPCode int

const (
	OPCodeNop OPCode = iota
	OPCodeAdd
	OPCodeMul
	OPCodeHalt OPCode = 99
)

// Machine is an intcode interpreter
type Machine struct {
	IP     int
	Memory []int
}

// Run executes until halt
func (m *Machine) Run() {
	for m.Step() {
	}
}

// Step executes one instruction
func (m *Machine) Step() bool {
	switch OPCode(m.Memory[m.IP]) {
	case OPCodeAdd:
		m.Memory[m.Memory[m.IP+3]] = m.Memory[m.Memory[m.IP+1]] + m.Memory[m.Memory[m.IP+2]]
		m.IP += 4
	case OPCodeMul:
		m.Memory[m.Memory[m.IP+3]] = m.Memory[m.Memory[m.IP+1]] * m.Memory[m.Memory[m.IP+2]]
		m.IP += 4
	case OPCodeHalt:
		return false
	}
	return true
}
