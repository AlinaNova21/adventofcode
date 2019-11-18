const fs = require('fs')
const out = fs.createWriteStream('day19.asm')
const input = require('./input').map(v => {
  const [opcode,a=0,b=0,c=0] = v.split(' ')
  return [opcode, +a, +b, +c]
})

let mem = []
input.splice(0,1)
input.splice(25,0,['raw','cmp byte [acc],1'])
input.splice(26,0,['raw','je add29'])
const w = (v) => {
  out.write(`${v.trim()}\n`)
}
w(`extern printf`)
w(`section .data`)
w(`acc: dd 1`)
w(`fmt: db "out: %d",10,0`)
w(`dbgfmt: db "ip=%d [%d, %d, %d, %d, %d, %d]",10,0`)
w(`section .text`)
w(`global main`)
w(`main:`,0)
w(`mov ebx,0`)
w(`mov ecx,0`)
w(`mov edx,0`)
w(`mov esi,0`)
w(`mov edi,0`)
const reg = ['[acc]','ebx','ecx','edx','esi','edi']
// const reg = ['r10','r11','r12','r13','r14','r15']
for(let i=0;i < input.length; i++) {
  const [opcode,a,b,c] = input[i]
//  w(`;${opcode} ${a} ${b} ${c}`)
  out.write(`add${+i}: `)
  w(`mov ${reg[2]}, ${+i}`)
  w(`call dbg`)
  const j = (v) => {
    if(v > input.length) {
      return w(`jmp finish`)
    }
    w(`jmp add${v}`)
  }
  switch(opcode) {
    case 'raw':
      w(a)
      break
    case 'addr':
      w(`mov eax, ${reg[a]}`)
      w(`add eax, ${reg[b]}`)
      w(`mov ${reg[c]}, eax`)
      break
    case 'addi':
      if(reg[c] !== reg[a])
        w(`mov ${reg[c]}, ${reg[a]}`)
      w(`add ${reg[c]}, ${b}`)
      if(c==2) j(i+b+1)
      break
    case 'mulr':
      w(`mov eax, ${reg[a]}`)
      w(`imul eax, ${reg[b]}`)
      w(`mov ${reg[c]}, eax`)
      if(c==2) w(`jmp finish`)
      break
    case 'muli':
      w(`mov eax, ${reg[a]}`)
      w(`imul eax, ${b}`)
      w(`mov ${reg[c]}, eax`)
      break
    case 'banr':
      w(`mov eax, ${reg[a]}`)
      w(`and eax, ${reg[b]}`)
      w(`mov ${reg[c]}, eax`)
      break
    case 'bani':
      if(reg[c] !== reg[a])
      w(`mov ${reg[c]}, ${reg[a]}`)
      w(`and ${reg[c]}, ${b}`)
      break
    case 'borr':
      w(`mov eax, ${reg[a]}`)
      w(`or eax, ${reg[b]}`)
      w(`mov ${reg[c]}, eax`)
      break
    case 'bori':
      if(reg[c] !== reg[a])
      w(`mov ${reg[c]}, ${reg[a]}`)
      w(`or ${reg[c]}, ${b}`)
      break
    case 'setr':
      if(reg[c] !== reg[a])
      w(`mov ${reg[c]}, ${reg[a]}`)
      break
    case 'seti':
      if(c===0) {
        w(`mov BYTE ${reg[c]}, ${a}`)
      } else {
        w(`mov ${reg[c]}, ${a}`)
      }
      if(c==2) j(a+1)
      break
    case 'gtrr':
      w(`mov eax, ${reg[a]}`)
      w(`cmp eax, ${reg[b]}`)
      w(`jg add${+i+3}`)
      i++
      break
    case 'eqrr':
      w(`mov eax, ${reg[a]}`)
      w(`cmp eax, ${reg[b]}`)
      w(`je add${+i+3}`)
      i++
      break
    default:
      w(`MISSING: ${opcode}`)
      break
    // gtir: () => [`r${c} = ${a} > r${b} ? 1 : 0`],
    // gtri: () => [`r${c} = r${a} > ${b} ? 1 : 0`],
    // gtrr: () => [`r${c} = r${a} > r${b} ? 1 : 0`],
    // eqir: () => [`r${c} = ${a} === r${b} ? 1 : 0`],
    // eqri: () => [`r${c} = r${a} === ${b} ? 1 : 0`],
    // eqrr: () => [`r${c} = r${a} === r${b} ? 1 : 0`],
  }
}
// w(`mov edx, 1`)
// w(`mov ecx, acc`)
// w(`mov ebx, 1`)
// w(`mov eax, 4`)
// w(`int 0x80`)
out.write(`dbg: `)
w(`ret`)
w(`push dword ${reg[5]}`)
w(`push dword ${reg[4]}`)
w(`push dword ${reg[3]}`)
w(`push dword ${reg[2]}`)
w(`push dword ${reg[1]}`)
w(`push dword ${reg[0]}`)
w(`push dword ${reg[5]}`)
w(`push dword ${reg[4]}`)
w(`push dword ${reg[3]}`)
w(`push dword ${reg[2]}`)
w(`push dword ${reg[1]}`)
w(`push dword ${reg[0]}`)
w(`push dword ${reg[2]}`)
w(`push dword dbgfmt`)
w(`call printf`)
w(`add esp, 32`)
w(`pop dword ${reg[0]}`)
w(`pop dword ${reg[1]}`)
w(`pop dword ${reg[2]}`)
w(`pop dword ${reg[3]}`)
w(`pop dword ${reg[4]}`)
w(`pop dword ${reg[5]}`)
w(`ret`)
out.write(`finish: `)
w(`push dword [acc]`)
w(`push dword fmt`)
w(`call printf`)
w(`call dbg`)
w(`add esp, 8`)
w(`retn`)
out.end()