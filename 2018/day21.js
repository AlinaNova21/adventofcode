const input = require('./input').map(v => {
  const [opcode,a=0,b=0,c=0] = v.split(' ')
  return [opcode, +a, +b, +c]
})

class Machine {
  constructor() {
    this.reg = [0,0,0,0,0,0]
    this.ip = 0
    this.ipo = 0
    this.ipr = 0
    this.mem = []
    this.ops = {
      addr: (a,b) => this.reg[a] + this.reg[b],
      addi: (a,b) => this.reg[a] + b,    
      mulr: (a,b) => this.reg[a] * this.reg[b],
      muli: (a,b) => this.reg[a] * b,
      banr: (a,b) => this.reg[a] & this.reg[b],
      bani: (a,b) => this.reg[a] & b,
      borr: (a,b) => this.reg[a] | this.reg[b],
      bori: (a,b) => this.reg[a] | b,
      setr: (a,b) => this.reg[a],
      seti: (a,b) => a,
      gtir: (a,b) => a > this.reg[b] ? 1 : 0,
      gtri: (a,b) => this.reg[a] > b ? 1 : 0,
      gtrr: (a,b) => this.reg[a] > this.reg[b] ? 1 : 0,
      eqir: (a,b) => a === this.reg[b] ? 1 : 0,
      eqri: (a,b) => this.reg[a] === b ? 1 : 0,
      eqrr: (a,b) => this.reg[a] === this.reg[b] ? 1 : 0,
    }
  }
  tick() {
    if(this.halted) return false
    if(this.ip >= this.mem.length) {
      this.halt()
      return false
    }
    const [opcode, a, b, c] = this.mem[this.ip]
    const fn = this.ops[opcode]
    this.reg[this.ipr] = this.ip
    const lastReg = ''//= Array.from(this.reg)
    this.reg[c] = fn(a,b)
    this.ip = this.reg[this.ipr]
    if(this.dbg) {
      console.log(`ip=${this.ip} [${lastReg}] ${opcode} ${a} ${b} ${c} [${this.reg}]`)
    }
    this.ip++
    return true
  }
  halt() {
    this.halted = true
  }
  reset() {
    this.ip = 0
  }
}
{
  const m = new Machine()
  m.dbg = false
  m.mem = input.slice(1)
  m.ipr = input[0][1]
  const s = new Set()
  const brk = 28
  while(!m.halted) {
    // m.dbg = m.ip === 28
    if(m.ip === brk) {
      const [opcode, a, b, c] = m.mem[m.ip]
      const other = a === 0 ? m.reg[b] : m.reg[a]
      // console.log(other,s.size)
      if(s.has(other)) {
        console.log('Repeat found')
        break
      } else {
        s.add(other)
      }
    }
    m.tick()
  }
  const arr = Array.from(s)
  console.log(arr[0], arr.slice(-1)[0])
}
