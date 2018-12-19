let r0 = 1
let r1 = 0
let r2 = 0
let r3 = 0
let r4 = 0
let r5 = 0

r4 = 987
if(r0 === 1) {
  r4 += 10550400
  r0 = 0
}
for(r1 = 1; r1 <= r4; r1++) {
  for(r3 = 1; r3 <= r4; r3++) {
    if(r4 === r1 * r3) {
      r0 += r1
    }
  }
}

console.log(r0)