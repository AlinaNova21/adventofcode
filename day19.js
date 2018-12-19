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
    const gr = (i) => this.reg[i]
    const sr = (i,v) => this.reg[i] = v
    this.ops = {
      addr: (a,b) => gr(a) + gr(b),
      addi: (a,b) => gr(a) + b,    
      mulr: (a,b) => gr(a) * gr(b),
      muli: (a,b) => gr(a) * b,
      banr: (a,b) => gr(a) & gr(b),
      bani: (a,b) => gr(a) & b,
      borr: (a,b) => gr(a) | gr(b),
      bori: (a,b) => gr(a) | b,
      setr: (a,b) => gr(a),
      seti: (a,b) => a,
      gtir: (a,b) => a > gr(b) ? 1 : 0,
      gtri: (a,b) => gr(a) > b ? 1 : 0,
      gtrr: (a,b) => gr(a) > gr(b) ? 1 : 0,
      eqir: (a,b) => a === gr(b) ? 1 : 0,
      eqri: (a,b) => gr(a) === b ? 1 : 0,
      eqrr: (a,b) => gr(a) === gr(b) ? 1 : 0,
    }
  }
  tick() {
    if(this.halted) return false
    if(this.ip >= this.mem.length || this.ip < 0) {
      this.halt()
      return false
    }
    const [opcode, a, b, c] = this.mem[this.ip]
    const fn = this.ops[opcode]
    this.reg[this.ipr] = this.ip
    const lastReg = Array.from(this.reg)
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
  // m.dbg = true
  m.mem = input.slice(1)
  m.ipr = input[0][1]
  while(!m.halted) {
    m.tick()
  }
  console.log(m.reg[0])
}
{
  const m = new Machine()
  m.dbg = true
  m.mem = input.slice(1)
  m.ipr = input[0][1]
  m.reg[0] = 1
  // while(!m.halted) {
    // m.tick()
  // }
  console.log(m.reg[0])
}