canvas.width = window.innerWidth
canvas.height = window.innerHeight

const clay = []
let world
async function run() {
  const res = await fetch('input/day18')
  const data = await res.text()
  const input = data.split('\n').filter(Boolean)
  const w = input[0].length
  const h = input.length

  world = new CAWorld({
    width: w,
    height: h,
    cellSize: 20
  })

  world.registerCellType('acre', {
    process (neighbors) {
      const trees = this.countSurroundingCellsWithValue(neighbors, 'wasTrees')
      const lumberyards = this.countSurroundingCellsWithValue(neighbors, 'wasLumberyard')
      this.lumberyard = (this.wasTrees && lumberyards >= 3) || (this.wasLumberyard && trees >= 1 && lumberyards >= 1)
      this.trees = (this.wasOpen && trees >= 3) || (this.wasTrees && !this.lumberyard)
      this.open = !this.trees && !this.lumberyard
    },
    reset () {
      this.wasOpen = this.open
      this.wasTrees = this.trees
      this.wasLumberyard = this.lumberyard
    }
  }, function() {
    // init
    this.open = false
    this.trees = false
    this.lumberyard = false
  })
  

  world.initialize([
    { name: 'acre', distribution: 100 }
  ])

  for(const y in input) {
    const row = input[y]
    for(let x = 0; x < row.length; x++) {
      const key = row[x] === '|' ? 'trees' : (row[x] === '#' ? 'lumberyard' : 'open')
      world.grid[y][x][key] = true
    }
  }

  let count = 1000000000
  function render() {
    const ctx = canvas.getContext('2d')
    ctx.clearRect(0,0,canvas.width,canvas.height)
    ctx.save()
    ctx.font = '10pt mono'
    ctx.translate(0,100)
    const CS = world.cellSize
    const counts = {
      open: 0,
      trees: 0,
      lumberyard: 0
    }
    for (let y=0; y<world.height; y++) {
      for (let x=0; x<world.width; x++) {
         const cell = world.grid[y][x]
         ctx.fillStyle = cell.trees ? 'green' : (cell.lumberyard ? 'brown' : 'black')
         ctx.beginPath()
         ctx.rect(x*CS,y*CS,CS,CS)
         ctx.fill()
         if(cell.open) counts.open++
         if(cell.trees) counts.trees++
         if(cell.lumberyard) counts.lumberyard++
      }
    }
    ctx.restore()
    ctx.save()
    ctx.fillStyle = 'white'
    ctx.fillText(`Open: ${counts.open}`,100,10)
    ctx.fillText(`Trees: ${counts.trees}`,100,20)
    ctx.fillText(`Lumberyard: ${counts.lumberyard}`,100,30)
    ctx.fillText(`Value: ${counts.trees * counts.lumberyard}`,100,40)
    ctx.fillText(`Rem: ${count}`,100,50)
    const value = counts.trees * counts.lumberyard
    ctx.restore()
    return value
  }
  //renderLoop()
  // render()
  btn.addEventListener('click', function () {
    // world.step()
    renderLoop()
  })
  btn.innerHTML = 'start'
  // while(count--) {
    // render()
  // }
  const values = new Set()
  const dup = new Map()
  let finalValue = 0
  // setTimeout(() => renderLoop(), 8000)
  function renderLoop() {
    requestAnimationFrame(renderLoop)
    world.step()
    if(count) {
      count--
    } else {
      const value = render()
      if(!finalValue) finalValue = value
      const ctx = canvas.getContext('2d')
      ctx.save()
      ctx.font = '24pt mono'
      ctx.fillStyle = 'green'
      //ctx.fillText(`Final Value: ${finalValue}`, 200, 50)
      ctx.fillText(`Loop Found!`, 200, 50)
      ctx.restore()
    }
    if(finalValue) return
    const value = render()
    if(values.has(value)) {
      if(dup.has(value)) {
        const last = dup.get(value)
        const diff = last - count
        console.log('test', last, diff)
        for(let i=0;i<diff;i++) {
          world.step()
        }
        count -= diff
        const val1 = render()
        for(let i=0;i<diff;i++) {
          world.step()
        }
        count -= diff
        const val2 = render()
        if(val1 === val2 && val1 === value) {
          count %= diff
          console.log('Loop Found!', diff, count)
          dup.clear()
        }
      }
      dup.set(value, count)
    }
    values.add(value)
  }
  
  function toXY (ind) {
    const rx = ind % w
    const ry = Math.floor(ind / w)
    return [rx, ry]
  }
  
  function toInd([x,y]) {
    return (y * w) + x
  }
}

run().catch(err => console.error(err))
