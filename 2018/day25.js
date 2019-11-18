const input = require('./input').map(v => ({
  pos: v.split(',').map(v => +v),
  con: 0,
  dists: []
}))

let lastCon = 0

for(let a of input) {
  a.dists = input.map(b => dist(a.pos,b.pos))
  const inds = Object.entries(a.dists).filter(([i,d]) => d > 0 && d <= 3).map(([i,v]) => i)
  if(a.pos[0] === 9) console.log(inds)
  let con = inds.map(i => input[i].con).find(Boolean)
  if(!con) con = Math.random().toString(36).slice(-6)
  a.con = con
  for(const i of inds) {
    const b = input[i]
    if(b.con && b.con !== con) {
      input.filter(c => c.con === b.con).forEach(c => c.con = con) // Merge
    }
    b.con = con
  }
}
const count = new Set(input.map(i => i.con)).size
console.log(input)
console.log(count)
function dist(a,b) {
  return Math.abs(a[0] - b[0]) + Math.abs(a[1] - b[1]) + Math.abs(a[2] - b[2]) + Math.abs(a[3] - b[3])
}