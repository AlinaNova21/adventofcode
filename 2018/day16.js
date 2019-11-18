const input = require('./input')

class Machine {
  constructor() {
    this.reg = [0,0,0,0]
    this.pc = 0
    this.mem = []
    const gr = (i) => this.reg[i]
    const sr = (i,v) => this.reg[i] = v
    this.ops = [
      /* addr */ (a,b) => gr(a) + gr(b),
      /* addi */ (a,b) => gr(a) + b,    
      /* mulr */ (a,b) => gr(a) * gr(b),
      /* muli */ (a,b) => gr(a) * b,
      /* banr */ (a,b) => gr(a) & gr(b),
      /* bani */ (a,b) => gr(a) & b,
      /* borr */ (a,b) => gr(a) | gr(b),
      /* bori */ (a,b) => gr(a) | b,
      /* setr */ (a,b) => gr(a),
      /* seti */ (a,b) => a,
      /* gtir */ (a,b) => a > gr(b) ? 1 : 0,
      /* gtri */ (a,b) => gr(a) > b ? 1 : 0,
      /* gtri */ (a,b) => gr(a) > gr(b) ? 1 : 0,
      /* eqir */ (a,b) => a === gr(b) ? 1 : 0,
      /* eqri */ (a,b) => gr(a) === b ? 1 : 0,
      /* eqri */ (a,b) => gr(a) === gr(b) ? 1 : 0,
    ]
    this.opmap = Array.from(this.ops.keys())
  }
  tick() {
     const [opcode, a, b, c] = this.mem[this.pc]
     const fn = this.ops[this.opmap[opcode]]
     this.reg[c] = fn(a,b)
     this.pc++
  }
  reset() {
    this.pc = 0
  }
}

const sets = []
const prog = []
for(let i=0;i<input.length;i+=3) {
  const [before,op,after] = input.slice(i, i+3)
  if(!before.startsWith('Before')) {
    prog.push(...input.slice(i).map(l => l.split(' ').map(o => +o)))
    break
  }
  sets.push({
    before: JSON.parse(before.split(': ')[1]),
    op: op.split(' ').map(o => +o),
    after: JSON.parse(after.split(': ')[1])
  })
}

for(const s of sets) {
  const m = new Machine()
  const ops = []
  for(let i=0;i<16;i++) {
    m.reset()
    m.reg = Array.from(s.before)
    m.mem = [Array.from(s.op)]
    m.mem[0][0] = i
    m.tick()
    if(m.reg.join(' ') === s.after.join(' ')) {
      ops.push(i)
    }
  }
  s.ops = ops
}

console.log(sets.filter(s => s.ops.length >= 3).length)

const opmap = []
for(let i = 0; i<16; i++) {
  const ops = new Set()
  const possible = new Set([0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15])
  sets.filter(s => s.op[0] === i).forEach(s => {
    possible.forEach(p => {
      if(!s.ops.includes(p)) possible.delete(p)
    })
  })
  opmap.push([i,possible])
}
const finalopmap = Array(16)
while(opmap.length) {
  const ind = opmap.findIndex(o => o[1].size === 1)
  const [[o,s]] = opmap.splice(ind, 1)
  const [v] = Array.from(s)
  opmap.forEach(([i,s]) => s.delete(v))
  finalopmap[o] = Array.from(s)[0]
}

{
  const m = new Machine()
  m.opmap = finalopmap
  m.mem = prog
  while(m.pc < prog.length) {
    console.log(m.mem[m.pc], m.reg)
    m.tick()
  }
  console.log(m.reg[0])
  console.log(m.opmap)
}