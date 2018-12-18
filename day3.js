const fs = require('fs')

const argv = process.argv.slice(2)
const input = require('./input').map(v => {
  const [,id,x,y,w,h] = v.match(/#(\d+) @ (\d+),(\d+): (\d+)x(\d+)/).map(v => parseInt(v))
  return { id, x, y, w, h }
})

const rect = new Array(1000 * 1000)
for(let i=0;i<(1000*1000);i++) {
  rect[i] = 0
}

for(const { id, x, y, w, h } of input) {
  for(let xx = x; xx < x + w; xx++) {
    for(let yy = y; yy < y + h; yy++) {
      rect[toInd(xx,yy)]++
    }
  }
}

let part2 = ''
for(const { id, x, y, w, h } of input) {
  let overlaps = false
  for(let xx = x; xx < x + w; xx++) {
    for(let yy = y; yy < y + h; yy++) {
      overlaps |= rect[toInd(xx,yy)] !== 1
    }
  }
  if (!overlaps) {
    part2 = id
  }
}


const part1 = rect.filter(v => v > 1).length
console.log(part1)
console.log(part2)

function toInd(x,y) {
  return (y * 1000) + x
}