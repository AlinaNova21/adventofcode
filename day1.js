const fs = require('fs')

const input = require('./input').map(v => parseInt(v))

const day1 = input.reduce((l,v) => l + v, 0)
console.log(day1)

const freqs = new Set()
let dup = false
let freq = 0
while(dup === false) {
  for (const diff of input) {
    freq += diff
    if (freqs.has(freq)) {
      dup = freq
      console.log(freq)
      break
    }
    freqs.add(freq)    
  }
}
console.log(dup)