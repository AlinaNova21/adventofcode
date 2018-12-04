const fs = require('fs')
const moment = require('moment')

const argv = process.argv.slice(2)
const input = fs.readFileSync(argv[0], 'utf8').split("\n").filter(Boolean).map(v => {
  const [,time,data] = v.match(/^\[(.+)\] (.+)$/)
  const ts = moment(time)
  console.log(time, ts, data)
  let event = ''
  let id = -1
  if (data === 'wakes up') event = 'wake'
  if (data === 'falls asleep') event = 'asleep'
  if (data.startsWith('Guard')) {
    id = parseInt(data.split(' ')[1].slice(1))
  }
  return { ts, event, id }
})

input.sort((a,b) => {
  return b.ts.isBefore(a.ts, 'minute') ? 1 : -1
})

let lastID = -1
input.forEach(i => {
  if (i.id === -1) i.id = lastID
  else lastID = i.id
})

const guards = {}
let sleep = 0
for(const { id, event, ts } of input) {
  const guard = guards[id] = guards[id] || { id, sleep: 0, minutes: Array(61).fill(0) }
  if (event === 'asleep') {
    sleep = ts
  }
  if (event === 'wake') {
    const dur = moment.duration(ts.diff(sleep)).asMinutes()
    guard.sleep += dur
    for(let min = 0; min < dur; min++) {
      guard.minutes[min + sleep.minute()]++
    }
  }
}

const arr = Object.values(guards)
arr.sort((a,b) => b.sleep - a.sleep)
arr.forEach(g => {
  const val = g.minutes.reduce((a,b) => Math.max(a,b), 0)
  const min = g.minutes.indexOf(val)
  g.mostAsleep = { min, val }
})
{
  const guard = arr[0]
  const max = guard.minutes.reduce((a,b) => Math.max(a,b), 0)
  const minute = guard.minutes.indexOf(max)
  console.log(guard.id, minute, guard.id * minute)
}
arr.sort((a,b) => b.mostAsleep.val - a.mostAsleep.val)
{
  const guard = arr[0]
  console.log(guard.id, guard.mostAsleep.min, guard.mostAsleep.val, guard.mostAsleep.min * guard.id)
}
