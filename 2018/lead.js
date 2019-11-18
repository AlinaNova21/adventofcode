const input = require('./leaderboard.json')
const moment = require('moment')
const members = Object.values(input.members)
/*members.forEach(m => {
  delete m.completion_day_level // cleanup output
})*/

members.forEach(m => {
  m.last_star_ts = parseInt(m.last_star_ts) * 1000
  m.times = Object.values(m.completion_day_level).map(d => parseInt(d['2'].get_star_ts) * 1000)
})

members.sort((a,b) => b.local_score - a.local_score)

console.log(members)

for(const [ind, member] of Object.entries(members)) {
  const firstPlace = members[0]
  const nextPlace = ind > 1 ? members[ind-1] : members[0]
  const parts = []
  parts.push(leftpad(parseInt(ind)+1,3))
  parts.push(leftpad(member.local_score,5))
  parts.push(leftpad(nextPlace.local_score - member.local_score,5))
  parts.push(leftpad(firstPlace.local_score - member.local_score,5))
  const vel = Math.floor(member.local_score / member.stars)
  parts.push(leftpad(vel,3))
  parts.push(rightpad(member.name, 20))
  parts.push(moment(parseInt(member.last_star_ts) * 1000).format('HH:mm'))
  const times = member.times.map(t => moment(t).format('HH:mm'))
  console.log(parts.join(' '))
  // times.forEach(t => console.log(t))
}

function leftpad(s,len,char = ' ') {
  return (char.repeat(len) + s).slice(-len)
}
function rightpad(s,len,char = ' ') {
  return (s+char.repeat(len)).slice(0,len)
}