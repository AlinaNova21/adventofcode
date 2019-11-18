const input = require('./input')
  .map(v => {
    const [a,b] = v.split(', ').map(v => v.split('='))
    a[1] = +a[1]
    b[1] = range(...b[1].split('..').map(v => +v))
    return [a,b]
  })

const clay = []

for(const [a,b] of input) {
  const p = a[0] === 'x' ? (v) => [a[1], v] : (v) => [v, a[1]]
  for(let i=0;i<b[1].length;i++) {
    clay.push(p(b[1][i]))
  }
}

const minx = clay.reduce((l,[v]) => Math.min(l,v), 1e10) - 1
const ominy = clay.reduce((l,[,v]) => Math.min(l,v), 1e10)
const miny = 0
const maxx = clay.reduce((l,[v]) => Math.max(l,v), 0) + 3
const maxy = clay.reduce((l,[,v]) => Math.max(l,v), 0) + 1

const w = maxx - minx
const h = maxy - miny

const spring = [500,0]

console.log(w,h,minx,miny,maxx,maxy)
const grid = Array(w*h)
grid.fill('.')
grid[toInd(spring)] = '+'
clay.forEach((c) => {
  grid[toInd(c)] = '#'
})

render()
let count = 0
let notSettled = 0
// while(count--) {
while(true) {
  count++
  if(!tick()) {
    console.log(notSettled)
    notSettled++
  } else {
    notSettled = 0
  }
  if(count % 10 === 0) {
    // render()
  }
  if(notSettled > 10) {
    render()
    console.log('stop',count)
    break
  }
}

{ 
  const settled = grid.slice(ominy * w).reduce((r,g) => r + (g === '~'), 0)
  const flowing = grid.slice(ominy * w).reduce((r,g) => r + (g === '|'), 0)
  console.log(settled + flowing, settled)
}

function tick() {
  let ind = toInd(spring)
  return flow(ind)
}

function flow(ind) {
  let solid = (v) => ['#','~'].includes(v)
  let settled = false
  outer: while(true) {
    let flowed = false
    while(true) {
      const down = grid[ind + w]
      if(solid(down)) {
        break
      }
      ind += w
      if(ind >= grid.length) {
  //      console.log('bottom')
        break outer
      }
      grid[ind] = '|'
    }
    let ldist = 1
    while(true) {
      const left = grid[ind - ldist]
      const down = grid[ind - ldist + w]
      if(solid(left)) break
      grid[ind-ldist] = '|'
      if(!solid(down)) {
        settled |= flow(ind - ldist)
        flowed = true
        break
      }
      ldist++
    }
    let rdist = 1
    while(true) {
      const right = grid[ind + rdist]
      const down = grid[ind + rdist + w]
      if(solid(right)) break
      grid[ind+rdist] = '|'
      if(!solid(down)) {
        settled |= flow(ind + rdist)
        flowed = true
        break
      }
      rdist++
    }
    if(!flowed) {
      settled = true
      for(let i=-ldist+1;i<rdist;i++) {
        grid[ind + i] = '~'
      }
    }
    break outer
  }
  return settled
}

function render() {
  let out = ''
  for(let i=0;i<grid.length;i++) {
    if(i % w === 0) out += "\n"
    out += grid[i]
  }
  console.log(out + "\n")
}

function range(s,e) {
  const arr = []
  for(let i = s; i <= e; i++) {
    arr.push(i)
  }
  return arr
}

function toXY (ind) {
  const rx = ind % w
  const ry = Math.floor(ind / w)
  return [rx + minx, ry + miny]
}

function toInd([x,y]) {
  return ((y - miny) * w) + (x - minx)
}
