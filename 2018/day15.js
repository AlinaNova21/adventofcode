const input = require('./input')
const PF = require('pathfinding')

const w = input.reduce((r,v) => Math.max(r, v.length), 0)
const h = input.length

const ELF_POWER = +process.argv[3] || 3

class Unit {
  constructor(type, cell) {
    this.type = type
    this.cell = cell
    this.ap = type === 'E' ? ELF_POWER : 3
    this.hits = 200
  }
  get elf() { return this.type === 'E' }
  get gob() { return this.type === 'G' }
  get pos() {
    return toXY(this.cell)
  }
  set pos(v) {
    this.cell = toInd(v)
  }
  attack(u) {
    u.hits -= this.ap
  }
  toJSON() {
    const { type, pos, cell } = this
    return { type, pos, cell }
  }
}

const grid = input.reduce((r,v) => [...r,...v.split('')],[])
const units = Object.entries(grid).filter(([,v]) => v === 'E' || v === 'G').map(([i,t]) => new Unit(t,+i))
units.forEach(u => grid[u.cell] = '.')

const pfgrid = new PF.Grid(w,h)
grid.forEach((g,i) => {
  pfgrid.setWalkableAt(...toXY(i), g === '.')
})

const finder = new PF.AStarFinder()

let cnt = 0
render()
while(round()) {
  cnt++
  render()
  // units.filter(u => u.hits > 0)
  // .forEach(u => console.log(`${u.type}(${u.hits}) ${u.cell} ${u.pos}`))
  console.log()

}
//console.log(`Round ${cnt}`)
//console.log(process.argv)
  
function round() {
  const alive = units.filter(u => u.hits > 0)
  for(const unit of alive) {
    units.sort((a,b) => a.cell - b.cell)
    const alive = units.filter(u => u.hits > 0)
    const elves = alive.filter(u => u.type === 'E')
    const goblins = alive.filter(u => u.type === 'G')
    if(unit.hits <= 0) continue
    const cgrid = pfgrid.clone()
    const enemies = unit.gob ? elves : goblins
    const friends = unit.elf ? elves : goblins
    if(!enemies.length) {
      const winner = unit.gob ? 'goblins' : 'elves'
      const score = (unit.gob ? goblins : elves).reduce((r,u) => r + u.hits, 0)
      console.log(winner, score, cnt, score * cnt)
      return false
    }
    alive.forEach(u => {
      cgrid.setWalkableAt(...u.pos, u.type !== unit.type || u === unit)
    })
    enemies.forEach(u => {
      u._path = finder.findPath(...unit.pos,...u.pos,cgrid.clone())
    })
    const closest = enemies.filter(u => u._path.length).reduce((r,u) => r && r._path.length <= u._path.length ? r : u, 0)
    if(!closest) continue

    if(closest._path.length !== 2) {
      const lgrid = Array.from(grid)
      alive.forEach(u => {
        lgrid[u.cell] = u.type
      })
      lgrid[closest.cell] = '.'
      flood(lgrid, closest.cell, 0)
      let dirs = [-w,-1,1,w].map(d => [+lgrid[unit.cell + d],(unit.cell + d)]).filter(a => !isNaN(a[0]))
      const shortest = dirs.reduce((r, [d]) => Math.min(r,d),1e10)
      dirs = dirs.filter(d => d[0] === shortest)
      dirs.sort((a,b) => a[1] - b[1])
      // console.log(unit, closest, dirs, shortest)
      // render(lgrid)
      const [[,ind]=[]] = dirs
      if(ind) {
        unit.cell = ind
      }
      // dx = Math.min(Math.max(dx, -1), 1)
      // dy = Math.min(Math.max(dy, -1), 1)
      // if(dy < 0) dx = 0
      // if(dx !== 0) dy = 0
      // unit.pos = [unit.pos[0] + dx, unit.pos[1] + dy]
    }
    {
      enemies.forEach(u => {
        u._path = finder.findPath(...unit.pos,...u.pos,cgrid.clone())
      })
      const neighbors = enemies.filter(u => u._path.length === 2)
      neighbors.sort((a,b) => a.cell - b.cell)
      if(neighbors.length) {
        // Adjacent      
        const weakest = neighbors.reduce((ret, u) => ret && ret.hits <= u.hits ? ret : u, null)
        unit.attack(weakest)
        if(weakest.type === 'E' && weakest.hits <= 0 && ELF_POWER != 3) {
          console.log('Elf died',ELF_POWER)
          return false
        }
        // console.log(weakest)
      }
    }
  }
  return true
}

function flood(grid, node) {
  const q = [node]
  const nq = []
  let dist = 0
  while(q.length) {
    const nq = []
    q.forEach(i => {
      grid[i] = dist
      nq.push(...[-w,-1,+1,+w].map(v => i + v).filter(i => grid[i] === '.'))
    })
    q.splice(0,q.length,...(new Set(nq)))
    dist++
  }
}

function render(rgrid = grid,spc='') {
  const lgrid = rgrid.map(a=>a)
  units.filter(u=>u.hits > 0).forEach(u => {
    lgrid[u.cell] = u.type
  })
  let out = `Round ${cnt}`
  for(let i=0;i<lgrid.length;i++) {
    if(i % w === 0) out += "\n"
    out += spc + lgrid[i]
  }
  console.log(out + "\n")
}


function dist(a,b) {
  return Math.abs(a[0] - b[0]) + Math.abs(a[1] - b[1])
}
function toXY (ind) {
  const rx = ind % w
  const ry = Math.floor(ind / w)
  return [rx, ry]
}

function toInd([x,y]) {
  return (y * w) + x
}
