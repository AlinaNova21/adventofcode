const input = require('./input')

let reacted = true
const regexParts = []
for(let i=65;i<65+26;i++) {
  const a = String.fromCharCode(i)
  const b = a.toLowerCase()
  regexParts.push(a+b,b+a)
}
const regex = new RegExp(regexParts.join('|'),'g')
console.log('part1',react(input[0]))

const results = []
for(let i=65;i<65+26;i++) {
  const a = String.fromCharCode(i)
  const b = a.toLowerCase()
  const result = react(input[0].replace(new RegExp(a,'ig'), ''))
  results.push(result)
}
const min = results.reduce((a,b) => Math.min(a, b), 1e100)
console.log('part2', min)

function react(res) {
  while(true) {
    let n = res.replace(regex, '')
    if (n === res) break
    res = n
  }
  return res.length
}

