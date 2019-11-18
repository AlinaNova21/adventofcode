const input = require('./input').map(v => {
  const [opcode,a=0,b=0,c=0] = v.split(' ')
  return [opcode, +a, +b, +c]
})

let out = ''
let mem = []
input.splice(0,1)
for(const [opcode,a,b,c] of input) {
  let ipr = 0
  const map = {
    addr: () => `r${c} = r${a} + r${b}`,
    addi: () => `r${c} = r${a} + ${b}`,    
    mulr: () => `r${c} = r${a} * r${b}`,
    muli: () => `r${c} = r${a} * ${b}`,
    banr: () => `r${c} = r${a} & r${b}`,
    bani: () => `r${c} = r${a} & ${b}`,
    borr: () => `r${c} = r${a} | r${b}`,
    bori: () => `r${c} = r${a} | ${b}`,
    setr: () => `r${c} = r${a}`,
    seti: () => `r${c} = ${a}`,
    gtir: () => `r${c} = ${a} > r${b} ? 1 : 0`,
    gtri: () => `r${c} = r${a} > ${b} ? 1 : 0`,
    gtrr: () => `r${c} = r${a} > r${b} ? 1 : 0`,
    eqir: () => `r${c} = ${a} === r${b} ? 1 : 0`,
    eqri: () => `r${c} = r${a} === ${b} ? 1 : 0`,
    eqrr: () => `r${c} = r${a} === r${b} ? 1 : 0`,
  }
  const fn = map[opcode]
  mem.push(fn())
}

let ip = 0
let ipr = 2
let r0 = 1
let r1 = 0
let r2 = 0
let r3 = 0
let r4 = 0
let r5 = 0
console.log(mem.join("\n"))
while(ip < mem.length) {
  r2 = ip
  let out = `ip=${ip} [${r0},${r1},${r2},${r3},${r4},${r5}] `
  eval(mem[ip])
  ip = r2
  ip++
  out += `[${r0},${r1},${r2},${r3},${r4},${r5}]`
  //console.log(out, mem[ip])
}
console.log(r0)