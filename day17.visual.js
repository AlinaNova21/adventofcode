canvas.width = window.innerWidth
canvas.height = window.innerHeight

const clay = []

async function run() {
  const res = await fetch('input.day17')
  const data = await res.text()
  const input = data.split('\n').filter(Boolean)
    .map(v => {
      const [a,b] = v.split(', ').map(v => v.split('='))
      a[1] = +a[1]
      b[1] = b[1].split('..').map(v => +v)
      return [a,b]
    })

  
  for(const [a,b] of input) {
    const p = a[0] === 'x' ? (v) => [a[1], v] : (v) => [v, a[1]]
    const r = range(...b[1])
    for(let i=0;i<r.length;i++) {
      clay.push(p(r[i]))
    }
  }
  const minx = clay.reduce((l,[v]) => Math.min(l,v), 1e10) - 1
  const ominy = clay.reduce((l,[,v]) => Math.min(l,v), 1e10)
  const miny = 0
  const maxx = clay.reduce((l,[v]) => Math.max(l,v), 0) + 3
  const maxy = clay.reduce((l,[,v]) => Math.max(l,v), 0) + 2
  
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
  
  let count = 0
  let notSettled = 0
  let q = []
  setTimeout(() => {
    loop()
    renderLoop()
  }, 5000)
  function loop(v) {
    count++
    let ind = q.shift() || toInd(spring)
    const [s, queue] = flow(ind)
    q.push(...queue)
    q = Array.from(new Set(q))
    if(!q.length) {
      console.log('stop',count)
      return
    } else {
      setTimeout(loop, 1)
    }
  }
  function renderLoop() {
    if(q.length) {
      requestAnimationFrame(renderLoop)
    }
    render()
  }
  
  
  { 
    const settled = grid.slice(ominy * w).reduce((r,g) => r + (g === '~'), 0)
    const flowing = grid.slice(ominy * w).reduce((r,g) => r + (g === '|'), 0)
    console.log(settled + flowing, settled)
  }
  
  function flow(ind) {
    let solid = (v) => ['#','~'].includes(v)
    let settled = false
    let queue = []
    outer: while(true) {
      let flowed = false
      while(true) {
        const down = grid[ind + w]
        if(solid(down)) {
          break
        }
        ind += w
        if(ind >= grid.length) {
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
          // settled |= flow(ind - ldist)
          flowed = true
          queue.push(ind - ldist)
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
          // settled |= flow(ind + rdist)
          flowed = true
          queue.push(ind + rdist)
          break
        }
        rdist++
      }
      if(!flowed) {
        settled = true
        for(let i=-ldist+1;i<rdist;i++) {
          grid[ind + i] = '~'
        }
        queue.push(ind - w)
      }
      break outer
    }
    return [settled, queue]
  }
  
  function render() {
    const ctx = canvas.getContext('2d')
    const settled = grid.reduce((r,g,i) => {
      g === '~' ? r.push(i) : ''
      return r
    }, [])
    const flowing = grid.reduce((r,g,i) => {
      g === '|' ? r.push(i) : ''
      return r
    }, [])
    const maxind = Math.max(settled.reduce((a,b) => Math.max(a,b),0), flowing.reduce((a,b) => Math.max(a,b),0))
    const [ox,oy] = toXY(maxind)
    canvas.viewOff = canvas.viewOff || 0
    const scale = 2
    const yoff = Math.max(0, oy - (canvas.width / (scale * 2)))
    if(canvas.viewOff < yoff) {
      if(yoff - canvas.viewOff < 50) {
        canvas.viewOff -= 0.5
      }
      if(yoff - canvas.viewOff > 100) canvas.viewOff+=3
      canvas.viewOff++
    }
    ctx.clearRect(0,0,canvas.width,canvas.height)
    ctx.save()
    ctx.scale(scale,scale)
    ctx.translate(-minx,Math.floor(-canvas.viewOff))
    ctx.strokeStyle = 'brown'
    ctx.lineWidth = 1.5
    ctx.fillStyle = 'white'
    ctx.beginPath()
    for(const [a,b] of input) {
      const p = a[0] === 'x' ? (v) => [a[1], v] : (v) => [v, a[1]]
      const start = a[0] === 'x' ? [a[1],b[1][0]] : [b[1][0],a[1]]
      const end = a[0] === 'x' ? [a[1],b[1][1]] : [b[1][1],a[1]]
      ctx.moveTo(...start)
      ctx.lineTo(...end)
    }
    ctx.stroke()
    
    ctx.beginPath()
    ctx.fillStyle = 'lightblue'
    for(const ind of flowing) {
      ctx.rect(...toXY(ind), 1, 1)
    }
    ctx.fill()
    ctx.beginPath()
    ctx.fillStyle = 'blue'
    for(const ind of settled) {
      ctx.rect(...toXY(ind), 1, 1)
    }
    ctx.fill()
    ctx.translate(650,canvas.viewOff)
    ctx.fillStyle = 'white'
    ctx.fillText(`Cycle: ${count}`, 0, 10)
    ctx.fillText(`Settled: ${settled.length}`, 0, 20)
    ctx.fillText(`Flowing: ${flowing.length}`, 0, 30)
    ctx.fillText(`All: ${settled.length+flowing.length}`, 0, 40)
    ctx.restore()
    return
    
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
}

run().catch(err => console.error(err))
