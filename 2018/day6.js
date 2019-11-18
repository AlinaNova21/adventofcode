const input = require('./input').map(v => v.split(', ').map(v => parseInt(v))).map((v,i) => [...v,String.fromCharCode('0'.charCodeAt(0) + i),i])

const minx = input.reduce((r, [x,y]) => Math.min(x, r),0) - 1
const maxx = input.reduce((r, [x,y]) => Math.max(x, r),0) + 1

const miny = input.reduce((r, [x,y]) => Math.min(y, r),0) - 1
const maxy = input.reduce((r, [x,y]) => Math.max(y, r),0) + 2

const w = maxx - minx
const h = maxy - miny

function toXY (ind) {
  const rx = ind % h
  const ry = Math.floor(ind / h)
  return [rx + minx, ry + miny]
}

function toInd([x,y]) {
  return ((y - miny) * w) + (x - minx)
}

function renderGrid(grid) {
  let out = ''
  for(let i=0;i < grid.length; i++) {
    out += grid[i]
    if(i % h === h - 1) {
      out += "\n"
    }
  }
  console.log(out)
}

function dist(a,b) {
  return Math.abs(a[0] - b[0]) + Math.abs(a[1] - b[1])
}

function popGrid(grid) {
  for(const i in grid) {
    grid[i] = ' '
  }
  for(const [x,y,l] of input) {
    const ind = toInd([x,y])
    grid[ind] = l
  }
}

function part1() {
  const grid = Array(w * h).fill(' ')
  popGrid(grid)
  for(let i=0;i < grid.length; i++) {
    const p1 = toXY(i)
    let shortest = 1e10000
    for(const p2 of input) {
      const d = dist(p1,p2)
      if (!d) continue
      if (d === shortest) {
        grid[i] = '.'
      }
      if (d < shortest) {
        grid[i] = p2[2]
        shortest = d
      }
    }
  }
  
  const counts = {}
  outer: for(const [,,l] of input) {
    const points = grid.map((v,i) => [v,i]).filter(([v]) => v === l)
    for(let [v,i] of points) {
      const [x,y] = toXY(i)
      if(x === minx || x === maxx || y === miny || y === maxy) {
        continue outer
      }
    }
    counts[l] = points.length
  }
  renderGrid(grid)
  console.log('part1',Object.values(counts).reduce((a,b) => Math.max(a,b)))
}
function part2() {
  const grid = Array(w * h).fill(' ')
  popGrid(grid)
  for(let i=0;i < grid.length; i++) {
    const p1 = toXY(i)
    const d = input.reduce((r, p2) => r + dist(p1,p2), 0)
    if (d < 10000) {
      grid[i] = '#'
    }
  }
  renderGrid(grid)
  const points = grid.filter(v => v === '#')
  console.log('part2', points.length)
}

part1()
part2()

