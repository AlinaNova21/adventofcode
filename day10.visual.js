const params = new URLSearchParams(location.search)
const ticks = parseInt(params.get('ticks') || 3)

function next() {
  location.href = `day10.html?ticks=${ticks+1}`
}
function prev() {
  location.href = `day10.html?ticks=${ticks-1}`
}

canvas.width = window.innerWidth
canvas.height = window.innerHeight

async function run() {
  const res = await fetch('input.day10')
  const data = await res.text()
  const input = data.split('\n').filter(Boolean).map(v => {
    const [x,y,vx,vy] = v.match(/(-?\d+)/g).map(v => parseInt(v))
    return [x,y,vx,vy]
  })

  const minx = input.reduce((r, [x,y]) => Math.min(x, r),0) - 1
  const maxx = input.reduce((r, [x,y]) => Math.max(x, r),0) + 1
  
  const miny = input.reduce((r, [x,y]) => Math.min(y, r),0) - 1
  const maxy = input.reduce((r, [x,y]) => Math.max(y, r),0) + 2
  const w = maxx - minx
  const h = maxy - miny

  const ctx = canvas.getContext('2d')
  const mult = 1
  // canvas.width = w / mult
  // canvas.height = h / mult
  // ctx.translate(-minx,-miny)
  canvas.width *= mult
  canvas.height *= mult
  console.log(minx,maxx,miny,maxy)
  // ctx.translate((maxx + minx) / 2 - 300, (maxy + miny) / 2 - 100)
  console.log(ticks)

  function render() {
    requestAnimationFrame(render)
    const ctx = canvas.getContext('2d')
    const ticks = 10659 + Math.sin(Date.now() / 5000) * 50
    const len = ticks
    ctx.clearRect(0,0,canvas.width,canvas.height)
    ctx.save()
    ctx.fillStyle = 'white'
    ctx.fillText(Math.floor(ticks).toString(),100,100)
    ctx.translate(-180,-100)
    ctx.translate(canvas.width/2, canvas.height/2)
    ctx.fillStyle = 'white'
    for(const [x,y,vx,vy] of input) {
      ctx.moveTo(x,y)
      ctx.lineTo(x+(vx*len),y + (vy*len))
    }
    // ctx.stroke()
    for(const [x,y,vx,vy] of input) {
      ctx.beginPath()
      ctx.rect(x + (vx * ticks), y + (vy * ticks), 1, 1)
      ctx.fill()
    }
    ctx.restore()
  }
  
  render()
}

run().catch(e => console.error(e))