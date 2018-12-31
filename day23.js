const input = require('./input').map(v => v.match(/(-?\d+)/g).map(v => +v))

const mostPowerful = input.reduce((r,v) => r && r[3] > v[3] ? r : v, false)
const dists = input.map(v => [...v,dist(v, mostPowerful)])
console.log(mostPowerful)
console.log(dists.map(v => [...v,v[4] <= mostPowerful[3]]))
const inRange = dists.filter(v => v[4] <= mostPowerful[3]).length

console.log(inRange)

// const rangeAvg = avg(input.map(v => v[3]))
// const centerPoint = [avg(input.map(v => v[0] * v[3])) / rangeAvg, avg(input.map(v => v[1] * v[3])) / rangeAvg, avg(input.map(v => v[2] * v[3])) / rangeAvg]
// console.log(centerPoint)

/*
const dirs = []

const counts = 4
for(let i = 0; i < counts; i++) {
  const r = (360 / counts) * i
  const x = Math.round(1 * Math.sin(r))
  const y = Math.round(1 * Math.cos(r))
  dirs.push([x,y])
}
console.log(dirs)

const dirRange = [0,0,0,0,0,0,0,0]

const closest = input.map(v => [...v, dist(v, [0,0,0]) - v[3]]).reduce((r,v) => r && r[4] > v[4] ? r : v, 0)
const farthest = input.map(v => [...v, dist(v, [0,0,0]) + v[3]]).reduce((r,v) => r && r[4] > v[4] ? r : v, 0)
let startRange = closest[4]
let endRange = farthest[4]
let found = false
const origin = [0,0,0]
function sweep(range) {
  const points = dirs.map(([x,y]) => [x*range,y*range])
  for(let i in points) {
    points[i].push(countInRange(point[i]))
  }
}
*/
// let scale = 100
let scale = 100000000
let best = [0,0,0,0,1e10]
while(scale !== 1) {
  scale /= 10
  console.log(scale,best)
  const r = 100 * scale
  for(let z = -r;z<=r;z+=scale) {
    for(let y = -r;y<=r;y+=scale) {
      for(let x = -r;x<=r;x+=scale) {
        const pos = [best[0]+x,best[1]+y,best[2]+z]
        const count = input
          .filter(v => dist(v,pos) <= v[3])
          .length
        const distance = pos.reduce((a,b) => a + Math.abs(b), 0)
        if(count > best[3] || (count === best[3] && distance <= best[5])) {
          best = [...pos,count,distance]
        }
      }
    }
  }
}

console.log(best)

function countInRange(pos) {
  return input.map(v => [...v,dist(v, pos)]).filter(v => v[4] <= v[5]).length
}

function dist(a,b) {
  return Math.abs(a[0]-b[0]) + Math.abs(a[1]-b[1]) + Math.abs(a[2]-b[2])
}

function avg(arr) {
  return arr.reduce((a,b) => a + b, 0) / arr.length
}
