const input = require('./input')

// left
// straight
// right
// repeat

const w = input.reduce((r,v) => Math.max(r, v.length), 0)
const h = input.length

const grid = input.reduce((r,v) => [...r,...v.split('')],[])

const base = grid.map(v => {
  if(['v','^'].includes(v)) return '|'
  if(['<','>'].includes(v)) return '-'
  return v
})

const transform = {
  '/': ([x,y]) => [-y,-x],
  '\\': ([x,y]) => [y,x],
  '|': ([x,y]) => [x,y],
  '-': ([x,y]) => [x,y],
  's': ([x,y]) => [x,y],
  'l': ([x,y]) => [y,-x],
  'r': ([x,y]) => [-y,x],
}
const vect = {
  'v': [0,1],
  '^': [0,-1],
  '<': [-1,0],
  '>': [1,0],
}

const revVect = Object.entries(vect)
  .reduce((ret, [d,v]) => {
    ret[v.toString()] = d
    return ret
  },{})

const gridDirs = { 'v': w, '^': -w, '<': -1, '>': 1 }

const carts = Object.entries(grid).filter(([,v]) => isCart(v)).map(([ind,val]) => {
  ind = parseInt(ind)
  return {
    pos: toXY(ind),
    vec: vect[val],
    crashed: false,
    turn: 0,
    ind
  }
})

const collisions = []

//render()
let cnt = 1
while(tick(cnt++));// render()

function tick(time) {
  //console.log('tick',time)
  const moved = new Set()
  carts.sort((a,b) => a.ind - b.ind)
  const occupied = new Set(carts.map(c=>c.ind))
  for(const cart of carts) {
    if(cart.crashed) continue
    move(cart)
    const ind = toInd(cart.pos)
    if(base[ind] === '+') {
      cart.vec = [transform.l,transform.s,transform.r][cart.turn++ % 3](cart.vec)
    } else {
      cart.vec = transform[base[ind]](cart.vec)
    }
    const otherCart = carts.find(c => !c.crashed && c !== cart && c.ind === ind)
    if(otherCart) {
      cart.crashed = true
      otherCart.crashed = true
      collisions.push({
        carts: [cart, otherCart],
        time,
        pos: cart.pos
      })
    }
  }
  const remCarts = carts.filter(c => !c.crashed)
  if(remCarts.length < 2) {
    console.log(`First collision was at ${collisions[0].pos}`)
    console.log(`Final cart was at ${remCarts[0].pos}`)
    return false
  }
  return true
}

function render() {
  const grid = Array.from(base)
  carts
    .filter(c => !c.crashed)
    .forEach(({ pos, vec }) => grid[toInd(pos)] = revVect[vec.toString()])
  let out = ''
  for(let i=0;i<grid.length;i++) {
    if(i % w === 0) out += "\n"
    out += grid[i]
  }
  console.log(out)
}

function rotate(cart, dir) {
  const dirs = ['^','>','v','<']
  const ind = dirs.indexOf(cart) + dir
  return dirs[(ind + 4) % 4]
}

function toXY (ind) {
  const rx = ind % w
  const ry = Math.floor(ind / w)
  return [rx, ry]
}

function toInd([x,y]) {
  return (y * w) + x
}

function isCart(v) {
  return ['v','^','<','>'].includes(v)
}

function move(cart) {
  cart.pos = [cart.pos[0] + cart.vec[0], cart.pos[1] + cart.vec[1]]
  cart.ind = toInd(cart.pos)
}
