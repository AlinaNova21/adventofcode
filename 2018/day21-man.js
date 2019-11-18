let r0 = 0
let r1 = 0
let r2 = 0
let r3 = 0
let r4 = 0
let r5 = 0

while(72 !== 123 & 456);
console.log('start')
r3 = 0

const s = new Set()

outer: while(true) {
  r5 = r3 | 65536
  r3 = 15028787
  while(true) {
    r2 = r5 & 255
    r3 = (((r3 + r2) & 16777215) * 65899) & 16777215
    // console.log(r5)
    if(256 > r5) {
      if(r3 === r0) {
        console.log('halted!',r2,r5,r3)
        break outer
      }
      if(s.has(r3)) {
        console.log('seen before', r3)
        break outer
      }
      s.add(r3)
      continue outer
    }
    r2 = 0
    do {
      r4 = (r2 + 1) * 256
      r2++
    } while(r4 <= r5)
    r5 = r2 - 1
  }
}
const arr = Array.from(s)
console.log(arr[0],arr.slice(-1)[0])
//console.log(s.size)
