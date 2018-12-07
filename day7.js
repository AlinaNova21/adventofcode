const input = require('./input').reduce((ret, v) => {
  const [,requires,,,,,,step] = v.split(' ')
  ret[step] = ret[step] || []
  ret[requires] = ret[requires] || []
  ret[step].push(requires)
  return ret
}, {})

const BASE_TIME = 60
const WORKERS = 5
const workers = []
for(let i=0;i<WORKERS;i++) {
  workers.push({ step: null, time: null })
}

console.log(input)
const steps = Object.entries(input)
const start = steps.filter(([s,r]) => !r.length).map(v => v[0])
const ret = []

const ready = new Set()
let time = 0
start.forEach(s => ready.add(s))
while(ready.size || ret.length < steps.length) {
  doStep()
  console.log(time++, workers, ret.join(' '))
  workers.forEach(worker => {
    if (!worker.step) return
    worker.time--
    if (worker.time === 0) {
      ret.push(worker.step)
      worker.step = null
    }
  })
}
console.log(time, workers, ret.join(' '))

function doStep() {
  const vals = Array.from(ready)
  vals.sort(alphaSort[0])
  // console.log(vals, ret)
  for(const step of vals) {
    const [,reqs] = steps.find(([s]) => s === step)
    const reqMet = !reqs.find(r => !ret.includes(r))
    // console.log(step, reqMet)
    if (!reqMet) continue
    const worker = workers.find(w => !w.step)
    if (!worker) break
    ready.delete(step)
    //ret.push(step)
    worker.step = step
    worker.time = BASE_TIME + step.charCodeAt(0) - 64
    const next = steps.filter(([s,r]) => r.includes(step)).map(v => v[0])
    next.sort(alphaSort())
    next.forEach((s) => ready.add(s))
    // console.log(next)
  }
}

console.log('ret',ret.join(''),time)
/*
function recurse(step) {
  const next = steps.filter(([s,r]) => r.includes(step)).map(v => v[0])
  next.sort(alphaSort(0))
  console.log('recurse', step, next)
  ret.push(step)
  next.forEach(recurse)
}
recurse(start)
ret.reverse()
ret.splice(0,ret.length,...Array.from(new Set(ret)))
ret.reverse()
console.log('rec',ret.join(''))
*/
/*
for(const step in input) {
  const requires = input[step]
  requires.sort(alphaSort)
  const ret = []
  requires.forEach(s => {
    ret.push(
  })
}
*/

function alphaSort(prop = null) {
  return (a,b) => {
    if(prop !== null) return b[prop].charCodeAt(0) - a[prop].charCodeAt(0)
    return b.charCodeAt(0) - a.charCodeAt(0)
  }
}

/*
const letters = input.map(([,r]) => r)
const [last] = input.find(([s]) => !letters.includes(s))
const steps = [...letters,last].map(s => {
  const requires = input.filter(([step
})


function recurse(last) {
  const ret = []
  const steps = input.filter(([s,r]) => s === last)
  steps.sort((a,b) => a[1].charCodeAt(0) - b[1].charCodeAt(0))
  console.log('recurse',last,steps.map(s=>s[1]))
  steps.forEach(([,r]) => {
    ret.push(recurse(r))
  })
  ret.push(last)
  return ret
}
const ret = recurse(last)
//ret.splice(0,ret.length,...Array.from(new Set(ret)))
console.log(ret)

/*


const start = input.find(({step}) => !letters.includes(step)).step
letters.unshift(start)

console.log(letters.join(''))
*/