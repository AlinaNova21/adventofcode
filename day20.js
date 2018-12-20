const input = require('./input')

const tests = [
  '^N(E|W)N$',
  '^ENWWW(NEEE|SSE(EE|N))$',
]

function stackIt(str) {
  const path = str.slice(1,-1)
  const stack = []
  const distances = {}
  let cur = [0,0]
  for(let i = 0; i < path.length; i++) {
    const v = path[i]
    if(v === '(') {
      stack.push(cur)
      continue
    }
    if(v === ')') {
      cur = stack.pop()
      continue
    }
    if(v === '|') {
      cur = stack.slice(-1)[0]
      continue
    }
    const npos = ({ N: [0,-1], E: [1,0], S: [0,1], W: [-1,0] })[v].map((v,i) => v + cur[i])
    const pos1 = distances[npos] || (distances[cur] + 1)
    const pos2 = distances[cur] + 1
    distances[npos] = Math.min(pos1 || 1, pos2 || 1)
    cur = npos
  }
  const max = Math.max(...Object.values(distances))
  const overk = Object.values(distances).map(x => x >= 1000).reduce((l,v) => l + v, 0)
  return [max,overk]
}

// tests
input
  .map(v => (console.log(v),v))
  .map(v => stackIt(v))
  .map(v => (console.log(v),v))