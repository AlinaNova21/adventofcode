// canvas.width = window.innerWidth
// canvas.height = window.innerHeight
const START = parseInt(process.argv[2] || 1)
const END = parseInt(process.argv[3] || 300)

const w = 300
const h = 300

function toXY (ind) {
  const rx = ind % h
  const ry = Math.floor(ind / h)
  return [rx, ry]
}

function toInd([x,y]) {
  return (y * w) + x
}

async function run() {
  // const res = await fetch('input.day11')

  const cells = new Array(300*300)

  const gridSerial = 8141

  let maxSize = [0,0,1,0]
  for(let i=0;i<w*h;i++) {
    const [x,y] = toXY(i)
    const rackID = x + 10
    cells[i] = (Math.floor((((rackID * y) + gridSerial) * rackID) / 100) % 10) - 5
  }

  for(let i=START;i<END;i++) {
    const groupSize = i
  
    const offs = [] // 0,1,2,w+0,w+1,w+2,w+w+0,w+w+1,w+w+2]
    for(let x = 0; x < groupSize; x++) {
      for(let y = 0; y < groupSize; y++) {
        offs.push(toInd([x,y]))
      }
    }
    const groups = cells
      .filter((_,i) => {
        const [x,y] = toXY(i)
        return (x < w - groupSize) && (y < h - groupSize)
      })
      .map((_,i) => offs.map(ii => cells[i + ii]).reduce((a,b) => a + b, 0))
  
    let max = [0,0]
  
    groups.forEach((pl,i) => {
      if (pl > max[0]) {
        max = [pl, i]
      }
    })
    if (maxSize[3] < max[0]) {
      const [x,y] = toXY(max[1])
      maxSize = [x,y,groupSize,max[0]]
    }
    console.log(max, ...toXY(max[1]), groupSize)
  }
  console.log(maxSize)
}

run().catch(e => console.error(e))